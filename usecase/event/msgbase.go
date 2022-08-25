package event

import (
	"fmt"
	"strings"

	"github.com/crypto-com/chain-indexing/entity/event"
)

const MSG_SUCCESS_SUFFIX = "Created"
const MSG_FAILED_SUFFIX = "Failed"

// MsgBase composes of Base except it has logical switch between succeeded and failed
type MsgBase struct {
	event.Base

	MsgName   string `json:"msgName"`
	MsgTxHash string `json:"txHash"`
	MsgIndex  int    `json:"msgIndex"`
}

func NewMsgBase(params MsgBaseParams) MsgBase {
	if strings.HasSuffix(params.MsgName, MSG_SUCCESS_SUFFIX) || strings.HasSuffix(params.MsgName, MSG_FAILED_SUFFIX) {
		panic("msg name should not have Created or Failed keyword")
	}

	return MsgBase{
		event.NewBase(event.BaseParams{
			Name:        eventName(params.MsgSuccess, params.MsgFailed, params.TxSuccess),
			Version:     params.Version,
			BlockHeight: params.BlockHeight,
		}),

		params.MsgName,
		params.TxHash,
		params.MsgIndex,
	}
}

func (base *MsgBase) MsgType() string {
	return base.MsgName
}

func (base *MsgBase) TxHash() string {
	return base.MsgTxHash
}

func (base *MsgBase) TxSuccess() bool {
	return strings.HasSuffix(base.Name(), MSG_SUCCESS_SUFFIX)
}

func eventName(msgSuccess string, msgFailed string, txSuccess bool) string {
	var eventName string
	if txSuccess {
		eventName = msgSuccess
	} else {
		eventName = msgFailed
	}
	return fmt.Sprintf("%s", eventName)
}

type MsgBaseParams struct {
	MsgName    string
	MsgType    string
	MsgSuccess string
	MsgFailed  string
	Version    int

	MsgCommonParams
}

type MsgCommonParams struct {
	BlockHeight int64
	TxHash      string
	TxSuccess   bool
	MsgIndex    int
}
