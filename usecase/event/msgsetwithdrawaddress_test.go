package event_test

import (
	event_entity "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_usecase "github.com/crypto-com/chainindex/usecase/event"
)

var _ = Describe("Event", func() {
	Describe("En/DecodeMsgSetWithdrawAddress", func() {
		registry := event_entity.NewRegistry()
		event_usecase.RegisterEvents(registry)

		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyDelegatorAddress := "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn"
			anyWithdrawAddress := "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh"
			anyParams := model.MsgSetWithdrawAddressParams{
				DelegatorAddress: anyDelegatorAddress,
				WithdrawAddress:  anyWithdrawAddress,
			}
			event := event_usecase.NewMsgSetWithdrawAddress(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_SET_WITHDRAW_ADDRESS_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgSetWithdrawAddress)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_SET_WITHDRAW_ADDRESS_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.TxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.DelegatorAddress).To(Equal(anyDelegatorAddress))
			Expect(typedEvent.WithdrawAddress).To(Equal(anyWithdrawAddress))
		})

		It("should able to encode and decode to failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyDelegatorAddress := "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn"
			anyWithdrawAddress := "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh"
			anyParams := model.MsgSetWithdrawAddressParams{
				DelegatorAddress: anyDelegatorAddress,
				WithdrawAddress:  anyWithdrawAddress,
			}
			event := event_usecase.NewMsgSetWithdrawAddress(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_SET_WITHDRAW_ADDRESS_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgSetWithdrawAddress)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_SET_WITHDRAW_ADDRESS_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.TxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.DelegatorAddress).To(Equal(anyDelegatorAddress))
			Expect(typedEvent.WithdrawAddress).To(Equal(anyWithdrawAddress))
		})
	})
})
