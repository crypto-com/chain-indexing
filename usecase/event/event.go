package event

import (
	"github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/usecase/event/blockcreated"
	"github.com/crypto-com/chainindex/usecase/event/rawblockcreated"
)

func RegisterEvents(registry *event.Registry) {
	registry.Register(blockcreated.NAME, 0, blockcreated.Decode)
	registry.Register(rawblockcreated.NAME, 0, rawblockcreated.Decode)
}
