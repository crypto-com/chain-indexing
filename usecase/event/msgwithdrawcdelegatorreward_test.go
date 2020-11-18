package event_test

import (
	"github.com/crypto-com/chainindex/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chainindex/usecase/coin"
	event_usecase "github.com/crypto-com/chainindex/usecase/event"
)

var _ = Describe("Event", func() {
	Describe("En/DecodeMsgWithdrawDelegatorReward", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyDelegatorAddress := "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn"
			anyValidatorAddress := "tcrocncl15grftg88l0gdw4mg9t9pwnl0pde2asjzekz0ek"
			anyRecipientAddress := "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh"
			anyAmount := coin.MustNewCoinFromString("123456")
			anyParams := model.MsgWithdrawDelegatorRewardParams{
				DelegatorAddress: anyDelegatorAddress,
				ValidatorAddress: anyValidatorAddress,
				RecipientAddress: anyRecipientAddress,
				Amount:           anyAmount,
			}
			event := event_usecase.NewMsgWithdrawDelegatorReward(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := event_usecase.DecodeMsgWithdrawDelegatorReward([]byte(encoded))
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgWithdrawDelegatorReward)
			Expect(typedEvent.Name()).To(Equal("MsgWithdrawDelegatorRewardCreated"))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.TxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.DelegatorAddress).To(Equal(anyDelegatorAddress))
			Expect(typedEvent.ValidatorAddress).To(Equal(anyValidatorAddress))
			Expect(typedEvent.RecipientAddress).To(Equal(anyRecipientAddress))
			Expect(typedEvent.Amount).To(Equal(anyAmount))
		})

		It("should able to encode and decode to failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyDelegatorAddress := "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn"
			anyValidatorAddress := "tcrocncl15grftg88l0gdw4mg9t9pwnl0pde2asjzekz0ek"
			anyRecipientAddress := "tcro14m5a4kxt2e82uqqs5gtqza29dm5wqzya2jw9sh"
			anyAmount := coin.MustNewCoinFromString("123456")
			anyParams := model.MsgWithdrawDelegatorRewardParams{
				DelegatorAddress: anyDelegatorAddress,
				ValidatorAddress: anyValidatorAddress,
				RecipientAddress: anyRecipientAddress,
				Amount:           anyAmount,
			}
			event := event_usecase.NewMsgWithdrawDelegatorReward(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := event_usecase.DecodeMsgWithdrawDelegatorReward([]byte(encoded))
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgWithdrawDelegatorReward)
			Expect(typedEvent.Name()).To(Equal("MsgWithdrawDelegatorRewardFailed"))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.TxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.DelegatorAddress).To(Equal(anyDelegatorAddress))
			Expect(typedEvent.ValidatorAddress).To(Equal(anyValidatorAddress))
			Expect(typedEvent.RecipientAddress).To(Equal(anyRecipientAddress))
			Expect(typedEvent.Amount).To(Equal(anyAmount))
		})
	})
})
