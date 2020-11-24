package event

import (
	"fmt"
	"strings"

	"github.com/crypto-com/chainindex/entity/event"
)

// MsgBase composes of Base except it has logical switch between succeeded and failed
type MsgBase struct {
	event.Base

	MsgName   string `json:"msgName"`
	MsgTxHash string `json:"txHash"`
	MsgIndex  int    `json:"msgIndex"`
}

func NewMsgBase(params MsgBaseParams) MsgBase {
	if strings.HasSuffix(params.MsgName, "Created") || strings.HasSuffix(params.MsgName, "Failed") {
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

func eventName(msgName string, txSuccess bool) string {
	if txSuccess {
		return fmt.Sprintf("%sCreated", msgName)
	}
	return fmt.Sprintf("%sFailed", msgName)
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
