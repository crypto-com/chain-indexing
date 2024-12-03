package parser

import (
	cosmos_gov_v1 "github.com/crypto-com/chain-indexing/usecase/parser/gov/v1"
	"github.com/crypto-com/chain-indexing/usecase/parser/ibc"
	"github.com/crypto-com/chain-indexing/usecase/parser/icaauth"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
	V0_42_7_ibcmsg "github.com/crypto-com/chain-indexing/usecase/parser/v0_42_7/ibcmsg"
	V1_4_0_cronos_msg "github.com/crypto-com/chain-indexing/usecase/parser/v1_4_0_cronos/msg"
)

const BEGIN_BLOCK_HEIGHT = 0

func InitParsers(manager *utils.CosmosParserManager) {
	// cosmos bank
	manager.RegisterParser("/cosmos.bank.v1beta1.MsgSend", BEGIN_BLOCK_HEIGHT, ParseMsgSend)
	manager.RegisterParser("/cosmos.bank.v1beta1.MsgMultiSend", BEGIN_BLOCK_HEIGHT, ParseMsgMultiSend)
	manager.RegisterParser("/cosmos.bank.v1beta1.MsgSetSendEnabled", BEGIN_BLOCK_HEIGHT, ParseMsgSetSendEnabled)

	// cosmos distribution
	manager.RegisterParser("/cosmos.distribution.v1beta1.MsgSetWithdrawAddress", BEGIN_BLOCK_HEIGHT, ParseMsgSetWithdrawAddress)
	manager.RegisterParser("/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward", BEGIN_BLOCK_HEIGHT, ParseMsgWithdrawDelegatorReward)
	manager.RegisterParser("/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission", BEGIN_BLOCK_HEIGHT, ParseMsgWithdrawValidatorCommission)
	manager.RegisterParser("/cosmos.distribution.v1beta1.MsgFundCommunityPool", BEGIN_BLOCK_HEIGHT, ParseMsgFundCommunityPool)

	// cosmos gov
	manager.RegisterParser("/cosmos.gov.v1beta1.MsgSubmitProposal", BEGIN_BLOCK_HEIGHT, ParseMsgSubmitProposal)
	manager.RegisterParser("/cosmos.gov.v1beta1.MsgVote", BEGIN_BLOCK_HEIGHT, ParseMsgVote)
	manager.RegisterParser("/cosmos.gov.v1beta1.MsgDeposit", BEGIN_BLOCK_HEIGHT, ParseMsgDeposit)

	// cosmos gov v1
	manager.RegisterParser("/cosmos.gov.v1.MsgDeposit", BEGIN_BLOCK_HEIGHT, cosmos_gov_v1.ParseMsgDeposit)
	manager.RegisterParser("/cosmos.gov.v1.MsgSubmitProposal", BEGIN_BLOCK_HEIGHT, cosmos_gov_v1.ParseMsgSubmitProposal)
	manager.RegisterParser("/cosmos.gov.v1.MsgVote", BEGIN_BLOCK_HEIGHT, cosmos_gov_v1.ParseMsgVote)
	manager.RegisterParser("/cosmos.gov.v1.MsgVoteWeighted", BEGIN_BLOCK_HEIGHT, cosmos_gov_v1.ParseMsgVoteWeighted)

	// cosmos staking
	manager.RegisterParser("/cosmos.staking.v1beta1.MsgDelegate", BEGIN_BLOCK_HEIGHT, ParseMsgDelegate)
	manager.RegisterParser("/cosmos.staking.v1beta1.MsgUndelegate", BEGIN_BLOCK_HEIGHT, ParseMsgUndelegate)
	manager.RegisterParser("/cosmos.staking.v1beta1.MsgBeginRedelegate", BEGIN_BLOCK_HEIGHT, ParseMsgBeginRedelegate)
	manager.RegisterParser("/cosmos.staking.v1beta1.MsgCreateValidator", BEGIN_BLOCK_HEIGHT, ParseMsgCreateValidator)
	manager.RegisterParser("/cosmos.staking.v1beta1.MsgEditValidator", BEGIN_BLOCK_HEIGHT, ParseMsgEditValidator)

	// cosmos slashing
	manager.RegisterParser("/cosmos.slashing.v1beta1.MsgUnjail", BEGIN_BLOCK_HEIGHT, ParseMsgUnjail)

	// chainmain nft
	manager.RegisterParser("/chainmain.nft.v1.MsgIssueDenom", BEGIN_BLOCK_HEIGHT, ParseMsgNFTIssueDenom)
	manager.RegisterParser("/chainmain.nft.v1.MsgMintNFT", BEGIN_BLOCK_HEIGHT, ParseMsgNFTMintNFT)
	manager.RegisterParser("/chainmain.nft.v1.MsgTransferNFT", BEGIN_BLOCK_HEIGHT, ParseMsgNFTTransferNFT)
	manager.RegisterParser("/chainmain.nft.v1.MsgEditNFT", BEGIN_BLOCK_HEIGHT, ParseMsgNFTEditNFT)
	manager.RegisterParser("/chainmain.nft.v1.MsgBurnNFT", BEGIN_BLOCK_HEIGHT, ParseMsgNFTBurnNFT)

	// chainmain icaauth
	manager.RegisterParser("/chainmain.icaauth.v1.MsgRegisterAccount", BEGIN_BLOCK_HEIGHT, icaauth.ParseChainmainMsgRegisterAccount)
	manager.RegisterParser("/chainmain.icaauth.v1.MsgSubmitTx", BEGIN_BLOCK_HEIGHT, icaauth.ParseChainmainMsgSubmitTx)

	// cronos icaauth
	manager.RegisterParser("/icaauth.v1.MsgRegisterAccount", BEGIN_BLOCK_HEIGHT, icaauth.ParseMsgRegisterAccount)
	manager.RegisterParser("/icaauth.v1.MsgSubmitTx", BEGIN_BLOCK_HEIGHT, icaauth.ParseMsgSubmitTx)

	// ibc core client
	manager.RegisterParser("/ibc.core.client.v1.MsgCreateClient", BEGIN_BLOCK_HEIGHT, ibc.ParseMsgCreateClient)
	manager.RegisterParser("/ibc.core.client.v1.MsgUpdateClient", BEGIN_BLOCK_HEIGHT, ibc.ParseMsgUpdateClient)

	// ibc core connection
	manager.RegisterParser("/ibc.core.connection.v1.MsgConnectionOpenInit", BEGIN_BLOCK_HEIGHT, ibc.ParseMsgConnectionOpenInit)
	manager.RegisterParser("/ibc.core.connection.v1.MsgConnectionOpenTry", BEGIN_BLOCK_HEIGHT, ibc.ParseMsgConnectionOpenTry)
	manager.RegisterParser("/ibc.core.connection.v1.MsgConnectionOpenAck", BEGIN_BLOCK_HEIGHT, ibc.ParseMsgConnectionOpenAck)
	manager.RegisterParser("/ibc.core.connection.v1.MsgConnectionOpenConfirm", BEGIN_BLOCK_HEIGHT, ibc.ParseMsgConnectionOpenConfirm)

	// ibc core channel
	manager.RegisterParser("/ibc.core.channel.v1.MsgChannelOpenInit", BEGIN_BLOCK_HEIGHT, ibc.ParseMsgChannelOpenInit)
	manager.RegisterParser("/ibc.core.channel.v1.MsgChannelOpenTry", BEGIN_BLOCK_HEIGHT, ibc.ParseMsgChannelOpenTry)
	manager.RegisterParser("/ibc.core.channel.v1.MsgChannelOpenAck", BEGIN_BLOCK_HEIGHT, ibc.ParseMsgChannelOpenAck)
	manager.RegisterParser("/ibc.core.channel.v1.MsgChannelOpenConfirm", BEGIN_BLOCK_HEIGHT, ibc.ParseMsgChannelOpenConfirm)
	manager.RegisterParser("/ibc.core.channel.v1.MsgRecvPacket", BEGIN_BLOCK_HEIGHT, ibc.ParseMsgRecvPacket)
	manager.RegisterParser("/ibc.core.channel.v1.MsgAcknowledgement", BEGIN_BLOCK_HEIGHT, ibc.ParseMsgAcknowledgement)
	manager.RegisterParser("/ibc.core.channel.v1.MsgTimeout", BEGIN_BLOCK_HEIGHT, ibc.ParseMsgTimeout)
	manager.RegisterParser("/ibc.core.channel.v1.MsgTimeoutOnClose", BEGIN_BLOCK_HEIGHT, ibc.ParseMsgTimeoutOnClose)
	manager.RegisterParser("/ibc.core.channel.v1.MsgChannelCloseInit", BEGIN_BLOCK_HEIGHT, ibc.ParseMsgChannelCloseInit)
	manager.RegisterParser("/ibc.core.channel.v1.MsgChannelCloseConfirm", BEGIN_BLOCK_HEIGHT, ibc.ParseMsgChannelCloseConfirm)

	// ibc applications transfer
	manager.RegisterParser("/ibc.applications.transfer.v1.MsgTransfer", BEGIN_BLOCK_HEIGHT, ibc.ParseMsgTransfer)

	// cosmos authz
	manager.RegisterParser("/cosmos.authz.v1beta1.MsgGrant", BEGIN_BLOCK_HEIGHT, ParseMsgGrant)
	manager.RegisterParser("/cosmos.authz.v1beta1.MsgRevoke", BEGIN_BLOCK_HEIGHT, ParseMsgRevoke)
	// FIXME: https://github.com/crypto-com/chain-indexing/issues/673
	//manager.RegisterParser("/cosmos.authz.v1beta1.MsgExec", BEGIN_BLOCK_HEIGHT, ParseMsgExec)

	// cosmos feegrant
	manager.RegisterParser("/cosmos.feegrant.v1beta1.MsgGrantAllowance", BEGIN_BLOCK_HEIGHT, ParseMsgGrantAllowance)
	manager.RegisterParser("/cosmos.feegrant.v1beta1.MsgRevokeAllowance", BEGIN_BLOCK_HEIGHT, ParseMsgRevokeAllowance)

	// cosmos vesting
	manager.RegisterParser("/cosmos.vesting.v1beta1.MsgCreateVestingAccount", BEGIN_BLOCK_HEIGHT, ParseMsgCreateVestingAccount)

	// ethermint evm
	manager.RegisterParser("/ethermint.evm.v1.MsgEthereumTx", BEGIN_BLOCK_HEIGHT, ParseMsgEthereumTx)
}

func RegisterBreakingVersionParsers(manager *utils.CosmosParserManager) {
	// cosmos sdk v0.42.7
	manager.RegisterParser("/ibc.core.channel.v1.MsgRecvPacket", manager.GetCosmosV0_42_7BlockHeight(), V0_42_7_ibcmsg.ParseMsgRecvPacket)

	/*
	   Cronos v1.4.0
	*/

	if manager.GetCronosV1_4_0BlockHeight() > 0 {
		manager.RegisterParser("/ethermint.evm.v1.MsgEthereumTx", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEthereumTx)

		// cosmos bank
		manager.RegisterParser("/cosmos.bank.v1beta1.MsgSend", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ParseMsgSend))
		manager.RegisterParser("/cosmos.bank.v1beta1.MsgMultiSend", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ParseMsgMultiSend))
		manager.RegisterParser("/cosmos.bank.v1beta1.MsgSetSendEnabled", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ParseMsgSetSendEnabled))

		// cosmos distribution
		manager.RegisterParser("/cosmos.distribution.v1beta1.MsgSetWithdrawAddress", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ParseMsgSetWithdrawAddress))
		manager.RegisterParser("/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ParseMsgWithdrawDelegatorReward))
		manager.RegisterParser("/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ParseMsgWithdrawValidatorCommission))
		manager.RegisterParser("/cosmos.distribution.v1beta1.MsgFundCommunityPool", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ParseMsgFundCommunityPool))

		// cosmos gov
		manager.RegisterParser("/cosmos.gov.v1beta1.MsgSubmitProposal", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ParseMsgSubmitProposal))
		manager.RegisterParser("/cosmos.gov.v1beta1.MsgVote", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ParseMsgVote))
		manager.RegisterParser("/cosmos.gov.v1beta1.MsgDeposit", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ParseMsgDeposit))

		// cosmos gov v1
		manager.RegisterParser("/cosmos.gov.v1.MsgDeposit", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(cosmos_gov_v1.ParseMsgDeposit))
		manager.RegisterParser("/cosmos.gov.v1.MsgSubmitProposal", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(cosmos_gov_v1.ParseMsgSubmitProposal))
		manager.RegisterParser("/cosmos.gov.v1.MsgVote", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(cosmos_gov_v1.ParseMsgVote))
		manager.RegisterParser("/cosmos.gov.v1.MsgVoteWeighted", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(cosmos_gov_v1.ParseMsgVoteWeighted))

		// cosmos staking
		manager.RegisterParser("/cosmos.staking.v1beta1.MsgDelegate", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ParseMsgDelegate))
		manager.RegisterParser("/cosmos.staking.v1beta1.MsgUndelegate", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ParseMsgUndelegate))
		manager.RegisterParser("/cosmos.staking.v1beta1.MsgBeginRedelegate", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ParseMsgBeginRedelegate))
		manager.RegisterParser("/cosmos.staking.v1beta1.MsgCreateValidator", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ParseMsgCreateValidator))
		manager.RegisterParser("/cosmos.staking.v1beta1.MsgEditValidator", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ParseMsgEditValidator))

		// cosmos slashing
		manager.RegisterParser("/cosmos.slashing.v1beta1.MsgUnjail", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ParseMsgUnjail))

		// cronos icaauth
		manager.RegisterParser("/icaauth.v1.MsgRegisterAccount", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(icaauth.ParseMsgRegisterAccount))
		manager.RegisterParser("/icaauth.v1.MsgSubmitTx", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(icaauth.ParseMsgSubmitTx))

		// ibc core client
		manager.RegisterParser("/ibc.core.client.v1.MsgCreateClient", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ibc.ParseMsgCreateClient))
		manager.RegisterParser("/ibc.core.client.v1.MsgUpdateClient", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ibc.ParseMsgUpdateClient))

		// ibc core connection
		manager.RegisterParser("/ibc.core.connection.v1.MsgConnectionOpenInit", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ibc.ParseMsgConnectionOpenInit))
		manager.RegisterParser("/ibc.core.connection.v1.MsgConnectionOpenTry", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ibc.ParseMsgConnectionOpenTry))
		manager.RegisterParser("/ibc.core.connection.v1.MsgConnectionOpenAck", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ibc.ParseMsgConnectionOpenAck))
		manager.RegisterParser("/ibc.core.connection.v1.MsgConnectionOpenConfirm", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ibc.ParseMsgConnectionOpenConfirm))

		// ibc core channel
		manager.RegisterParser("/ibc.core.channel.v1.MsgChannelOpenInit", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ibc.ParseMsgChannelOpenInit))
		manager.RegisterParser("/ibc.core.channel.v1.MsgChannelOpenTry", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ibc.ParseMsgChannelOpenTry))
		manager.RegisterParser("/ibc.core.channel.v1.MsgChannelOpenAck", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ibc.ParseMsgChannelOpenAck))
		manager.RegisterParser("/ibc.core.channel.v1.MsgChannelOpenConfirm", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ibc.ParseMsgChannelOpenConfirm))
		manager.RegisterParser("/ibc.core.channel.v1.MsgRecvPacket", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(V0_42_7_ibcmsg.ParseMsgRecvPacket))
		manager.RegisterParser("/ibc.core.channel.v1.MsgAcknowledgement", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ibc.ParseMsgAcknowledgement))
		manager.RegisterParser("/ibc.core.channel.v1.MsgTimeout", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ibc.ParseMsgTimeout))
		manager.RegisterParser("/ibc.core.channel.v1.MsgTimeoutOnClose", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ibc.ParseMsgTimeoutOnClose))
		manager.RegisterParser("/ibc.core.channel.v1.MsgChannelCloseInit", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ibc.ParseMsgChannelCloseInit))
		manager.RegisterParser("/ibc.core.channel.v1.MsgChannelCloseConfirm", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ibc.ParseMsgChannelCloseConfirm))

		// ibc applications transfer
		manager.RegisterParser("/ibc.applications.transfer.v1.MsgTransfer", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ibc.ParseMsgTransfer))

		// cosmos authz
		manager.RegisterParser("/cosmos.authz.v1beta1.MsgGrant", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ParseMsgGrant))
		manager.RegisterParser("/cosmos.authz.v1beta1.MsgRevoke", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ParseMsgRevoke))
		// FIXME: https://github.com/crypto-com/chain-indexing/issues/673
		//manager.RegisterParser("/cosmos.authz.v1beta1.MsgExec",  manager.GetCronosV1_4_0BlockHeight(), ParseMsgExec))

		// cosmos feegrant
		manager.RegisterParser("/cosmos.feegrant.v1beta1.MsgGrantAllowance", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ParseMsgGrantAllowance))
		manager.RegisterParser("/cosmos.feegrant.v1beta1.MsgRevokeAllowance", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ParseMsgRevokeAllowance))

		// cosmos vesting
		manager.RegisterParser("/cosmos.vesting.v1beta1.MsgCreateVestingAccount", manager.GetCronosV1_4_0BlockHeight(), V1_4_0_cronos_msg.ParseMsgEventsToLog(ParseMsgCreateVestingAccount))
	}
}
