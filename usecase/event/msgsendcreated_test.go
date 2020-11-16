package event_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chainindex/usecase/coin"
	event_usecase "github.com/crypto-com/chainindex/usecase/event"
)

var _ = Describe("Event", func() {
	Describe("En/DecodeMsgSendCreated", func() {
		It("should able to encode and decode to the same Event", func() {
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
			event := event_usecase.NewMsgSendCreated(anyHeight, anyTxHash, anyMsgIndex, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			fmt.Println(encoded)

			decodedEvent, err := event_usecase.DecodeMsgSendCreated([]byte(encoded))
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent := decodedEvent.(*event_usecase.MsgSendCreated)
			Expect(typedEvent.TxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.FromAddress).To(Equal(anyFromAddress))
			Expect(typedEvent.ToAddress).To(Equal(anyToAddress))
			Expect(typedEvent.Amount).To(Equal(anyAmount))
		})
	})
})
