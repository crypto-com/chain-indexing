package event

import (
	"github.com/crypto-com/chainindex/entity/event"
)

func RegisterEvents(registry *event.Registry) {
	registry.Register(BLOCK_CREATED_NAME, 1, DecodeBlockCreated)
	registry.Register(RAW_BLOCK_CREATED_NAME, 1, DecodeRawBlockCreated)
	registry.Register(TRANSACTION_CREATED_NAME, 1, DecodeTransactionCreated)
	registry.Register(TRANSACTION_FAILED_NAME, 1, DecodeTransactionFailed)

	registry.Register(MSG_SEND_CREATED_NAME, 1, DecodeMsgSendCreated)
}
