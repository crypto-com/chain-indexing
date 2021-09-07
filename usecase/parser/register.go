package parser

import (
	"github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/parser/ibcmsg"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"
	V0_43_0_ibcmsg "github.com/crypto-com/chain-indexing/usecase/parser/v0_43_0/ibcmsg"
)

const BEGIN_BLOCK_HEIGHT = 0

func InitParsers(manager *utils.CosmosParserManager) {
	manager.RegisterParser(event.MSG_GRANT, BEGIN_BLOCK_HEIGHT, ParseMsgGrant)
	manager.RegisterParser(event.MSG_IBC_RECV_PACKET, BEGIN_BLOCK_HEIGHT, ibcmsg.ParseMsgRecvPacket)
}

func RegisterBreakingVersionParsers(manager *utils.CosmosParserManager) {
	//v0.43.0
	manager.RegisterParser(event.MSG_IBC_RECV_PACKET, manager.GetConfig().CosmosVersionBLockHeight.V0_43_0, V0_43_0_ibcmsg.ParseMsgRecvPacket)
}
