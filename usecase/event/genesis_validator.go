package event

import (
	"bytes"

	"github.com/crypto-com/chain-indexing/projection/validator/constants"

	"github.com/crypto-com/chain-indexing/usecase/coin"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"github.com/crypto-com/chain-indexing/usecase/model/genesis"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const GENESIS_VALIDATOR_CREATED = "GenesisValidatorCreated"

type CreateGenesisValidator struct {
	event_entity.Base

	Status            constants.Status           `json:"status"`
	Description       model.ValidatorDescription `json:"description"`
	CommissionRates   model.ValidatorCommission  `json:"commissionRates"`
	MinSelfDelegation string                     `json:"minSelfDelegation"`
	DelegatorAddress  string                     `json:"delegatorAddress"`
	ValidatorAddress  string                     `json:"validatorAddress"`
	TendermintPubkey  string                     `json:"tendermintPubkey"`
	Amount            coin.Coin                  `json:"amount"`
	Jailed            bool                       `json:"jailed"`
}

func NewCreateGenesisValidator(
	params genesis.CreateGenesisValidatorParams,
) *CreateGenesisValidator {
	return &CreateGenesisValidator{
		event_entity.NewBase(event_entity.BaseParams{
			Name:        GENESIS_VALIDATOR_CREATED,
			Version:     1,
			BlockHeight: 0,
		}),

		params.Status,
		params.Description,
		params.Commission,
		params.MinSelfDelegation,
		params.DelegatorAddress,
		params.ValidatorAddress,
		params.TendermintPubkey,
		params.Amount,
		params.Jailed,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *CreateGenesisValidator) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *CreateGenesisValidator) String() string {
	return render.Render(event)
}

func DecodeCreateGenesisValidator(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *CreateGenesisValidator
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
