package forking

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/holiman/uint256"
	"golang.org/x/exp/maps"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/trie/trienode"
)

type accountDiff struct {
	// no diff if nil
	Nonce *uint64

	// no diff if nil
	Balance *uint256.Int

	// no diff if not present in map. Deletions are zero-value entries.
	Storage map[common.Hash]common.Hash

	CodeHash *common.Hash // no code-diff if nil
}

func (d *accountDiff) Copy() *accountDiff {
	var out accountDiff
	if d.Nonce != nil {
		v := *d.Nonce // copy the value
		out.Nonce = &v
	}
	if d.Balance != nil {
		out.Balance = d.Balance.Clone()
	}
	if d.Storage != nil {
		out.Storage = maps.Clone(d.Storage)
	}
	if d.CodeHash != nil {
		h := *d.CodeHash
		out.CodeHash = &h
	}
	return &out
}

type ForkedAccountsTrie struct {
	// stateRoot that this diff is based on top of
	stateRoot common.Hash

	// source to retrieve data from when it's not in the diff
	src ForkSource

	// Accounts diff. Deleted accounts are set to nil.
	// Warning: this only contains finalized state changes.
	// The state itself holds on to non-flushed changes.
	accountDiff map[common.Address]*accountDiff

	// Stores new contract codes by code-hash
	codeDiff map[common.Hash][]byte
}

var _ state.Trie = (*ForkedAccountsTrie)(nil)

func (f *ForkedAccountsTrie) Copy() *ForkedAccountsTrie {
	accountsCopy := make(map[common.Address]*accountDiff, len(f.accountDiff))
	for k, v := range f.accountDiff {
		accountsCopy[k] = v.Copy()
	}
	codesCopy := make(map[common.Hash][]byte)
	for k, v := range f.codeDiff {
		codesCopy[k] = bytes.Clone(v)
	}
	return &ForkedAccountsTrie{
		stateRoot:   f.stateRoot,
		accountDiff: accountsCopy,
		codeDiff:    codesCopy,
	}
}

// ClearDiff clears the flushed changes. This does not clear the warm state changes.
// To fully clear, first Finalise the forked state that uses this trie, and then clear the diff.
func (f *ForkedAccountsTrie) ClearDiff() {
	f.accountDiff = make(map[common.Address]*accountDiff)
	f.codeDiff = make(map[common.Hash][]byte)
}

// ContractCode is not directly part of the vm.State interface,
// but is used by the ForkDB to retrieve the contract code.
func (f *ForkedAccountsTrie) ContractCode(addr common.Address, codeHash common.Hash) ([]byte, error) {
	diffAcc, ok := f.accountDiff[addr]
	if ok {
		if diffAcc.CodeHash != nil && *diffAcc.CodeHash != codeHash {
			return nil, fmt.Errorf("account code changed to %s, cannot get code %s of account %s", *diffAcc.CodeHash, codeHash, addr)
		}
		if code, ok := f.codeDiff[codeHash]; ok {
			return code, nil
		}
		// if not in codeDiff, the actual code has not changed.
	}
	code, err := f.src.Code(addr)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve code: %w", err)
	}
	// sanity-check the retrieved code matches the expected codehash
	if h := crypto.Keccak256Hash(code); h != codeHash {
		return nil, fmt.Errorf("retrieved code of %s hashed to %s, but expected %s", addr, h, codeHash)
	}
	return code, nil
}

// ContractCodeSize is not directly part of the vm.State interface,
// but is used by the ForkDB to retrieve the contract code-size.
func (f *ForkedAccountsTrie) ContractCodeSize(addr common.Address, codeHash common.Hash) (int, error) {
	code, err := f.ContractCode(addr, codeHash)
	if err != nil {
		return 0, fmt.Errorf("cannot get contract code to determine code size: %w", err)
	}
	return len(code), nil
}

func (f *ForkedAccountsTrie) GetKey(bytes []byte) []byte {
	panic("arbitrary key lookups on ForkedAccountsTrie are not supported")
}

func (f *ForkedAccountsTrie) GetAccount(address common.Address) (*types.StateAccount, error) {
	acc := &types.StateAccount{
		Nonce:    0,
		Balance:  nil,
		Root:     fakeRoot,
		CodeHash: nil,
	}
	diffAcc := f.accountDiff[address]
	if diffAcc.Nonce != nil {
		acc.Nonce = *diffAcc.Nonce
	} else {
		v, err := f.src.Nonce(address)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve nonce of account %s: %w", address, err)
		}
		acc.Nonce = v
	}
	if diffAcc.Balance != nil {
		acc.Balance = new(uint256.Int).Set(diffAcc.Balance)
	} else {
		v, err := f.src.Balance(address)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve balance of account %s: %w", address, err)
		}
		acc.Balance = new(uint256.Int).Set(v)
	}
	if diffAcc.CodeHash != nil {
		cpy := *diffAcc.CodeHash
		acc.CodeHash = cpy.Bytes()
	} else {
		v, err := f.src.Code(address)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve code of account %s: %w", address, err)
		}
		acc.CodeHash = crypto.Keccak256Hash(v).Bytes()
	}
	return acc, nil
}

func (f *ForkedAccountsTrie) GetStorage(addr common.Address, key []byte) ([]byte, error) {
	k := common.BytesToHash(key)
	diffAcc, ok := f.accountDiff[addr]
	if ok { // if there is a diff, try and see if it contains a storage diff
		v, ok := diffAcc.Storage[k]
		if ok { // if the storage has changed, return that change
			return v.Bytes(), nil
		}
	}
	v, err := f.src.StorageAt(addr, k)
	if err != nil {
		return nil, err
	}
	return v.Bytes(), nil
}

func (f *ForkedAccountsTrie) UpdateAccount(address common.Address, account *types.StateAccount) error {
	nonce := account.Nonce
	b := account.Balance.Clone()
	codeHash := common.BytesToHash(account.CodeHash)
	out := &accountDiff{
		Nonce:    &nonce,
		Balance:  b,
		Storage:  nil,
		CodeHash: &codeHash,
	}
	// preserve the storage diff
	if diffAcc, ok := f.accountDiff[address]; ok {
		out.Storage = diffAcc.Storage
	}
	f.accountDiff[address] = out
	return nil
}

func (f *ForkedAccountsTrie) UpdateStorage(addr common.Address, key, value []byte) error {
	diffAcc, ok := f.accountDiff[addr]
	if !ok {
		return fmt.Errorf("unknown account %s", addr)
	}
	if diffAcc.Storage == nil {
		diffAcc.Storage = make(map[common.Hash]common.Hash)
	}
	k := common.BytesToHash(key)
	v := common.BytesToHash(value)
	diffAcc.Storage[k] = v
	return nil
}

func (f *ForkedAccountsTrie) DeleteAccount(address common.Address) error {
	f.accountDiff[address] = nil
	return nil
}

func (f *ForkedAccountsTrie) DeleteStorage(addr common.Address, key []byte) error {
	return f.UpdateStorage(addr, key, nil)
}

func (f *ForkedAccountsTrie) UpdateContractCode(addr common.Address, codeHash common.Hash, code []byte) error {
	diffAcc, ok := f.accountDiff[addr]
	if !ok {
		return fmt.Errorf("unknown account %s", addr)
	}
	diffAcc.CodeHash = &codeHash
	f.codeDiff[codeHash] = code
	return nil
}

func (f *ForkedAccountsTrie) Hash() common.Hash {
	return f.stateRoot
}

func (f *ForkedAccountsTrie) Commit(collectLeaf bool) (common.Hash, *trienode.NodeSet) {
	panic("cannot commit state-changes of a forked trie")
}

func (f *ForkedAccountsTrie) Witness() map[string]struct{} {
	panic("witness generation of a ForkedAccountsTrie is not supported")
}

func (f *ForkedAccountsTrie) NodeIterator(startKey []byte) (trie.NodeIterator, error) {
	return nil, errors.New("node iteration of a ForkedAccountsTrie is not supported")
}

func (f *ForkedAccountsTrie) Prove(key []byte, proofDb ethdb.KeyValueWriter) error {
	return errors.New("proving of a ForkedAccountsTrie is not supported")
}

func (f *ForkedAccountsTrie) IsVerkle() bool {
	return false
}
