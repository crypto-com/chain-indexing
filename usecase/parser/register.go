package parser

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/parser/ibcmsg"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
	V0_43_0_ibcmsg "github.com/crypto-com/chain-indexing/usecase/parser/v0_43_0/ibcmsg"
)

const BEGIN_BLOCK_HEIGHT = 0

var MSG_TYPE_TO_EVENT_NAME_MAP = map[string]utils.CosmosParserKey{}

func InitParsers(manager *utils.CosmosParserManager) {
	// cosmos bank
	MSG_TYPE_TO_EVENT_NAME_MAP["/cosmos.bank.v1beta1.MsgSend"] = event.MSG_SEND
	manager.RegisterParser(event.MSG_SEND, BEGIN_BLOCK_HEIGHT, ParseMsgSend)
	MSG_TYPE_TO_EVENT_NAME_MAP["/cosmos.bank.v1beta1.MsgMultiSend"] = event.MSG_MULTI_SEND
	manager.RegisterParser(event.MSG_MULTI_SEND, BEGIN_BLOCK_HEIGHT, ParseMsgMultiSend)

	// cosmos distribution
	MSG_TYPE_TO_EVENT_NAME_MAP["/cosmos.distribution.v1beta1.MsgSetWithdrawAddress"] = event.MSG_SET_WITHDRAW_ADDRESS
	manager.RegisterParser(event.MSG_SET_WITHDRAW_ADDRESS, BEGIN_BLOCK_HEIGHT, ParseMsgSetWithdrawAddress)
	MSG_TYPE_TO_EVENT_NAME_MAP["/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward"] = event.MSG_WITHDRAW_DELEGATOR_REWARD
	manager.RegisterParser(event.MSG_WITHDRAW_DELEGATOR_REWARD, BEGIN_BLOCK_HEIGHT, ParseMsgWithdrawDelegatorReward)
	MSG_TYPE_TO_EVENT_NAME_MAP["/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission"] = event.MSG_WITHDRAW_VALIDATOR_COMMISSION
	manager.RegisterParser(event.MSG_WITHDRAW_VALIDATOR_COMMISSION, BEGIN_BLOCK_HEIGHT, ParseMsgWithdrawValidatorCommission)
	MSG_TYPE_TO_EVENT_NAME_MAP["/cosmos.distribution.v1beta1.MsgFundCommunityPool"] = event.MSG_FUND_COMMUNITY_POOL
	manager.RegisterParser(event.MSG_FUND_COMMUNITY_POOL, BEGIN_BLOCK_HEIGHT, ParseMsgFundCommunityPool)

	// cosmos gov
	MSG_TYPE_TO_EVENT_NAME_MAP["/cosmos.gov.v1beta1.MsgSubmitProposal"] = event.MSG_SUBMIT_PROPOSAL
	manager.RegisterParser(event.MSG_SUBMIT_PROPOSAL, BEGIN_BLOCK_HEIGHT, ParseMsgSubmitProposal)
	MSG_TYPE_TO_EVENT_NAME_MAP["/cosmos.gov.v1beta1.MsgVote"] = event.MSG_VOTE
	manager.RegisterParser(event.MSG_VOTE, BEGIN_BLOCK_HEIGHT, ParseMsgVote)
	MSG_TYPE_TO_EVENT_NAME_MAP["/cosmos.gov.v1beta1.MsgDeposit"] = event.MSG_DEPOSIT
	manager.RegisterParser(event.MSG_DEPOSIT, BEGIN_BLOCK_HEIGHT, ParseMsgDeposit)

	// cosmos staking
	MSG_TYPE_TO_EVENT_NAME_MAP["/cosmos.staking.v1beta1.MsgDelegate"] = event.MSG_DELEGATE
	manager.RegisterParser(event.MSG_DELEGATE, BEGIN_BLOCK_HEIGHT, ParseMsgDelegate)
	MSG_TYPE_TO_EVENT_NAME_MAP["/cosmos.staking.v1beta1.MsgUndelegate"] = event.MSG_UNDELEGATE
	manager.RegisterParser(event.MSG_UNDELEGATE, BEGIN_BLOCK_HEIGHT, ParseMsgUndelegate)
	MSG_TYPE_TO_EVENT_NAME_MAP["/cosmos.staking.v1beta1.MsgBeginRedelegate"] = event.MSG_BEGIN_REDELEGATE
	manager.RegisterParser(event.MSG_BEGIN_REDELEGATE, BEGIN_BLOCK_HEIGHT, ParseMsgBeginRedelegate)
	MSG_TYPE_TO_EVENT_NAME_MAP["/cosmos.staking.v1beta1.MsgCreateValidator"] = event.MSG_CREATE_VALIDATOR
	manager.RegisterParser(event.MSG_CREATE_VALIDATOR, BEGIN_BLOCK_HEIGHT, ParseMsgCreateValidator)
	MSG_TYPE_TO_EVENT_NAME_MAP["/cosmos.staking.v1beta1.MsgEditValidator"] = event.MSG_EDIT_VALIDATOR
	manager.RegisterParser(event.MSG_EDIT_VALIDATOR, BEGIN_BLOCK_HEIGHT, ParseMsgEditValidator)

	// cosmos slashing
	MSG_TYPE_TO_EVENT_NAME_MAP["/cosmos.slashing.v1beta1.MsgUnjail"] = event.MSG_UNJAIL
	manager.RegisterParser(event.MSG_UNJAIL, BEGIN_BLOCK_HEIGHT, ParseMsgUnjail)

	// chainmain nft
	MSG_TYPE_TO_EVENT_NAME_MAP["/chainmain.nft.v1.MsgIssueDenom"] = event.MSG_NFT_ISSUE_DENOM
	manager.RegisterParser(event.MSG_NFT_ISSUE_DENOM, BEGIN_BLOCK_HEIGHT, ParseMsgNFTIssueDenom)
	MSG_TYPE_TO_EVENT_NAME_MAP["/chainmain.nft.v1.MsgMintNFT"] = event.MSG_NFT_MINT_NFT
	manager.RegisterParser(event.MSG_NFT_MINT_NFT, BEGIN_BLOCK_HEIGHT, ParseMsgNFTMintNFT)
	MSG_TYPE_TO_EVENT_NAME_MAP["/chainmain.nft.v1.MsgTransferNFT"] = event.MSG_NFT_TRANSFER_NFT
	manager.RegisterParser(event.MSG_NFT_TRANSFER_NFT, BEGIN_BLOCK_HEIGHT, ParseMsgNFTTransferNFT)
	MSG_TYPE_TO_EVENT_NAME_MAP["/chainmain.nft.v1.MsgEditNFT"] = event.MSG_NFT_EDIT_NFT
	manager.RegisterParser(event.MSG_NFT_EDIT_NFT, BEGIN_BLOCK_HEIGHT, ParseMsgNFTEditNFT)
	MSG_TYPE_TO_EVENT_NAME_MAP["/chainmain.nft.v1.MsgBurnNFT"] = event.MSG_NFT_BURN_NFT
	manager.RegisterParser(event.MSG_NFT_BURN_NFT, BEGIN_BLOCK_HEIGHT, ParseMsgNFTBurnNFT)

	// ibc core client
	MSG_TYPE_TO_EVENT_NAME_MAP["/ibc.core.client.v1.MsgCreateClient"] = event.MSG_IBC_CREATE_CLIENT
	manager.RegisterParser(event.MSG_IBC_CREATE_CLIENT, BEGIN_BLOCK_HEIGHT, ibcmsg.ParseMsgCreateClient)
	MSG_TYPE_TO_EVENT_NAME_MAP["/ibc.core.client.v1.MsgUpdateClient"] = event.MSG_IBC_UPDATE_CLIENT
	manager.RegisterParser(event.MSG_IBC_UPDATE_CLIENT, BEGIN_BLOCK_HEIGHT, ibcmsg.ParseMsgUpdateClient)

	// ibc core connection
	MSG_TYPE_TO_EVENT_NAME_MAP["/ibc.core.connection.v1.MsgConnectionOpenInit"] = event.MSG_IBC_CONNECTION_OPEN_INIT
	manager.RegisterParser(event.MSG_IBC_CONNECTION_OPEN_INIT, BEGIN_BLOCK_HEIGHT, ibcmsg.ParseMsgConnectionOpenInit)
	MSG_TYPE_TO_EVENT_NAME_MAP["/ibc.core.connection.v1.MsgConnectionOpenTry"] = event.MSG_IBC_CONNECTION_OPEN_TRY
	manager.RegisterParser(event.MSG_IBC_CONNECTION_OPEN_TRY, BEGIN_BLOCK_HEIGHT, ibcmsg.ParseMsgConnectionOpenTry)
	MSG_TYPE_TO_EVENT_NAME_MAP["/ibc.core.connection.v1.MsgConnectionOpenAck"] = event.MSG_IBC_CONNECTION_OPEN_ACK
	manager.RegisterParser(event.MSG_IBC_CONNECTION_OPEN_ACK, BEGIN_BLOCK_HEIGHT, ibcmsg.ParseMsgConnectionOpenAck)
	MSG_TYPE_TO_EVENT_NAME_MAP["/ibc.core.connection.v1.MsgConnectionOpenConfirm"] = event.MSG_IBC_CONNECTION_OPEN_CONFIRM
	manager.RegisterParser(event.MSG_IBC_CONNECTION_OPEN_CONFIRM, BEGIN_BLOCK_HEIGHT, ibcmsg.ParseMsgConnectionOpenConfirm)

	// ibc core channel
	MSG_TYPE_TO_EVENT_NAME_MAP["/ibc.core.channel.v1.MsgChannelOpenInit"] = event.MSG_IBC_CHANNEL_OPEN_INIT
	manager.RegisterParser(event.MSG_IBC_CHANNEL_OPEN_INIT, BEGIN_BLOCK_HEIGHT, ibcmsg.ParseMsgChannelOpenInit)
	MSG_TYPE_TO_EVENT_NAME_MAP["/ibc.core.channel.v1.ParseMsgChannelOpenTry"] = event.MSG_IBC_CHANNEL_OPEN_TRY
	manager.RegisterParser(event.MSG_IBC_CHANNEL_OPEN_TRY, BEGIN_BLOCK_HEIGHT, ibcmsg.ParseMsgChannelOpenTry)
	MSG_TYPE_TO_EVENT_NAME_MAP["/ibc.core.channel.v1.ParseMsgChannelOpenAck"] = event.MSG_IBC_CHANNEL_OPEN_ACK
	manager.RegisterParser(event.MSG_IBC_CHANNEL_OPEN_ACK, BEGIN_BLOCK_HEIGHT, ibcmsg.ParseMsgChannelOpenAck)
	MSG_TYPE_TO_EVENT_NAME_MAP["/ibc.core.channel.v1.ParseMsgChannelOpenConfirm"] = event.MSG_IBC_CHANNEL_OPEN_CONFIRM
	manager.RegisterParser(event.MSG_IBC_CHANNEL_OPEN_CONFIRM, BEGIN_BLOCK_HEIGHT, ibcmsg.ParseMsgChannelOpenConfirm)
	MSG_TYPE_TO_EVENT_NAME_MAP["/ibc.core.channel.v1.MsgRecvPacket"] = event.MSG_IBC_RECV_PACKET
	manager.RegisterParser(event.MSG_IBC_RECV_PACKET, BEGIN_BLOCK_HEIGHT, ibcmsg.ParseMsgRecvPacket)
	MSG_TYPE_TO_EVENT_NAME_MAP["/ibc.core.channel.v1.MsgAcknowledgement"] = event.MSG_IBC_ACKNOWLEDGEMENT
	manager.RegisterParser(event.MSG_IBC_ACKNOWLEDGEMENT, BEGIN_BLOCK_HEIGHT, ibcmsg.ParseMsgAcknowledgement)
	MSG_TYPE_TO_EVENT_NAME_MAP["/ibc.core.channel.v1.MsgTimeout"] = event.MSG_IBC_TIMEOUT
	manager.RegisterParser(event.MSG_IBC_TIMEOUT, BEGIN_BLOCK_HEIGHT, ibcmsg.ParseMsgTimeout)
	MSG_TYPE_TO_EVENT_NAME_MAP["/ibc.core.channel.v1.MsgTimeoutOnClose"] = event.MSG_IBC_TIMEOUT_ON_CLOSE
	manager.RegisterParser(event.MSG_IBC_TIMEOUT_ON_CLOSE, BEGIN_BLOCK_HEIGHT, ibcmsg.ParseMsgTimeoutOnClose)

	// ibc applications transfer
	MSG_TYPE_TO_EVENT_NAME_MAP["/ibc.applications.transfer.v1.MsgTransfer"] = event.MSG_IBC_TRANSFER_TRANSFER
	manager.RegisterParser(event.MSG_IBC_TRANSFER_TRANSFER, BEGIN_BLOCK_HEIGHT, ibcmsg.ParseMsgTransfer)

	// cosmos authz
	MSG_TYPE_TO_EVENT_NAME_MAP["/cosmos.authz.v1beta1.MsgGrant"] = event.MSG_GRANT
	manager.RegisterParser(event.MSG_GRANT, BEGIN_BLOCK_HEIGHT, ParseMsgGrant)
	MSG_TYPE_TO_EVENT_NAME_MAP["/cosmos.authz.v1beta1.MsgRevoke"] = event.MSG_REVOKE
	manager.RegisterParser(event.MSG_REVOKE, BEGIN_BLOCK_HEIGHT, ParseMsgRevoke)
	MSG_TYPE_TO_EVENT_NAME_MAP["/cosmos.authz.v1beta1.MsgExec"] = event.MSG_EXEC
	manager.RegisterParser(event.MSG_EXEC, BEGIN_BLOCK_HEIGHT, ParseMsgExec)

	// cosmos feegrant
	MSG_TYPE_TO_EVENT_NAME_MAP["/cosmos.feegrant.v1beta1.MsgGrantAllowance"] = event.MSG_GRANT_ALLOWANCE
	manager.RegisterParser(event.MSG_GRANT_ALLOWANCE, BEGIN_BLOCK_HEIGHT, ParseMsgGrantAllowance)
	MSG_TYPE_TO_EVENT_NAME_MAP["/cosmos.feegrant.v1beta1.MsgRevokeAllowance"] = event.MSG_REVOKE_ALLOWANCE
	manager.RegisterParser(event.MSG_REVOKE_ALLOWANCE, BEGIN_BLOCK_HEIGHT, ParseMsgRevokeAllowance)

	// cosmos vesting
	MSG_TYPE_TO_EVENT_NAME_MAP["/cosmos.vesting.v1beta1.MsgCreateVestingAccount"] = event.MSG_CREATE_VESTING_ACCOUNT
	manager.RegisterParser(event.MSG_CREATE_VESTING_ACCOUNT, BEGIN_BLOCK_HEIGHT, ParseMsgCreateVestingAccount)
}

func RegisterBreakingVersionParsers(manager *utils.CosmosParserManager) {
	//v0.43.0
	manager.RegisterParser(event.MSG_IBC_RECV_PACKET, manager.GetCosmosV0_43_0BlockHeight(), V0_43_0_ibcmsg.ParseMsgRecvPacket)
}

// GetParserKeyFromMsgType panic if msgType is not registered
func GetParserKeyFromMsgType(msgType string) utils.CosmosParserKey {
	key, ok := MSG_TYPE_TO_EVENT_NAME_MAP[msgType]
	if !ok {
		panic(fmt.Sprintf("Requesting invalid parser key from type %s", msgType))
	}
	return key
}
