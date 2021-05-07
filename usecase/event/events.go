package event

import (
	"github.com/crypto-com/chain-indexing/entity/event"
)

func RegisterEvents(registry *event.Registry) {
	registry.Register(GENESIS_CREATED, 1, DecodeGenesisCreated)

	registry.Register(BLOCK_CREATED, 1, DecodeBlockCreated)
	registry.Register(RAW_BLOCK_CREATED, 1, DecodeRawBlockCreated)
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

	// IBC
	registry.Register(MSG_CREATE_CLIENT_CREATED, 1, DecodeMsgCreateClient)
	registry.Register(MSG_CREATE_CLIENT_FAILED, 1, DecodeMsgCreateClient)
	registry.Register(MSG_CONNECTION_OPEN_INIT_CREATED, 1, DecodeMsgConnectionOpenInit)
	registry.Register(MSG_CONNECTION_OPEN_INIT_FAILED, 1, DecodeMsgConnectionOpenInit)
}
