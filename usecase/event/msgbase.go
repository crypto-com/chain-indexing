package event

import (
	"fmt"
	"strings"

	"github.com/crypto-com/chainindex/entity/event"
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
			Name:        eventName(params.MsgName, params.TxSuccess),
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

func eventName(msgName string, txSuccess bool) string {
	var suffix string
	if txSuccess {
		suffix = MSG_SUCCESS_SUFFIX
	} else {
		suffix = MSG_FAILED_SUFFIX
	}
	return fmt.Sprintf("%s%s", msgName, suffix)
}

type MsgBaseParams struct {
	MsgName string
	Version int

	MsgCommonParams
}

type MsgCommonParams struct {
	BlockHeight int64
	TxHash      string
	TxSuccess   bool
	MsgIndex    int
}
