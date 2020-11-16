package event_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chainindex/usecase/coin"
	event_usecase "github.com/crypto-com/chainindex/usecase/event"
)

var _ = Describe("Event", func() {
	Describe("En/DecodeMsgSend", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyFromAddress := "tcro165tzcrh2yl83g8qeqxueg2g5gzgu57y3fe3kc3"
			anyToAddress := "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3"
			anyAmount := coin.MustNewCoinFromString("123456")
			anyParams := event_usecase.MsgSendCreatedParams{
				FromAddress: anyFromAddress,
				ToAddress:   anyToAddress,
				Amount:      anyAmount,
			}
			event := event_usecase.NewMsgSend(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := event_usecase.DecodeMsgSend([]byte(encoded))
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgSend)
			Expect(typedEvent.Name()).To(Equal("MsgSendCreated"))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.TxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.FromAddress).To(Equal(anyFromAddress))
			Expect(typedEvent.ToAddress).To(Equal(anyToAddress))
			Expect(typedEvent.Amount).To(Equal(anyAmount))
		})

		It("should able to encode and decode to failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyFromAddress := "tcro165tzcrh2yl83g8qeqxueg2g5gzgu57y3fe3kc3"
			anyToAddress := "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3"
			anyAmount := coin.MustNewCoinFromString("123456")
			anyParams := event_usecase.MsgSendCreatedParams{
				FromAddress: anyFromAddress,
				ToAddress:   anyToAddress,
				Amount:      anyAmount,
			}
			event := event_usecase.NewMsgSend(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := event_usecase.DecodeMsgSend([]byte(encoded))
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgSend)
			Expect(typedEvent.Name()).To(Equal("MsgSendFailed"))
			Expect(typedEvent.Version()).To(Equal(1))
		})
	})
})
