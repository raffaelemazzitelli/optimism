package upgrades

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum-optimism/optimism/op-chain-ops/genesis"
	"github.com/ethereum-optimism/optimism/op-chain-ops/safe"

	"github.com/ethereum-optimism/superchain-registry/superchain"
)

const (
	// upgradeAndCall represents the signature of the upgradeAndCall function
	// on the ProxyAdmin contract.
	upgradeAndCall = "upgradeAndCall(address,address,bytes)"

	method = "setBytes32"
)

var (
	// storageSetterAddr represents the address of the StorageSetter contract.
	storageSetterAddr = common.HexToAddress("0xd81f43eDBCAcb4c29a9bA38a13Ee5d79278270cC")
)

// L1 will add calls for upgrading each of the L1 contracts.
func L1(batch *safe.Batch, addresses superchain.StandardContractAddresses, list superchain.AddressList, config *genesis.DeployConfig, chainConfig *superchain.ChainConfig, superchainConfig *superchain.Superchain, backend bind.ContractBackend) error {
	l1_cross_domain_messenger, err := common.NewMixedcaseAddressFromString(addresses.L1CrossDomainMessenger)
	if err != nil {
		return fmt.Errorf("invalid l1_cross_domain_messenger address: %s", addresses.L1CrossDomainMessenger)
	}
	l1_erc721_bridge, err := common.NewMixedcaseAddressFromString(addresses.L1ERC721Bridge)
	if err != nil {
		return fmt.Errorf("invalid l1_erc721_bridge address: %s", addresses.L1ERC721Bridge)
	}
	l1_standard_bridge, err := common.NewMixedcaseAddressFromString(addresses.L1StandardBridge)
	if err != nil {
		return fmt.Errorf("invalid l1_standard_bridge address: %s", addresses.L1StandardBridge)
	}
	l2_output_oracle, err := common.NewMixedcaseAddressFromString(addresses.L2OutputOracle)
	if err != nil {
		return fmt.Errorf("invalid l2_output_oracle address: %s", addresses.L2OutputOracle)
	}
	optimism_mintable_erc20_factory, err := common.NewMixedcaseAddressFromString(addresses.OptimismMintableERC20Factory)
	if err != nil {
		return fmt.Errorf("invalid optimism_mintable_erc20_factory address: %s", addresses.OptimismMintableERC20Factory)
	}
	optimism_portal, err := common.NewMixedcaseAddressFromString(addresses.OptimismPortal)
	if err != nil {
		return fmt.Errorf("invalid optimism_portal address: %s", addresses.OptimismPortal)
	}
	system_config, err := common.NewMixedcaseAddressFromString(addresses.SystemConfig)
	if err != nil {
		return fmt.Errorf("invalid system_config address: %s", addresses.SystemConfig)
	}

	if err := L1CrossDomainMessenger(batch, l1_cross_domain_messenger.Address(), list, config, chainConfig, superchainConfig, backend); err != nil {
		return fmt.Errorf("upgrading L1CrossDomainMessenger: %w", err)
	}
	if err := L1ERC721Bridge(batch, l1_erc721_bridge.Address(), list, config, chainConfig, superchainConfig, backend); err != nil {
		return fmt.Errorf("upgrading L1ERC721Bridge: %w", err)
	}
	if err := L1StandardBridge(batch, l1_standard_bridge.Address(), list, config, chainConfig, superchainConfig, backend); err != nil {
		return fmt.Errorf("upgrading L1StandardBridge: %w", err)
	}
	if err := L2OutputOracle(batch, l2_output_oracle.Address(), list, config, chainConfig, superchainConfig, backend); err != nil {
		return fmt.Errorf("upgrading L2OutputOracle: %w", err)
	}
	if err := OptimismMintableERC20Factory(batch, optimism_mintable_erc20_factory.Address(), list, config, chainConfig, superchainConfig, backend); err != nil {
		return fmt.Errorf("upgrading OptimismMintableERC20Factory: %w", err)
	}
	if err := OptimismPortal(batch, optimism_portal.Address(), list, config, chainConfig, superchainConfig, backend); err != nil {
		return fmt.Errorf("upgrading OptimismPortal: %w", err)
	}
	if err := SystemConfig(batch, system_config.Address(), list, config, chainConfig, superchainConfig, backend); err != nil {
		return fmt.Errorf("upgrading SystemConfig: %w", err)
	}
	return nil
}

// L1CrossDomainMessenger will add a call to the batch that upgrades the L1CrossDomainMessenger.
func L1CrossDomainMessenger(batch *safe.Batch, address common.Address, list superchain.AddressList, config *genesis.DeployConfig, chainConfig *superchain.ChainConfig, superchainConfig *superchain.Superchain, backend bind.ContractBackend) error {
	proxyAdminABI, err := bindings.ProxyAdminMetaData.GetAbi()
	if err != nil {
		return err
	}

	// 2 Step Upgrade
	{
		storageSetterABI, err := bindings.StorageSetterMetaData.GetAbi()
		if err != nil {
			return err
		}

		input := []bindings.StorageSetterSlot{
			// https://github.com/ethereum-optimism/optimism/blob/86a96023ffd04d119296dff095d02fff79fa15de/packages/contracts-bedrock/.storage-layout#L11-L13
			{
				Key:   common.Hash{},
				Value: common.Hash{},
			},
		}

		calldata, err := storageSetterABI.Pack(method, input)
		if err != nil {
			return err
		}
		args := []any{
			common.Address(list.L1CrossDomainMessengerProxy),
			storageSetterAddr,
			calldata,
		}
		proxyAdmin := common.Address(list.ProxyAdmin)
		if err := batch.AddCall(proxyAdmin, common.Big0, upgradeAndCall, args, proxyAdminABI); err != nil {
			return err
		}
	}

	l1CrossDomainMessengerABI, err := bindings.L1CrossDomainMessengerMetaData.GetAbi()
	if err != nil {
		return err
	}

	l1CrossDomainMessenger, err := bindings.NewL1CrossDomainMessengerCaller(common.Address(list.L1CrossDomainMessengerProxy), backend)
	if err != nil {
		return err
	}
	optimismPortal, err := l1CrossDomainMessenger.PORTAL(&bind.CallOpts{})
	if err != nil {
		return err
	}
	otherMessenger, err := l1CrossDomainMessenger.OTHERMESSENGER(&bind.CallOpts{})
	if err != nil {
		return err
	}

	if optimismPortal != common.Address(list.OptimismPortalProxy) {
		return fmt.Errorf("Portal address doesn't match config")
	}

	if otherMessenger != predeploys.L2CrossDomainMessengerAddr {
		return fmt.Errorf("OtherMessenger address doesn't match config")
	}

	calldata, err := l1CrossDomainMessengerABI.Pack("initialize", common.Address(*superchainConfig.Config.SuperchainConfigAddr), optimismPortal)
	if err != nil {
		return err
	}

	args := []any{
		common.Address(list.L1CrossDomainMessengerProxy),
		common.Address(address),
		calldata,
	}

	proxyAdmin := common.Address(list.ProxyAdmin)
	if err := batch.AddCall(proxyAdmin, common.Big0, upgradeAndCall, args, proxyAdminABI); err != nil {
		return err
	}

	return nil
}

// L1ERC721Bridge will add a call to the batch that upgrades the L1ERC721Bridge.
func L1ERC721Bridge(batch *safe.Batch, address common.Address, list superchain.AddressList, config *genesis.DeployConfig, chainConfig *superchain.ChainConfig, superchainConfig *superchain.Superchain, backend bind.ContractBackend) error {
	proxyAdminABI, err := bindings.ProxyAdminMetaData.GetAbi()
	if err != nil {
		return err
	}

	// 2 Step Upgrade
	{
		storageSetterABI, err := bindings.StorageSetterMetaData.GetAbi()
		if err != nil {
			return err
		}

		input := []bindings.StorageSetterSlot{
			// https://github.com/ethereum-optimism/optimism/blob/86a96023ffd04d119296dff095d02fff79fa15de/packages/contracts-bedrock/.storage-layout#L100-L102
			{
				Key:   common.Hash{},
				Value: common.Hash{},
			},
		}

		calldata, err := storageSetterABI.Pack(method, input)
		if err != nil {
			return fmt.Errorf("setBytes32: %w", err)
		}
		args := []any{
			common.Address(list.L1ERC721BridgeProxy),
			storageSetterAddr,
			calldata,
		}
		proxyAdmin := common.Address(list.ProxyAdmin)
		if err := batch.AddCall(proxyAdmin, common.Big0, upgradeAndCall, args, proxyAdminABI); err != nil {
			return err
		}
	}

	l1ERC721BridgeABI, err := bindings.L1ERC721BridgeMetaData.GetAbi()
	if err != nil {
		return err
	}

	l1ERC721Bridge, err := bindings.NewL1ERC721BridgeCaller(common.Address(list.L1ERC721BridgeProxy), backend)
	if err != nil {
		return err
	}
	messenger, err := l1ERC721Bridge.Messenger(&bind.CallOpts{})
	if err != nil {
		return err
	}
	otherBridge, err := l1ERC721Bridge.OtherBridge(&bind.CallOpts{})
	if err != nil {
		return err
	}

	if messenger != common.Address(list.L1CrossDomainMessengerProxy) {
		return fmt.Errorf("Messenger address doesn't match config")
	}

	if otherBridge != predeploys.L2ERC721BridgeAddr {
		return fmt.Errorf("OtherBridge address doesn't match config")
	}

	calldata, err := l1ERC721BridgeABI.Pack("initialize", messenger, common.Address(*(superchainConfig.Config.SuperchainConfigAddr)))
	if err != nil {
		return err
	}

	args := []any{
		common.Address(list.L1ERC721BridgeProxy),
		common.Address(address),
		calldata,
	}

	proxyAdmin := common.Address(list.ProxyAdmin)
	if err := batch.AddCall(proxyAdmin, common.Big0, upgradeAndCall, args, proxyAdminABI); err != nil {
		return err
	}

	return nil
}

// L1StandardBridge will add a call to the batch that upgrades the L1StandardBridge.
func L1StandardBridge(batch *safe.Batch, address common.Address, list superchain.AddressList, config *genesis.DeployConfig, chainConfig *superchain.ChainConfig, superchainConfig *superchain.Superchain, backend bind.ContractBackend) error {
	proxyAdminABI, err := bindings.ProxyAdminMetaData.GetAbi()
	if err != nil {
		return err
	}

	// 2 Step Upgrade
	{
		storageSetterABI, err := bindings.StorageSetterMetaData.GetAbi()
		if err != nil {
			return err
		}

		input := []bindings.StorageSetterSlot{
			// https://github.com/ethereum-optimism/optimism/blob/86a96023ffd04d119296dff095d02fff79fa15de/packages/contracts-bedrock/.storage-layout#L36-L37
			{
				Key:   common.Hash{},
				Value: common.Hash{},
			},
		}

		calldata, err := storageSetterABI.Pack(method, input)
		if err != nil {
			return err
		}
		args := []any{
			common.Address(list.L1StandardBridgeProxy),
			storageSetterAddr,
			calldata,
		}
		proxyAdmin := common.Address(list.ProxyAdmin)
		if err := batch.AddCall(proxyAdmin, common.Big0, upgradeAndCall, args, proxyAdminABI); err != nil {
			return err
		}
	}

	l1StandardBridgeABI, err := bindings.L1StandardBridgeMetaData.GetAbi()
	if err != nil {
		return err
	}

	l1StandardBridge, err := bindings.NewL1StandardBridgeCaller(common.Address(list.L1StandardBridgeProxy), backend)
	if err != nil {
		return err
	}

	messenger, err := l1StandardBridge.MESSENGER(&bind.CallOpts{})
	if err != nil {
		return err
	}

	otherBridge, err := l1StandardBridge.OTHERBRIDGE(&bind.CallOpts{})
	if err != nil {
		return err
	}

	if messenger != common.Address(list.L1CrossDomainMessengerProxy) {
		return fmt.Errorf("Messenger address doesn't match config")
	}

	if otherBridge != predeploys.L2StandardBridgeAddr {
		return fmt.Errorf("OtherBridge address doesn't match config")
	}

	calldata, err := l1StandardBridgeABI.Pack("initialize", messenger, common.Address(*(superchainConfig.Config.SuperchainConfigAddr)))
	if err != nil {
		return err
	}

	args := []any{
		common.Address(list.L1StandardBridgeProxy),
		common.Address(address),
		calldata,
	}

	proxyAdmin := common.Address(list.ProxyAdmin)
	if err := batch.AddCall(proxyAdmin, common.Big0, upgradeAndCall, args, proxyAdminABI); err != nil {
		return err
	}

	return nil
}

// L2OutputOracle will add a call to the batch that upgrades the L2OutputOracle.
func L2OutputOracle(batch *safe.Batch, address common.Address, list superchain.AddressList, config *genesis.DeployConfig, chainConfig *superchain.ChainConfig, superchainConfig *superchain.Superchain, backend bind.ContractBackend) error {
	proxyAdminABI, err := bindings.ProxyAdminMetaData.GetAbi()
	if err != nil {
		return err
	}

	// 2 Step Upgrade
	{
		storageSetterABI, err := bindings.StorageSetterMetaData.GetAbi()
		if err != nil {
			return err
		}

		input := []bindings.StorageSetterSlot{
			// https://github.com/ethereum-optimism/optimism/blob/86a96023ffd04d119296dff095d02fff79fa15de/packages/contracts-bedrock/.storage-layout#L50-L51
			{
				Key:   common.Hash{},
				Value: common.Hash{},
			},
		}

		calldata, err := storageSetterABI.Pack(method, input)
		if err != nil {
			return err
		}
		args := []any{
			common.Address(list.L2OutputOracleProxy),
			storageSetterAddr,
			calldata,
		}
		proxyAdmin := common.Address(list.ProxyAdmin)
		if err := batch.AddCall(proxyAdmin, common.Big0, upgradeAndCall, args, proxyAdminABI); err != nil {
			return err
		}
	}

	l2OutputOracleABI, err := bindings.L2OutputOracleMetaData.GetAbi()
	if err != nil {
		return err
	}

	l2OutputOracle, err := bindings.NewL2OutputOracleCaller(common.Address(list.L2OutputOracleProxy), backend)
	if err != nil {
		return err
	}

	l2OutputOracleSubmissionInterval, err := l2OutputOracle.SUBMISSIONINTERVAL(&bind.CallOpts{})
	if err != nil {
		return err
	}

	l2BlockTime, err := l2OutputOracle.L2BLOCKTIME(&bind.CallOpts{})
	if err != nil {
		return err
	}

	l2OutputOracleStartingBlockNumber, err := l2OutputOracle.StartingBlockNumber(&bind.CallOpts{})
	if err != nil {
		return err
	}

	l2OutputOracleStartingTimestamp, err := l2OutputOracle.StartingTimestamp(&bind.CallOpts{})
	if err != nil {
		return err
	}

	l2OutputOracleProposer, err := l2OutputOracle.PROPOSER(&bind.CallOpts{})
	if err != nil {
		return err
	}

	l2OutputOracleChallenger, err := l2OutputOracle.CHALLENGER(&bind.CallOpts{})
	if err != nil {
		return err
	}

	finalizationPeriodSeconds, err := l2OutputOracle.FINALIZATIONPERIODSECONDS(&bind.CallOpts{})
	if err != nil {
		return err
	}

	if config != nil {
		if l2OutputOracleSubmissionInterval.Uint64() != config.L2OutputOracleSubmissionInterval {
			return fmt.Errorf("L2OutputOracleSubmissionInterval address doesn't match config")
		}

		if l2BlockTime.Uint64() != config.L2BlockTime {
			return fmt.Errorf("L2BlockTime address doesn't match config")
		}

		if l2OutputOracleStartingBlockNumber.Uint64() != config.L2OutputOracleStartingBlockNumber {
			return fmt.Errorf("L2OutputOracleStartingBlockNumber address doesn't match config")
		}

		if config.L2OutputOracleStartingTimestamp < 0 {
			return fmt.Errorf("L2OutputOracleStartingTimestamp must be concrete")
		}

		if int(l2OutputOracleStartingTimestamp.Int64()) != config.L2OutputOracleStartingTimestamp {
			return fmt.Errorf("L2OutputOracleStartingTimestamp address doesn't match config")
		}

		if l2OutputOracleProposer != config.L2OutputOracleProposer {
			return fmt.Errorf("L2OutputOracleProposer address doesn't match config")
		}

		if l2OutputOracleChallenger != config.L2OutputOracleChallenger {
			return fmt.Errorf("L2OutputOracleChallenger address doesn't match config")
		}

		if finalizationPeriodSeconds.Uint64() != config.FinalizationPeriodSeconds {
			return fmt.Errorf("FinalizationPeriodSeconds address doesn't match config")
		}
	}

	calldata, err := l2OutputOracleABI.Pack(
		"initialize",
		l2OutputOracleSubmissionInterval,
		l2BlockTime,
		l2OutputOracleStartingBlockNumber,
		l2OutputOracleStartingTimestamp,
		l2OutputOracleProposer,
		l2OutputOracleChallenger,
		finalizationPeriodSeconds,
	)
	if err != nil {
		return err
	}

	args := []any{
		common.Address(list.L2OutputOracleProxy),
		common.Address(address),
		calldata,
	}

	proxyAdmin := common.Address(list.ProxyAdmin)
	if err := batch.AddCall(proxyAdmin, common.Big0, upgradeAndCall, args, proxyAdminABI); err != nil {
		return err
	}

	return nil
}

// OptimismMintableERC20Factory will add a call to the batch that upgrades the OptimismMintableERC20Factory.
func OptimismMintableERC20Factory(batch *safe.Batch, address common.Address, list superchain.AddressList, config *genesis.DeployConfig, chainConfig *superchain.ChainConfig, superchainConfig *superchain.Superchain, backend bind.ContractBackend) error {
	proxyAdminABI, err := bindings.ProxyAdminMetaData.GetAbi()
	if err != nil {
		return err
	}

	// 2 Step Upgrade
	{
		storageSetterABI, err := bindings.StorageSetterMetaData.GetAbi()
		if err != nil {
			return err
		}

		input := []bindings.StorageSetterSlot{
			// https://github.com/ethereum-optimism/optimism/blob/86a96023ffd04d119296dff095d02fff79fa15de/packages/contracts-bedrock/.storage-layout#L287-L289
			{
				Key:   common.Hash{},
				Value: common.Hash{},
			},
		}

		calldata, err := storageSetterABI.Pack(method, input)
		if err != nil {
			return err
		}
		args := []any{
			common.Address(list.OptimismMintableERC20FactoryProxy),
			storageSetterAddr,
			calldata,
		}
		proxyAdmin := common.Address(list.ProxyAdmin)
		if err := batch.AddCall(proxyAdmin, common.Big0, upgradeAndCall, args, proxyAdminABI); err != nil {
			return err
		}
	}

	optimismMintableERC20FactoryABI, err := bindings.OptimismMintableERC20FactoryMetaData.GetAbi()
	if err != nil {
		return err
	}

	optimismMintableERC20Factory, err := bindings.NewOptimismMintableERC20FactoryCaller(common.Address(list.OptimismMintableERC20FactoryProxy), backend)
	if err != nil {
		return err
	}

	bridge, err := optimismMintableERC20Factory.BRIDGE(&bind.CallOpts{})
	if err != nil {
		return err
	}

	if bridge != common.Address(list.L1StandardBridgeProxy) {
		return fmt.Errorf("Bridge address doesn't match config")
	}

	calldata, err := optimismMintableERC20FactoryABI.Pack("initialize", bridge)
	if err != nil {
		return err
	}

	args := []any{
		common.Address(list.OptimismMintableERC20FactoryProxy),
		common.Address(address),
		calldata,
	}

	proxyAdmin := common.Address(list.ProxyAdmin)
	if err := batch.AddCall(proxyAdmin, common.Big0, upgradeAndCall, args, proxyAdminABI); err != nil {
		return err
	}

	return nil
}

// OptimismPortal will add a call to the batch that upgrades the OptimismPortal.
func OptimismPortal(batch *safe.Batch, address common.Address, list superchain.AddressList, config *genesis.DeployConfig, chainConfig *superchain.ChainConfig, superchainConfig *superchain.Superchain, backend bind.ContractBackend) error {
	proxyAdminABI, err := bindings.ProxyAdminMetaData.GetAbi()
	if err != nil {
		return err
	}

	// 2 Step Upgrade
	{
		storageSetterABI, err := bindings.StorageSetterMetaData.GetAbi()
		if err != nil {
			return err
		}

		input := []bindings.StorageSetterSlot{
			// https://github.com/ethereum-optimism/optimism/blob/86a96023ffd04d119296dff095d02fff79fa15de/packages/contracts-bedrock/.storage-layout#L64-L65
			{
				Key:   common.Hash{},
				Value: common.Hash{},
			},
		}

		calldata, err := storageSetterABI.Pack(method, input)
		if err != nil {
			return err
		}
		args := []any{
			common.Address(list.OptimismPortalProxy),
			storageSetterAddr,
			calldata,
		}
		proxyAdmin := common.Address(list.ProxyAdmin)
		if err := batch.AddCall(proxyAdmin, common.Big0, upgradeAndCall, args, proxyAdminABI); err != nil {
			return err
		}
	}

	optimismPortalABI, err := bindings.OptimismPortalMetaData.GetAbi()
	if err != nil {
		return err
	}

	optimismPortal, err := bindings.NewOptimismPortalCaller(common.Address(list.OptimismPortalProxy), backend)
	if err != nil {
		return err
	}

	l2OutputOracle, err := optimismPortal.L2Oracle(&bind.CallOpts{})
	if err != nil {
		// Handle legacy `L2_ORACLE()(address)` getter
		addr := common.Address(list.OptimismPortalProxy)
		msg := ethereum.CallMsg{
			To:   &addr,
			Data: common.FromHex("0x001c2ff6"),
		}
		data, err := backend.CallContract(context.Background(), msg, nil)
		if err != nil {
			return err
		}
		if len(data) != 32 {
			return fmt.Errorf("unexpected result length: %d", len(data))
		}
		l2OutputOracle = common.BytesToAddress(data[12:])
	}

	systemConfig, err := optimismPortal.SystemConfig(&bind.CallOpts{})
	if err != nil {
		// Handle legacy `SYSTEM_CONFIG()(address)` getter
		addr := common.Address(list.OptimismPortalProxy)
		msg := ethereum.CallMsg{
			To:   &addr,
			Data: common.FromHex("0xf0498750"),
		}
		data, err := backend.CallContract(context.Background(), msg, nil)
		if err != nil {
			return err
		}
		if len(data) != 32 {
			return fmt.Errorf("unexpected result length: %d", len(data))
		}
		systemConfig = common.BytesToAddress(data[12:])
	}

	if l2OutputOracle != common.Address(list.L2OutputOracleProxy) {
		return fmt.Errorf("L2OutputOracle address doesn't match config")
	}

	if systemConfig != common.Address(list.SystemConfigProxy) {
		return fmt.Errorf("SystemConfig address doesn't match config")
	}

	calldata, err := optimismPortalABI.Pack("initialize", l2OutputOracle, systemConfig, common.Address(*superchainConfig.Config.SuperchainConfigAddr))
	if err != nil {
		return err
	}

	args := []any{
		common.Address(list.OptimismPortalProxy),
		common.Address(address),
		calldata,
	}

	proxyAdmin := common.Address(list.ProxyAdmin)
	if err := batch.AddCall(proxyAdmin, common.Big0, upgradeAndCall, args, proxyAdminABI); err != nil {
		return err
	}

	return nil
}

// SystemConfig will add a call to the batch that upgrades the SystemConfig.
func SystemConfig(batch *safe.Batch, address common.Address, list superchain.AddressList, config *genesis.DeployConfig, chainConfig *superchain.ChainConfig, superchainConfig *superchain.Superchain, backend bind.ContractBackend) error {
	proxyAdminABI, err := bindings.ProxyAdminMetaData.GetAbi()
	if err != nil {
		return err
	}

	// 2 Step Upgrade
	{
		storageSetterABI, err := bindings.StorageSetterMetaData.GetAbi()
		if err != nil {
			return err
		}

		var startBlock common.Hash
		if config != nil {
			startBlock = common.BigToHash(new(big.Int).SetUint64(config.SystemConfigStartBlock))
		} else {
			val, err := strconv.ParseUint(os.Getenv("SYSTEM_CONFIG_START_BLOCK"), 10, 64)
			if err != nil {
				return err
			}
			startBlock = common.BigToHash(new(big.Int).SetUint64(val))
		}

		input := []bindings.StorageSetterSlot{
			// https://github.com/ethereum-optimism/optimism/blob/86a96023ffd04d119296dff095d02fff79fa15de/packages/contracts-bedrock/.storage-layout#L82-L83
			{
				Key:   common.Hash{},
				Value: common.Hash{},
			},
			// bytes32 public constant START_BLOCK_SLOT = bytes32(uint256(keccak256("systemconfig.startBlock")) - 1);
			{
				Key:   common.HexToHash("0xa11ee3ab75b40e88a0105e935d17cd36c8faee0138320d776c411291bdbbb19f"),
				Value: startBlock,
			},
		}

		calldata, err := storageSetterABI.Pack(method, input)
		if err != nil {
			return err
		}
		args := []any{
			common.Address(list.SystemConfigProxy),
			storageSetterAddr,
			calldata,
		}
		proxyAdmin := common.Address(list.ProxyAdmin)
		if err := batch.AddCall(proxyAdmin, common.Big0, upgradeAndCall, args, proxyAdminABI); err != nil {
			return err
		}
	}

	systemConfigABI, err := bindings.SystemConfigMetaData.GetAbi()
	if err != nil {
		return err
	}

	systemConfig, err := bindings.NewSystemConfigCaller(common.Address(list.SystemConfigProxy), backend)
	if err != nil {
		return err
	}

	gasPriceOracleOverhead, err := systemConfig.Overhead(&bind.CallOpts{})
	if err != nil {
		return err
	}

	gasPriceOracleScalar, err := systemConfig.Scalar(&bind.CallOpts{})
	if err != nil {
		return err
	}

	batcherHash, err := systemConfig.BatcherHash(&bind.CallOpts{})
	if err != nil {
		return err
	}

	l2GenesisBlockGasLimit, err := systemConfig.GasLimit(&bind.CallOpts{})
	if err != nil {
		return err
	}

	p2pSequencerAddress, err := systemConfig.UnsafeBlockSigner(&bind.CallOpts{})
	if err != nil {
		return err
	}

	finalSystemOwner, err := systemConfig.Owner(&bind.CallOpts{})
	if err != nil {
		return err
	}

	if config != nil {
		if gasPriceOracleOverhead.Uint64() != config.GasPriceOracleOverhead {
			return fmt.Errorf("GasPriceOracleOverhead address doesn't match config")
		}
		if gasPriceOracleScalar.Uint64() != config.GasPriceOracleScalar {
			return fmt.Errorf("GasPriceOracleScalar address doesn't match config")
		}
		if batcherHash != common.BytesToHash(config.BatchSenderAddress.Bytes()) {
			return fmt.Errorf("BatchSenderAddress address doesn't match config")
		}
		if l2GenesisBlockGasLimit != uint64(config.L2GenesisBlockGasLimit) {
			return fmt.Errorf("L2GenesisBlockGasLimit address doesn't match config")
		}
		if p2pSequencerAddress != config.P2PSequencerAddress {
			return fmt.Errorf("P2PSequencerAddress address doesn't match config")
		}
		if finalSystemOwner != config.FinalSystemOwner {
			return fmt.Errorf("FinalSystemOwner address doesn't match config")
		}
	}

	resourceConfig, err := systemConfig.ResourceConfig(&bind.CallOpts{})
	if err != nil {
		return err
	}

	if resourceConfig.MaxResourceLimit != genesis.DefaultResourceConfig.MaxResourceLimit {
		return fmt.Errorf("DefaultResourceConfig MaxResourceLimit doesn't match contract MaxResourceLimit")
	}
	if resourceConfig.ElasticityMultiplier != genesis.DefaultResourceConfig.ElasticityMultiplier {
		return fmt.Errorf("DefaultResourceConfig ElasticityMultiplier doesn't match contract ElasticityMultiplier")
	}
	if resourceConfig.BaseFeeMaxChangeDenominator != genesis.DefaultResourceConfig.BaseFeeMaxChangeDenominator {
		return fmt.Errorf("DefaultResourceConfig BaseFeeMaxChangeDenominator doesn't match contract BaseFeeMaxChangeDenominator")
	}
	if resourceConfig.MinimumBaseFee != genesis.DefaultResourceConfig.MinimumBaseFee {
		return fmt.Errorf("DefaultResourceConfig MinimumBaseFee doesn't match contract MinimumBaseFee")
	}
	if resourceConfig.SystemTxMaxGas != genesis.DefaultResourceConfig.SystemTxMaxGas {
		return fmt.Errorf("DefaultResourceConfig SystemTxMaxGas doesn't match contract SystemTxMaxGas")
	}
	if resourceConfig.MaximumBaseFee.Cmp(genesis.DefaultResourceConfig.MaximumBaseFee) != 0 {
		return fmt.Errorf("DefaultResourceConfig MaximumBaseFee doesn't match contract MaximumBaseFee")
	}

	calldata, err := systemConfigABI.Pack(
		"initialize",
		finalSystemOwner,
		gasPriceOracleOverhead,
		gasPriceOracleScalar,
		batcherHash,
		l2GenesisBlockGasLimit,
		p2pSequencerAddress,
		genesis.DefaultResourceConfig,
		chainConfig.BatchInboxAddr,
		bindings.SystemConfigAddresses{
			L1CrossDomainMessenger:       common.Address(list.L1CrossDomainMessengerProxy),
			L1ERC721Bridge:               common.Address(list.L1ERC721BridgeProxy),
			L1StandardBridge:             common.Address(list.L1StandardBridgeProxy),
			L2OutputOracle:               common.Address(list.L2OutputOracleProxy),
			OptimismPortal:               common.Address(list.OptimismPortalProxy),
			OptimismMintableERC20Factory: common.Address(list.OptimismMintableERC20FactoryProxy),
		},
	)
	if err != nil {
		return err
	}

	args := []any{
		common.Address(list.SystemConfigProxy),
		common.Address(address),
		calldata,
	}

	proxyAdmin := common.Address(list.ProxyAdmin)
	if err := batch.AddCall(proxyAdmin, common.Big0, upgradeAndCall, args, proxyAdminABI); err != nil {
		return err
	}

	return nil
}
