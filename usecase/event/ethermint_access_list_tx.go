package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const ACCESS_LIST_TX = "/ethermint.evm.v1.AccessListTx"
const ACCESS_LIST_TX_CREATED = "/ethermint.evm.v1.AccessListTx.Created"
const ACCESS_LIST_TX_FAILED = "/ethermint.evm.v1.AccessListTx.Failed"

type EthermintAccessListTx struct {
	MsgBase

	Params model.EthermintAccessListTxParams `json:"params"`
}

func NewEthermintAccessListTx(
	msgCommonParams MsgCommonParams,
	params model.EthermintAccessListTxParams,
) *EthermintAccessListTx {
	return &EthermintAccessListTx{
		NewMsgBase(MsgBaseParams{
			MsgName:         ACCESS_LIST_TX,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *EthermintAccessListTx) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *EthermintAccessListTx) String() string {
	return render.Render(event)
}

func DecodeEthermintAccessListTx(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *EthermintAccessListTx
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
