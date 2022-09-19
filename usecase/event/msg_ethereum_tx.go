package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_ETHEREUM_TX = "/ethermint.evm.v1.MsgEthereumTx"
const MSG_ETHEREUM_TX_CREATED = "/ethermint.evm.v1.MsgEthereumTx.Created"
const MSG_ETHEREUM_TX_FAILED = "/ethermint.evm.v1.MsgEthereumTx.Failed"

type MsgEthereumTx struct {
	MsgBase

	Params model.MsgLegacyTxParams `json:"params"`
}

func NewMsgEthereumTx(
	msgCommonParams MsgCommonParams,
	params model.MsgLegacyTxParams,
) *MsgEthereumTx {
	return &MsgEthereumTx{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_ETHEREUM_TX,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgEthereumTx) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgEthereumTx) String() string {
	return render.Render(event)
}

func DecodeMsgEthereumTx(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgEthereumTx
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
