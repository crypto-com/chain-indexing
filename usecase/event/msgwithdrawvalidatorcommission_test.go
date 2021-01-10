package event_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

var _ = Describe("Event", func() {
	Describe("En/DecodeMsgWithdrawValidatorCommission", func() {
		registry := event_entity.NewRegistry()
		event_usecase.RegisterEvents(registry)

		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyValidatorAddress := "tcrocncl15grftg88l0gdw4mg9t9pwnl0pde2asjzekz0ek"
			anyRecipientAddress := "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh"
			anyAmount := coin.MustParseCoinsNormalized("123456basetcro,456789tcro")
			anyParams := model.MsgWithdrawValidatorCommissionParams{
				ValidatorAddress: anyValidatorAddress,
				RecipientAddress: anyRecipientAddress,
				Amount:           anyAmount,
			}
			event := event_usecase.NewMsgWithdrawValidatorCommission(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_WITHDRAW_VALIDATOR_COMMISSION_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgWithdrawValidatorCommission)
			Expect(typedEvent.Name()).To(Equal("MsgWithdrawValidatorCommissionCreated"))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.ValidatorAddress).To(Equal(anyValidatorAddress))
			Expect(typedEvent.RecipientAddress).To(Equal(anyRecipientAddress))
			Expect(typedEvent.Amount).To(Equal(anyAmount))
		})

		It("should able to encode and decode to failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyValidatorAddress := "tcrocncl15grftg88l0gdw4mg9t9pwnl0pde2asjzekz0ek"
			anyRecipientAddress := "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh"
			anyAmount := coin.MustParseCoinsNormalized("123456basetcro,456789tcro")
			anyParams := model.MsgWithdrawValidatorCommissionParams{
				ValidatorAddress: anyValidatorAddress,
				RecipientAddress: anyRecipientAddress,
				Amount:           anyAmount,
			}
			event := event_usecase.NewMsgWithdrawValidatorCommission(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_WITHDRAW_VALIDATOR_COMMISSION_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgWithdrawValidatorCommission)
			Expect(typedEvent.Name()).To(Equal("MsgWithdrawValidatorCommissionFailed"))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.ValidatorAddress).To(Equal(anyValidatorAddress))
			Expect(typedEvent.RecipientAddress).To(Equal(anyRecipientAddress))
			Expect(typedEvent.Amount).To(Equal(anyAmount))
		})
	})
})
