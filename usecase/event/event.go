package event

import (
	"github.com/crypto-com/chainindex/entity/event"
)

func RegisterEvents(registry *event.Registry) {
	registry.Register(BLOCK_CREATED_NAME, 1, DecodeBlockCreated)
	registry.Register(RAW_BLOCK_CREATED_NAME, 1, DecodeRawBlockCreated)
	registry.Register(TRANSACTION_CREATED_NAME, 1, DecodeTransactionCreated)
	registry.Register(TRANSACTION_FAILED_NAME, 1, DecodeTransactionFailed)

	registry.Register(MSG_SEND_CREATED, 1, DecodeMsgSend)
	registry.Register(MSG_SEND_FAILED, 1, DecodeMsgSend)
	registry.Register(MSG_MULTI_SEND_CREATED, 1, DecodeMsgMultiSend)
	registry.Register(MSG_MULTI_SEND_FAILED, 1, DecodeMsgMultiSend)
	registry.Register(MSG_WITHDRAW_DELEGATOR_REWARD_CREATED, 1, DecodeMsgWithdrawDelegatorReward)
	registry.Register(MSG_WITHDRAW_DELEGATOR_REWARD_FAILED, 1, DecodeMsgWithdrawDelegatorReward)
}
