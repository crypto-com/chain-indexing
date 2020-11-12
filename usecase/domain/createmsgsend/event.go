package createmsgsend

import (
	entity_event "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/usecase"
)

type MsgSendCreated struct {
	entity_event.Base

	FromAddress string
	ToAddress   string
	Amount      usecase.Coin
}
