package domain

import (
	"github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/usecase/domain/createblock"
	"github.com/crypto-com/chainindex/usecase/domain/createmsgsend"
	"github.com/crypto-com/chainindex/usecase/domain/createrawblock"
	"github.com/crypto-com/chainindex/usecase/domain/createtransaction"
)

func RegisterEvents(registry *event.Registry) {
	registry.Register(createblock.EVENT_NAME, 1, createblock.DecodeEvent)
	registry.Register(createrawblock.EVENT_NAME, 1, createrawblock.DecodeEvent)
	registry.Register(createtransaction.EVENT_NAME, 1, createtransaction.DecodeEvent)

	registry.Register(createmsgsend.EVENT_NAME, 1, createmsgsend.DecodeEvent)
}
