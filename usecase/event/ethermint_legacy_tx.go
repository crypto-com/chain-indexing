package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/usecase/model"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const LEGACY_TX = "/ethermint.evm.v1.LegacyTx"
const LEGACY_TX_CREATED = "/ethermint.evm.v1.LegacyTx.Created"
const LEGACY_TX_FAILED = "/ethermint.evm.v1.LegacyTx.Failed"

type EthermintLegacyTx struct {
	MsgBase

	Params model.EthermintLegacyTxParams `json:"params"`
}

func NewEthermintLegacyx(
	msgCommonParams MsgCommonParams,
	params model.EthermintLegacyTxParams,
) *EthermintLegacyTx {
	return &EthermintLegacyTx{
		NewMsgBase(MsgBaseParams{
			MsgName:         LEGACY_TX,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),

		params,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *EthermintLegacyTx) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *EthermintLegacyTx) String() string {
	return render.Render(event)
}

func DecodeEthermintLegacyTx(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *EthermintLegacyTx
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
