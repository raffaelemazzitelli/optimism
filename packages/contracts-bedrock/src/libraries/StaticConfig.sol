// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { IFeeVault } from "src/L2/interfaces/IFeeVault.sol";
import { Encoding } from "src/libraries/Encoding.sol";

/// @notice Enum representing different types of configurations that can be set on L1BlockIsthmus.
/// @custom:value SET_GAS_PAYING_TOKEN  Represents the config type for setting the gas paying token.
/// @custom:value
/// @custom:value ADD_DEPENDENCY        Represents the config type for adding a chain to the interop dependency set.
/// @custom:value REMOVE_DEPENDENCY     Represents the config type for removing a chain from the interop dependency set.
enum ConfigType {
    SET_GAS_PAYING_TOKEN,
    SET_BASE_FEE_VAULT_CONFIG,
    SET_L1_FEE_VAULT_CONFIG,
    SET_SEQUENCER_FEE_VAULT_CONFIG,
    SET_L1_CROSS_DOMAIN_MESSENGER_ADDRESS,
    SET_L1_ERC_721_BRIDGE_ADDRESS,
    SET_L1_STANDARD_BRIDGE_ADDRESS,
    SET_REMOTE_CHAIN_ID,
    ADD_DEPENDENCY,
    REMOVE_DEPENDENCY
}

// The main benefit of this library is to give the reader the ability to observe that
// the encode decode logic is correct by being able to see it right next to each other.
// Otherwise the encode/decode will exist in many different locations, between mocks and
// runtime code. It is overly verbose and I don't like it.

/// @title StaticConfig
/// @notice Library for encoding and decoding static configuration data.
library StaticConfig {
    /// @notice Encodes the static configuration data for setting a gas paying token.
    /// @param _token    Address of the gas paying token.
    /// @param _decimals Number of decimals for the gas paying token.
    /// @param _name     Name of the gas paying token.
    /// @param _symbol   Symbol of the gas paying token.
    /// @return Encoded static configuration data.
    function encodeSetGasPayingToken(
        address _token,
        uint8 _decimals,
        bytes32 _name,
        bytes32 _symbol
    )
        internal
        pure
        returns (bytes memory)
    {
        return abi.encode(_token, _decimals, _name, _symbol);
    }

    /// @notice Decodes the static configuration data for setting a gas paying token.
    /// @param _data Encoded static configuration data.
    /// @return Decoded gas paying token data (token address, decimals, name, symbol).
    function decodeSetGasPayingToken(bytes memory _data) internal pure returns (address, uint8, bytes32, bytes32) {
        return abi.decode(_data, (address, uint8, bytes32, bytes32));
    }

    /// @notice Encodes the static configuration data for adding a dependency.
    /// @param _chainId Chain ID of the dependency to add.
    /// @return Encoded static configuration data.
    function encodeAddDependency(uint256 _chainId) internal pure returns (bytes memory) {
        return abi.encode(_chainId);
    }

    /// @notice Decodes the static configuration data for adding a dependency.
    /// @param _data Encoded static configuration data.
    /// @return Decoded chain ID of the dependency to add.
    function decodeAddDependency(bytes memory _data) internal pure returns (uint256) {
        return abi.decode(_data, (uint256));
    }

    /// @notice Encodes the static configuration data for removing a dependency.
    /// @param _chainId Chain ID of the dependency to remove.
    /// @return Encoded static configuration data.
    function encodeRemoveDependency(uint256 _chainId) internal pure returns (bytes memory) {
        return abi.encode(_chainId);
    }

    /// @notice Decodes the static configuration data for removing a dependency.
    /// @param _data Encoded static configuration data.
    /// @return Decoded chain ID of the dependency to remove.
    function decodeRemoveDependency(bytes memory _data) internal pure returns (uint256) {
        return abi.decode(_data, (uint256));
    }

    /// @notice TODO: deprecate this in favor of inline abi encode/decode
    // This will not work
    // SystemConfig takes 3 params -> bytes32 -> bytes -> deposit
    // L1Block -> bytes -> bytes32
    function encodeSetFeeVaultConfig(address _recipient, uint256 _min, IFeeVault.WithdrawalNetwork _network) internal pure returns (bytes memory) {
        bytes32 encoded = Encoding.encodeFeeVaultConfig(_recipient, _min, _network);
        return abi.encode(encoded);
    }

    /// @notice
    /// This is not being used so it should be deleted
    function decodeSetFeeVaultConfig(bytes memory _data) internal pure returns (address, uint256, IFeeVault.WithdrawalNetwork) {
        bytes32 encoded = abi.decode(_data, (bytes32));
        return Encoding.decodeFeeVaultConfig(encoded);
    }

    /// @notice
    function encodeSetAddress(address _address) internal pure returns (bytes memory) {
        return abi.encode(_address);
    }

    /// @notice
    function decodeSetAddress(bytes memory _data) internal pure returns (address) {
        return abi.decode(_data, (address));
    }

    /// @notice
    function encodeSetRemoteChainId(uint256 _chainId) internal pure returns (bytes memory) {
        return abi.encode(_chainId);
    }

    /// @notice
    function decodeSetRemoteChainId(bytes memory _data) internal pure returns (uint256) {
        return abi.decode(_data, (uint256));
    }
}
