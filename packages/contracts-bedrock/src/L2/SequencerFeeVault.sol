// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

// Contracts
import { FeeVault } from "src/universal/FeeVault.sol";

// Libraries
import { ISemver } from "src/universal/interfaces/ISemver.sol";

/// @custom:proxied true
/// @custom:predeploy 0x4200000000000000000000000000000000000011
/// @title SequencerFeeVault
/// @notice The SequencerFeeVault is the contract that holds any fees paid to the Sequencer during
///         transaction processing and block production.
contract SequencerFeeVault is FeeVault, ISemver {
    /// @custom:semver 1.5.0-beta.3
    string public constant version = "1.5.0-beta.3";

    /// @notice Constructs the SequencerFeeVault contract.
    /// @param _recipient           Wallet that will receive the fees.
    /// @param _minWithdrawalAmount Minimum balance for withdrawals.
    /// @param _withdrawalNetwork   Network which the recipient will receive fees on.
    constructor(
        address _recipient,
        uint256 _minWithdrawalAmount,
        WithdrawalNetwork _withdrawalNetwork
    )
        FeeVault(_recipient, _minWithdrawalAmount, _withdrawalNetwork)
    { }

    /// @custom:legacy
    /// @notice Legacy getter for the recipient address.
    /// @return The recipient address.
    function l1FeeWallet() public view returns (address) {
        return RECIPIENT;
    }
}
