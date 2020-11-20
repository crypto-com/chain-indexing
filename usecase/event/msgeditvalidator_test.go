package event_test

import (
	event_entity "github.com/crypto-com/chainindex/entity/event"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chainindex/usecase/event"
	event_usecase "github.com/crypto-com/chainindex/usecase/event"
	"github.com/crypto-com/chainindex/usecase/model"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeMsgEditValidator", func() {
		It("should able to encode and decode to the same event", func() {

			event := event_usecase.NewMsgEditValidator(
				event.MsgCommonParams{
					BlockHeight: int64(503978),
					TxHash:      "E69985AC8168383A81B7952DBE03EB9B3400FF80AEC0F362369DD7F38B1C2FE9",
					TxSuccess:   true,
					MsgIndex:    0,
				},
				model.MsgEditValidatorParams{
					ValidatorAddress:  "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
					CommissionRate:    "0.1",
					MinSelfDelegation: "2",
				},
			)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_EDIT_VALIDATOR_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgEditValidator)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_EDIT_VALIDATOR_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.TxHash).To(Equal(event.TxHash))
			Expect(typedEvent.MsgIndex).To(Equal(event.MsgIndex))
			Expect(typedEvent.ValidatorAddress).To(Equal(event.ValidatorAddress))
			Expect(typedEvent.CommissionRate).To(Equal(event.CommissionRate))
			Expect(typedEvent.MinSelfDelegation).To(Equal(event.MinSelfDelegation))
		})

		It("should able to encode and decode to failed event", func() {
			event := event_usecase.NewMsgEditValidator(
				event.MsgCommonParams{
					BlockHeight: int64(503978),
					TxHash:      "E69985AC8168383A81B7952DBE03EB9B3400FF80AEC0F362369DD7F38B1C2FE9",
					TxSuccess:   false,
					MsgIndex:    0,
				},
				model.MsgEditValidatorParams{
					ValidatorAddress:  "tcrocncl1fmprm0sjy6lz9llv7rltn0v2azzwcwzvr4ufus",
					CommissionRate:    "0.1",
					MinSelfDelegation: "2",
				},
			)
			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_EDIT_VALIDATOR_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgEditValidator)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_EDIT_VALIDATOR_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
		})
	})
})
