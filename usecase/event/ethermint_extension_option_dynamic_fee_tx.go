package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const DYNAMIC_FEE_TX = "/ethermint.evm.v1.DynamicFeeTx"
const DYNAMIC_FEE_TX_CREATED = "/ethermint.evm.v1.DynamicFeeTx.Created"
const DYNAMIC_FEE_TX_FAILED = "/ethermint.evm.v1.DynamicFeeTx.Failed"

type EthermintExtensionOptionDynamicFeeTx struct {
	MsgBase

	Params model.EthermintDynamicFeeTxParams `json:"params"`
}

func NewEthermintExtensionOptionDynamicFeeTx(
	msgCommonParams MsgCommonParams,
	params model.EthermintDynamicFeeTxParams,
) *EthermintExtensionOptionDynamicFeeTx {
	return &EthermintExtensionOptionDynamicFeeTx{
		NewMsgBase(MsgBaseParams{
			MsgName:         DYNAMIC_FEE_TX,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *EthermintExtensionOptionDynamicFeeTx) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *EthermintExtensionOptionDynamicFeeTx) String() string {
	return render.Render(event)
}

func DecodeEthermintExtensionOptionDynamicFeeTx(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *EthermintExtensionOptionDynamicFeeTx
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
