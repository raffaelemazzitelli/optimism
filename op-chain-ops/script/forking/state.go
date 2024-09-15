package forking

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/stateless"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/trie/utils"
	"github.com/holiman/uint256"
)

type forkStateEntry struct {
	state *state.StateDB
}

func (fe *forkStateEntry) DB() *ForkDB {
	fe.state.GetTrie()
	return fe.state.Database().(*ForkDB)
}

// ForkableState implements the vm.StateDB interface,
// and a few other methods as defined in the VMStateDB interface.
// This state can be forked in-place,
// swapping over operations to route to in-memory states that wrap fork sources.
type ForkableState struct {
	selected VMStateDB

	activeFork ForkID
	forks      map[ForkID]*forkStateEntry

	// persistent accounts will override any interactions
	// to be directly with the fallback state, rather than whatever fork is active
	persistent map[common.Address]struct{}

	fallback VMStateDB

	idCounter uint64
}

var _ VMStateDB = (*ForkableState)(nil)

func NewForkableState(base VMStateDB) *ForkableState {
	return &ForkableState{
		selected:   base,
		activeFork: ForkID{},
		forks:      make(map[ForkID]*forkStateEntry),
		persistent: make(map[common.Address]struct{}),
		fallback:   base,
		idCounter:  0,
	}
}

// CreateSelectFork is like vm.createSelectFork, it creates a fork, and selects it immediately.
func (fst *ForkableState) CreateSelectFork(source ForkSource) (ForkID, error) {
	id, err := fst.CreateFork(source)
	if err != nil {
		return id, err
	}
	return id, fst.SelectFork(id)
}

// CreateFork is like vm.createFork, it creates a fork, but does not select it yet.
func (fst *ForkableState) CreateFork(source ForkSource) (ForkID, error) {
	fst.idCounter += 1 // increment first, don't use ID 0
	id := ForkID(*uint256.NewInt(fst.idCounter))
	_, ok := fst.forks[id]
	if ok { // sanity check our ID counter is consistent with the tracked forks
		return id, fmt.Errorf("cannot create fork, fork %q already exists", id)
	}
	forkDB := NewForkDB(source)
	st, err := state.New(forkDB.active.stateRoot, forkDB, nil)
	if err != nil {
		return id, fmt.Errorf("failed to construct fork state: %w", err)
	}
	fst.forks[id] = &forkStateEntry{
		state: st,
	}
	return id, nil
}

// SelectFork is like vm.selectFork, it activates the usage of a previously created fork.
func (fst *ForkableState) SelectFork(id ForkID) error {
	if id == (ForkID{}) {
		fst.selected = fst.fallback
		fst.activeFork = ForkID{}
		return nil
	}
	f, ok := fst.forks[id]
	if !ok {
		return fmt.Errorf("cannot select fork, fork %q is unknown", id)
	}
	fst.selected = f.state
	fst.activeFork = id
	return nil
}

// ActiveFork returns the ID current active fork, or active == false if no fork is active.
func (fst *ForkableState) ActiveFork() (id ForkID, active bool) {
	return fst.activeFork, fst.activeFork != (ForkID{})
}

// ForkURLOrAlias returns the URL or alias that the fork was configured with as source.
func (fst *ForkableState) ForkURLOrAlias(id ForkID) (string, error) {
	f, ok := fst.forks[id]
	if !ok {
		return "", fmt.Errorf("unknown fork %q", id)
	}
	return f.DB().active.src.URLOrAlias(), nil
}

// ChangeFork changes the fork with the given ID to anchor to the given source.
// This can be used to roll the fork to a different block,
// by providing a source with a state anchored at the new block.
// To roll to an accurate specific intra-block tx state,
// the transactions should be replayed on top of the start of the block.
func (fst *ForkableState) ChangeFork(id ForkID, source ForkSource, keepChanges bool) error {
	if id == (ForkID{}) {
		return errors.New("cannot change non-fork state")
	}
	forkTrie := fst.forks[id].DB().active
	forkTrie.src = source
	forkTrie.stateRoot = source.StateRoot()
	if !keepChanges {
		forkTrie.ClearDiff()
	}
	return nil
}

// SubstituteBaseState substitutes in a fallback state.
func (fst *ForkableState) SubstituteBaseState(base VMStateDB) {
	fst.fallback = base
}

// MakePersistent is like vm.makePersistent, it maintains this account context across all forks.
// It does not make the account of a fork persistent, it makes an account override what might be in a fork.
func (fst *ForkableState) MakePersistent(addr common.Address) {
	fst.persistent[addr] = struct{}{}
}

// RevokePersistent is like vm.revokePersistent, it undoes a previous vm.makePersistent.
func (fst *ForkableState) RevokePersistent(addr common.Address) {
	delete(fst.persistent, addr)
}

// IsPersistent is like vm.isPersistent, it checks if an account persists across forks.
func (fst *ForkableState) IsPersistent(addr common.Address) bool {
	_, ok := fst.persistent[addr]
	return ok
}

func (fst *ForkableState) stateFor(addr common.Address) VMStateDB {
	// if not forked, shortcut to the selected state, avoid map lookup
	if fst.activeFork == (ForkID{}) {
		return fst.fallback
	}
	// if forked, check if we persisted this account across forks
	_, ok := fst.persistent[addr]
	if ok {
		return fst.fallback
	}
	return fst.selected
}

// Finalise finalises the state by removing the destructed objects and clears
// the journal as well as the refunds. Finalise, however, will not push any updates
// into the tries just yet.
//
// The changes will be flushed to the underlying DB.
// A *ForkDB if the state is currently forked.
func (fst *ForkableState) Finalise(deleteEmptyObjects bool) {
	fst.selected.Finalise(deleteEmptyObjects)
}

func (fst *ForkableState) CreateAccount(address common.Address) {
	fst.stateFor(address).CreateAccount(address)
}

func (fst *ForkableState) CreateContract(address common.Address) {
	fst.stateFor(address).CreateContract(address)
}

func (fst *ForkableState) SubBalance(address common.Address, u *uint256.Int, reason tracing.BalanceChangeReason) {
	fst.stateFor(address).SubBalance(address, u, reason)
}

func (fst *ForkableState) AddBalance(address common.Address, u *uint256.Int, reason tracing.BalanceChangeReason) {
	fst.stateFor(address).AddBalance(address, u, reason)
}

func (fst *ForkableState) GetBalance(address common.Address) *uint256.Int {
	return fst.stateFor(address).GetBalance(address)
}

func (fst *ForkableState) GetNonce(address common.Address) uint64 {
	return fst.stateFor(address).GetNonce(address)
}

func (fst *ForkableState) SetNonce(address common.Address, u uint64) {
	fst.stateFor(address).SetNonce(address, u)
}

func (fst *ForkableState) GetCodeHash(address common.Address) common.Hash {
	return fst.stateFor(address).GetCodeHash(address)
}

func (fst *ForkableState) GetCode(address common.Address) []byte {
	return fst.stateFor(address).GetCode(address)
}

func (fst *ForkableState) SetCode(address common.Address, bytes []byte) {
	fst.stateFor(address).SetCode(address, bytes)
}

func (fst *ForkableState) GetCodeSize(address common.Address) int {
	return fst.stateFor(address).GetCodeSize(address)
}

func (fst *ForkableState) AddRefund(u uint64) {
	fst.selected.AddRefund(u)
}

func (fst *ForkableState) SubRefund(u uint64) {
	fst.selected.SubRefund(u)
}

func (fst *ForkableState) GetRefund() uint64 {
	return fst.selected.GetRefund()
}

func (fst *ForkableState) GetCommittedState(address common.Address, hash common.Hash) common.Hash {
	return fst.stateFor(address).GetCommittedState(address, hash)
}

func (fst *ForkableState) GetState(address common.Address, k common.Hash) common.Hash {
	return fst.stateFor(address).GetState(address, k)
}

func (fst *ForkableState) SetState(address common.Address, k common.Hash, v common.Hash) {
	fst.stateFor(address).SetState(address, k, v)
}

func (fst *ForkableState) GetStorageRoot(addr common.Address) common.Hash {
	return fst.stateFor(addr).GetStorageRoot(addr)
}

func (fst *ForkableState) GetTransientState(addr common.Address, key common.Hash) common.Hash {
	return fst.stateFor(addr).GetTransientState(addr, key)
}

func (fst *ForkableState) SetTransientState(addr common.Address, key, value common.Hash) {
	fst.stateFor(addr).SetTransientState(addr, key, value)
}

func (fst *ForkableState) SelfDestruct(address common.Address) {
	fst.stateFor(address).SelfDestruct(address)
}

func (fst *ForkableState) HasSelfDestructed(address common.Address) bool {
	return fst.stateFor(address).HasSelfDestructed(address)
}

func (fst *ForkableState) Selfdestruct6780(address common.Address) {
	fst.stateFor(address).Selfdestruct6780(address)
}

func (fst *ForkableState) Exist(address common.Address) bool {
	return fst.stateFor(address).Exist(address)
}

func (fst *ForkableState) Empty(address common.Address) bool {
	return fst.stateFor(address).Empty(address)
}

func (fst *ForkableState) AddressInAccessList(addr common.Address) bool {
	return fst.stateFor(addr).AddressInAccessList(addr)
}

func (fst *ForkableState) SlotInAccessList(addr common.Address, slot common.Hash) (addressOk bool, slotOk bool) {
	return fst.stateFor(addr).SlotInAccessList(addr, slot)
}

func (fst *ForkableState) AddAddressToAccessList(addr common.Address) {
	fst.stateFor(addr).AddAddressToAccessList(addr)
}

func (fst *ForkableState) AddSlotToAccessList(addr common.Address, slot common.Hash) {
	fst.stateFor(addr).AddSlotToAccessList(addr, slot)
}

func (fst *ForkableState) PointCache() *utils.PointCache {
	return fst.selected.PointCache()
}

func (fst *ForkableState) Prepare(rules params.Rules, sender, coinbase common.Address, dest *common.Address, precompiles []common.Address, txAccesses types.AccessList) {
	fst.selected.Prepare(rules, sender, coinbase, dest, precompiles, txAccesses)
}

func (fst *ForkableState) RevertToSnapshot(i int) {
	fst.selected.RevertToSnapshot(i)
}

func (fst *ForkableState) Snapshot() int {
	return fst.selected.Snapshot()
}

func (fst *ForkableState) AddLog(log *types.Log) {
	fst.selected.AddLog(log)
}

func (fst *ForkableState) AddPreimage(hash common.Hash, img []byte) {
	fst.selected.AddPreimage(hash, img)
}

func (fst *ForkableState) Witness() *stateless.Witness {
	return fst.selected.Witness()
}

func (fst *ForkableState) SetBalance(addr common.Address, amount *uint256.Int, reason tracing.BalanceChangeReason) {
	fst.stateFor(addr).SetBalance(addr, amount, reason)
}
