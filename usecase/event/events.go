package event

import (
	"github.com/crypto-com/chain-indexing/entity/event"
)

func RegisterEvents(registry *event.Registry) {
	registry.Register(GENESIS_CREATED, 1, DecodeGenesisCreated)

	registry.Register(BLOCK_CREATED, 1, DecodeBlockCreated)
	registry.Register(RAW_BLOCK_CREATED, 1, DecodeRawBlockCreated)
	registry.Register(BLOCK_RAW_EVENT_CREATED, 1, DecodeBlockRawEventCreated)
	registry.Register(TRANSACTION_CREATED, 1, DecodeTransactionCreated)
	registry.Register(TRANSACTION_FAILED, 1, DecodeTransactionFailed)

	registry.Register(ACCOUNT_TRANSFERRED, 1, DecodeAccountTransferred)
	registry.Register(BLOCK_PROPOSER_REWARDED, 1, DecodeBlockProposerRewarded)
	registry.Register(BLOCK_REWARDED, 1, DecodeBlockRewarded)
	registry.Register(BLOCK_COMMISSIONED, 1, DecodeBlockCommissioned)
	registry.Register(MINTED, 1, DecodeMinted)

	registry.Register(POWER_CHANGED, 1, DecodePowerChanged)
	registry.Register(VALIDATOR_SLASHED, 1, DecodeValidatorSlashed)
	registry.Register(VALIDATOR_JAILED, 1, DecodeValidatorJailed)

	// Bank
	registry.Register(MSG_SEND_CREATED, 1, DecodeMsgSend)
	registry.Register(MSG_SEND_FAILED, 1, DecodeMsgSend)
	registry.Register(MSG_MULTI_SEND_CREATED, 1, DecodeMsgMultiSend)
	registry.Register(MSG_MULTI_SEND_FAILED, 1, DecodeMsgMultiSend)

	// Distribution
	registry.Register(MSG_SET_WITHDRAW_ADDRESS_CREATED, 1, DecodeMsgSetWithdrawAddress)
	registry.Register(MSG_SET_WITHDRAW_ADDRESS_FAILED, 1, DecodeMsgSetWithdrawAddress)
	registry.Register(MSG_WITHDRAW_DELEGATOR_REWARD_CREATED, 1, DecodeMsgWithdrawDelegatorReward)
	registry.Register(MSG_WITHDRAW_DELEGATOR_REWARD_FAILED, 1, DecodeMsgWithdrawDelegatorReward)
	registry.Register(MSG_WITHDRAW_VALIDATOR_COMMISSION_CREATED, 1, DecodeMsgWithdrawValidatorCommission)
	registry.Register(MSG_WITHDRAW_VALIDATOR_COMMISSION_FAILED, 1, DecodeMsgWithdrawValidatorCommission)
	registry.Register(MSG_FUND_COMMUNITY_POOL_CREATED, 1, DecodeMsgFundCommunityPool)
	registry.Register(MSG_FUND_COMMUNITY_POOL_FAILED, 1, DecodeMsgFundCommunityPool)

	// Gov
	registry.Register(MSG_SUBMIT_PARAM_CHANGE_PROPOSAL_CREATED, 1, DecodeMsgSubmitParamChangeProposal)
	registry.Register(MSG_SUBMIT_PARAM_CHANGE_PROPOSAL_FAILED, 1, DecodeMsgSubmitParamChangeProposal)
	registry.Register(MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL_CREATED, 1, DecodeMsgSubmitCommunityPoolSpendProposal)
	registry.Register(MSG_SUBMIT_COMMUNITY_POOL_SPEND_PROPOSAL_FAILED, 1, DecodeMsgSubmitCommunityPoolSpendProposal)
	registry.Register(MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL_CREATED, 1, DecodeMsgSubmitSoftwareUpgradeProposal)
	registry.Register(MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL_FAILED, 1, DecodeMsgSubmitSoftwareUpgradeProposal)
	registry.Register(MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL_CREATED, 1, DecodeMsgSubmitCancelSoftwareUpgradeProposal)
	registry.Register(MSG_SUBMIT_CANCEL_SOFTWARE_UPGRADE_PROPOSAL_FAILED, 1, DecodeMsgSubmitCancelSoftwareUpgradeProposal)
	registry.Register(MSG_SUBMIT_TEXT_PROPOSAL_CREATED, 1, DecodeMsgSubmitTextProposal)
	registry.Register(MSG_SUBMIT_TEXT_PROPOSAL_FAILED, 1, DecodeMsgSubmitTextProposal)
	registry.Register(MSG_SUBMIT_UNKNOWN_PROPOSAL_CREATED, 1, DecodeMsgSubmitUnknownProposal)
	registry.Register(MSG_SUBMIT_UNKNOWN_PROPOSAL_FAILED, 1, DecodeMsgSubmitUnknownProposal)
	registry.Register(MSG_DEPOSIT_CREATED, 1, DecodeMsgDeposit)
	registry.Register(MSG_DEPOSIT_FAILED, 1, DecodeMsgDeposit)
	registry.Register(MSG_VOTE_CREATED, 1, DecodeMsgVote)
	registry.Register(MSG_VOTE_FAILED, 1, DecodeMsgVote)

	registry.Register(PROPOSAL_VOTING_PERIOD_STARTED, 1, DecodeProposalVotingPeriodStarted)
	registry.Register(PROPOSAL_ENDED, 1, DecodeProposalEnded)
	registry.Register(PROPOSAL_INACTIVED, 1, DecodeProposalInactived)

	// Staking
	registry.Register(MSG_CREATE_VALIDATOR_CREATED, 1, DecodeMsgCreateValidator)
	registry.Register(MSG_CREATE_VALIDATOR_FAILED, 1, DecodeMsgCreateValidator)
	registry.Register(MSG_EDIT_VALIDATOR_CREATED, 1, DecodeMsgEditValidator)
	registry.Register(MSG_EDIT_VALIDATOR_FAILED, 1, DecodeMsgEditValidator)
	registry.Register(MSG_DELEGATE_CREATED, 1, DecodeMsgDelegate)
	registry.Register(MSG_DELEGATE_FAILED, 1, DecodeMsgDelegate)
	registry.Register(MSG_UNDELEGATE_CREATED, 1, DecodeMsgUndelegate)
	registry.Register(MSG_UNDELEGATE_FAILED, 1, DecodeMsgUndelegate)
	registry.Register(MSG_BEGIN_REDELEGATE_CREATED, 1, DecodeMsgBeginRedelegate)
	registry.Register(MSG_BEGIN_REDELEGATE_FAILED, 1, DecodeMsgBeginRedelegate)

	registry.Register(UNBONDING_COMPLETED, 1, DecodeUnbondingCompleted)

	// Slashing
	registry.Register(MSG_UNJAIL_CREATED, 1, DecodeMsgUnjail)
	registry.Register(MSG_UNJAIL_FAILED, 1, DecodeMsgUnjail)

	// NFT
	registry.Register(MSG_NFT_ISSUE_DENOM_CREATED, 1, DecodeMsgNFTIssueDenom)
	registry.Register(MSG_NFT_ISSUE_DENOM_FAILED, 1, DecodeMsgNFTIssueDenom)
	registry.Register(MSG_NFT_MINT_NFT_CREATED, 1, DecodeMsgNFTMintNFT)
	registry.Register(MSG_NFT_MINT_NFT_FAILED, 1, DecodeMsgNFTMintNFT)
	registry.Register(MSG_NFT_TRANSFER_NFT_CREATED, 1, DecodeMsgNFTTransferNFT)
	registry.Register(MSG_NFT_TRANSFER_NFT_FAILED, 1, DecodeMsgNFTTransferNFT)
	registry.Register(MSG_NFT_EDIT_NFT_CREATED, 1, DecodeMsgNFTEditNFT)
	registry.Register(MSG_NFT_EDIT_NFT_FAILED, 1, DecodeMsgNFTEditNFT)
	registry.Register(MSG_NFT_BURN_NFT_CREATED, 1, DecodeMsgNFTBurnNFT)
	registry.Register(MSG_NFT_BURN_NFT_FAILED, 1, DecodeMsgNFTBurnNFT)

	// IBC
	registry.Register(MSG_IBC_CREATE_CLIENT_CREATED, 1, DecodeMsgIBCCreateClient)
	registry.Register(MSG_IBC_CREATE_CLIENT_FAILED, 1, DecodeMsgIBCCreateClient)
	registry.Register(MSG_IBC_UPDATE_CLIENT_CREATED, 1, DecodeMsgIBCUpdateClient)
	registry.Register(MSG_IBC_UPDATE_CLIENT_FAILED, 1, DecodeMsgIBCUpdateClient)
	registry.Register(MSG_IBC_CONNECTION_OPEN_INIT_CREATED, 1, DecodeMsgIBCConnectionOpenInit)
	registry.Register(MSG_IBC_CONNECTION_OPEN_INIT_FAILED, 1, DecodeMsgIBCConnectionOpenInit)
	registry.Register(MSG_IBC_CONNECTION_OPEN_TRY_CREATED, 1, DecodeMsgIBCConnectionOpenTry)
	registry.Register(MSG_IBC_CONNECTION_OPEN_TRY_FAILED, 1, DecodeMsgIBCConnectionOpenTry)
	registry.Register(MSG_IBC_CONNECTION_OPEN_ACK_CREATED, 1, DecodeMsgIBCConnectionOpenAck)
	registry.Register(MSG_IBC_CONNECTION_OPEN_ACK_FAILED, 1, DecodeMsgIBCConnectionOpenAck)
	registry.Register(MSG_IBC_CONNECTION_OPEN_CONFIRM_CREATED, 1, DecodeMsgIBCConnectionOpenConfirm)
	registry.Register(MSG_IBC_CONNECTION_OPEN_CONFIRM_FAILED, 1, DecodeMsgIBCConnectionOpenConfirm)
	registry.Register(MSG_IBC_CHANNEL_OPEN_INIT_CREATED, 1, DecodeMsgIBCChannelOpenInit)
	registry.Register(MSG_IBC_CHANNEL_OPEN_INIT_FAILED, 1, DecodeMsgIBCChannelOpenInit)
	registry.Register(MSG_IBC_CHANNEL_OPEN_TRY_CREATED, 1, DecodeMsgIBCChannelOpenTry)
	registry.Register(MSG_IBC_CHANNEL_OPEN_TRY_FAILED, 1, DecodeMsgIBCChannelOpenTry)
	registry.Register(MSG_IBC_CHANNEL_OPEN_ACK_CREATED, 1, DecodeMsgIBCChannelOpenAck)
	registry.Register(MSG_IBC_CHANNEL_OPEN_ACK_FAILED, 1, DecodeMsgIBCChannelOpenAck)
	registry.Register(MSG_IBC_CHANNEL_OPEN_CONFIRM_CREATED, 1, DecodeMsgIBCChannelOpenConfirm)
	registry.Register(MSG_IBC_CHANNEL_OPEN_CONFIRM_FAILED, 1, DecodeMsgIBCChannelOpenConfirm)
	registry.Register(MSG_IBC_CHANNEL_CLOSE_INIT_CREATED, 1, DecodeMsgIBCChannelCloseInit)
	registry.Register(MSG_IBC_CHANNEL_CLOSE_INIT_FAILED, 1, DecodeMsgIBCChannelCloseInit)
	registry.Register(MSG_IBC_CHANNEL_CLOSE_CONFIRM_CREATED, 1, DecodeMsgIBCChannelCloseConfirm)
	registry.Register(MSG_IBC_CHANNEL_CLOSE_CONFIRM_FAILED, 1, DecodeMsgIBCChannelCloseConfirm)

	registry.Register(MSG_IBC_RECV_PACKET_CREATED, 1, DecodeMsgIBCRecvPacket)
	registry.Register(MSG_IBC_RECV_PACKET_FAILED, 1, DecodeMsgIBCRecvPacket)
	registry.Register(MSG_ALREADY_RELAYED_IBC_RECV_PACKET_CREATED, 1, DecodeMsgAlreadyRelayedIBCRecvPacket)
	registry.Register(MSG_ALREADY_RELAYED_IBC_RECV_PACKET_FAILED, 1, DecodeMsgAlreadyRelayedIBCRecvPacket)
	registry.Register(MSG_IBC_ACKNOWLEDGEMENT_CREATED, 1, DecodeMsgIBCAcknowledgement)
	registry.Register(MSG_IBC_ACKNOWLEDGEMENT_FAILED, 1, DecodeMsgIBCAcknowledgement)
	registry.Register(MSG_ALREADY_RELAYED_IBC_ACKNOWLEDGEMENT_CREATED, 1, DecodeMsgAlreadyRelayedIBCAcknowledgement)
	registry.Register(MSG_ALREADY_RELAYED_IBC_ACKNOWLEDGEMENT_FAILED, 1, DecodeMsgAlreadyRelayedIBCAcknowledgement)
	registry.Register(MSG_IBC_TIMEOUT_CREATED, 1, DecodeMsgIBCTimeout)
	registry.Register(MSG_IBC_TIMEOUT_FAILED, 1, DecodeMsgIBCTimeout)
	registry.Register(MSG_ALREADY_RELAYED_IBC_TIMEOUT_CREATED, 1, DecodeMsgAlreadyRelayedIBCTimeout)
	registry.Register(MSG_ALREADY_RELAYED_IBC_TIMEOUT_FAILED, 1, DecodeMsgAlreadyRelayedIBCTimeout)
	registry.Register(MSG_IBC_TIMEOUT_ON_CLOSE_CREATED, 1, DecodeMsgIBCTimeoutOnClose)
	registry.Register(MSG_IBC_TIMEOUT_ON_CLOSE_FAILED, 1, DecodeMsgIBCTimeoutOnClose)
	registry.Register(MSG_ALREADY_RELAYED_IBC_TIMEOUT_ON_CLOSE_CREATED, 1, DecodeMsgAlreadyRelayedIBCTimeoutOnClose)
	registry.Register(MSG_ALREADY_RELAYED_IBC_TIMEOUT_ON_CLOSE_FAILED, 1, DecodeMsgAlreadyRelayedIBCTimeoutOnClose)

	registry.Register(MSG_IBC_TRANSFER_TRANSFER_CREATED, 1, DecodeMsgIBCTransferTransfer)
	registry.Register(MSG_IBC_TRANSFER_TRANSFER_FAILED, 1, DecodeMsgIBCTransferTransfer)

	// Authz
	registry.Register(MSG_GRANT_CREATED, 1, DecodeMsgGrant)
	registry.Register(MSG_GRANT_FAILED, 1, DecodeMsgGrant)
	registry.Register(MSG_REVOKE_CREATED, 1, DecodeMsgRevoke)
	registry.Register(MSG_REVOKE_FAILED, 1, DecodeMsgRevoke)
	registry.Register(MSG_EXEC_CREATED, 1, DecodeMsgExec)
	registry.Register(MSG_EXEC_FAILED, 1, DecodeMsgExec)

	// Feegrant
	registry.Register(MSG_GRANT_ALLOWANCE_CREATED, 1, DecodeMsgGrantAllowance)
	registry.Register(MSG_GRANT_ALLOWANCE_FAILED, 1, DecodeMsgGrantAllowance)
	registry.Register(MSG_REVOKE_ALLOWANCE_CREATED, 1, DecodeMsgRevokeAllowance)
	registry.Register(MSG_REVOKE_ALLOWANCE_FAILED, 1, DecodeMsgRevokeAllowance)

	// vesting
	registry.Register(MSG_CREATE_VESTING_ACCOUNT_CREATED, 1, DecodeMsgCreateVestingAccount)
	registry.Register(MSG_CREATE_VESTING_ACCOUNT_FAILED, 1, DecodeMsgCreateVestingAccount)

	// Gravity
	registry.Register(GRAVITY_ETHEREUM_SEND_TO_COSMOS_HANDLED, 1, DecodeGravityEthereumSendToCosmosHandled)

	// Cronos
	registry.Register(CRONOS_SEND_TO_IBC_CREATED, 1, DecodeCronosSendToIBCCreated)

	// Ethermint tx
	registry.Register(MSG_ETHEREUM_TX_CREATED, 1, DecodeMsgEthereumTx)
	registry.Register(MSG_ETHEREUM_TX_FAILED, 1, DecodeMsgEthereumTx)
}
