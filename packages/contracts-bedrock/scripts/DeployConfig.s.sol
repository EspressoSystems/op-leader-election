// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { Script } from "forge-std/Script.sol";
import { console2 as console } from "forge-std/console2.sol";
import { stdJson } from "forge-std/StdJson.sol";
import { Executables } from "./Executables.sol";
import { Chains } from "./Chains.sol";

/// @title DeployConfig
/// @notice Represents the configuration required to deploy the system. It is expected
///         to read the file from JSON. A future improvement would be to have fallback
///         values if they are not defined in the JSON themselves.
contract DeployConfig is Script {
    string internal _json;

    address public finalSystemOwner;
    address public portalGuardian;
    uint256 public l1ChainID;
    uint256 public l2ChainID;
    uint256 public l2BlockTime;
    uint256 public maxSequencerDrift;
    uint256 public sequencerWindowSize;
    uint256 public channelTimeout;
    address public p2pSequencerAddress;
    address public batchInboxAddress;
    address public batchSenderAddress;
    uint256 public l2OutputOracleSubmissionInterval;
    int256 internal _l2OutputOracleStartingTimestamp;
    uint256 public l2OutputOracleStartingBlockNumber;
    address public l2OutputOracleProposer;
    address public l2OutputOracleChallenger;
    uint256 public finalizationPeriodSeconds;
    address public proxyAdminOwner;
    address public baseFeeVaultRecipient;
    address public l1FeeVaultRecipient;
    address public sequencerFeeVaultRecipient;
    string public governanceTokenName;
    string public governanceTokenSymbol;
    address public governanceTokenOwner;
    uint256 public l2GenesisBlockGasLimit;
    uint256 public l2GenesisBlockBaseFeePerGas;
    uint256 public gasPriceOracleOverhead;
    uint256 public gasPriceOracleScalar;
    uint256 public eip1559Denominator;
    uint256 public eip1559Elasticity;
    uint256 public faultGameAbsolutePrestate;
    uint256 public faultGameMaxDepth;
    uint256 public faultGameMaxDuration;
    uint256 public systemConfigStartBlock;
    uint256 public leaderElectionNumberOfLeaders;
    uint256 public leaderElectionNumberOfSlotsPerLeader;
    address public batchInboxContractAddress;
    uint256 public initialBatcherVersion;
    uint256 public requiredProtocolVersion;
    uint256 public recommendedProtocolVersion;

    constructor(string memory _path) {
        console.log("DeployConfig: reading file %s", _path);
        try vm.readFile(_path) returns (string memory data) {
            _json = data;
        } catch {
            console.log("Warning: unable to read config. Do not deploy unless you are not using config.");
            return;
        }

        finalSystemOwner = stdJson.readAddress(_json, "$.finalSystemOwner");
        portalGuardian = stdJson.readAddress(_json, "$.portalGuardian");
        l1ChainID = stdJson.readUint(_json, "$.l1ChainID");
        l2ChainID = stdJson.readUint(_json, "$.l2ChainID");
        l2BlockTime = stdJson.readUint(_json, "$.l2BlockTime");
        maxSequencerDrift = stdJson.readUint(_json, "$.maxSequencerDrift");
        sequencerWindowSize = stdJson.readUint(_json, "$.sequencerWindowSize");
        channelTimeout = stdJson.readUint(_json, "$.channelTimeout");
        p2pSequencerAddress = stdJson.readAddress(_json, "$.p2pSequencerAddress");
        batchInboxAddress = stdJson.readAddress(_json, "$.batchInboxAddress");
        batchSenderAddress = stdJson.readAddress(_json, "$.batchSenderAddress");
        l2OutputOracleSubmissionInterval = stdJson.readUint(_json, "$.l2OutputOracleSubmissionInterval");
        _l2OutputOracleStartingTimestamp = stdJson.readInt(_json, "$.l2OutputOracleStartingTimestamp");
        l2OutputOracleStartingBlockNumber = stdJson.readUint(_json, "$.l2OutputOracleStartingBlockNumber");
        l2OutputOracleProposer = stdJson.readAddress(_json, "$.l2OutputOracleProposer");
        l2OutputOracleChallenger = stdJson.readAddress(_json, "$.l2OutputOracleChallenger");
        finalizationPeriodSeconds = stdJson.readUint(_json, "$.finalizationPeriodSeconds");
        proxyAdminOwner = stdJson.readAddress(_json, "$.proxyAdminOwner");
        baseFeeVaultRecipient = stdJson.readAddress(_json, "$.baseFeeVaultRecipient");
        l1FeeVaultRecipient = stdJson.readAddress(_json, "$.l1FeeVaultRecipient");
        sequencerFeeVaultRecipient = stdJson.readAddress(_json, "$.sequencerFeeVaultRecipient");
        governanceTokenName = stdJson.readString(_json, "$.governanceTokenName");
        governanceTokenSymbol = stdJson.readString(_json, "$.governanceTokenSymbol");
        governanceTokenOwner = stdJson.readAddress(_json, "$.governanceTokenOwner");
        l2GenesisBlockGasLimit = stdJson.readUint(_json, "$.l2GenesisBlockGasLimit");
        l2GenesisBlockBaseFeePerGas = stdJson.readUint(_json, "$.l2GenesisBlockBaseFeePerGas");
        gasPriceOracleOverhead = stdJson.readUint(_json, "$.gasPriceOracleOverhead");
        gasPriceOracleScalar = stdJson.readUint(_json, "$.gasPriceOracleScalar");
        eip1559Denominator = stdJson.readUint(_json, "$.eip1559Denominator");
        eip1559Elasticity = stdJson.readUint(_json, "$.eip1559Elasticity");
        systemConfigStartBlock = stdJson.readUint(_json, "$.systemConfigStartBlock");
        leaderElectionNumberOfLeaders = stdJson.readUint(_json, "$.leaderElectionNumberOfLeaders");
        leaderElectionNumberOfSlotsPerLeader = stdJson.readUint(_json, "$.leaderElectionNumberOfSlotsPerLeader");
        batchInboxContractAddress = stdJson.readAddress(_json, "$.batchInboxContractAddress");
        initialBatcherVersion = stdJson.readUint(_json, "$.initialBatcherVersion");

        if (block.chainid == Chains.LocalDevnet || block.chainid == Chains.GethDevnet) {
            faultGameAbsolutePrestate = stdJson.readUint(_json, "$.faultGameAbsolutePrestate");
            faultGameMaxDepth = stdJson.readUint(_json, "$.faultGameMaxDepth");
            faultGameMaxDuration = stdJson.readUint(_json, "$.faultGameMaxDuration");
            requiredProtocolVersion = stdJson.readUint(_json, "$.requiredProtocolVersion");
            recommendedProtocolVersion = stdJson.readUint(_json, "$.recommendedProtocolVersion");
        }

        if (block.chainid == Chains.Goerli || block.chainid == Chains.Sepolia) {
            requiredProtocolVersion = stdJson.readUint(_json, "$.requiredProtocolVersion");
            recommendedProtocolVersion = stdJson.readUint(_json, "$.recommendedProtocolVersion");
        }
    }

    function l1StartingBlockTag() public returns (bytes32) {
        try vm.parseJsonBytes32(_json, "$.l1StartingBlockTag") returns (bytes32 tag) {
            return tag;
        } catch {
            try vm.parseJsonString(_json, "$.l1StartingBlockTag") returns (string memory tag) {
                return _getBlockByTag(tag);
            } catch {
                try vm.parseJsonUint(_json, "$.l1StartingBlockTag") returns (uint256 tag) {
                    return _getBlockByTag(vm.toString(tag));
                } catch { }
            }
        }
        revert("l1StartingBlockTag must be a bytes32, string or uint256 or cannot fetch l1StartingBlockTag");
    }

    function l2OutputOracleStartingTimestamp() public returns (uint256) {
        if (_l2OutputOracleStartingTimestamp < 0) {
            bytes32 tag = l1StartingBlockTag();
            string[] memory cmd = new string[](3);
            cmd[0] = Executables.bash;
            cmd[1] = "-c";
            cmd[2] = string.concat("cast block ", vm.toString(tag), " --json | ", Executables.jq, " .timestamp");
            bytes memory res = vm.ffi(cmd);
            return stdJson.readUint(string(res), "");
        }
        return uint256(_l2OutputOracleStartingTimestamp);
    }

    function _getBlockByTag(string memory _tag) internal returns (bytes32) {
        string[] memory cmd = new string[](3);
        cmd[0] = Executables.bash;
        cmd[1] = "-c";
        cmd[2] = string.concat("cast block ", _tag, " --json | ", Executables.jq, " -r .hash");
        bytes memory res = vm.ffi(cmd);
        return abi.decode(res, (bytes32));
    }
}
