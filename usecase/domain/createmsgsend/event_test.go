package createmsgsend_test

import (
	"fmt"

	"github.com/crypto-com/chainindex/usecase/coin"
	"github.com/crypto-com/chainindex/usecase/domain/createmsgsend"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Event", func() {
	Describe("En/DecodeEvent", func() {
		It("should able to encode and decode to the same Event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyFromAddress := "tcro165tzcrh2yl83g8qeqxueg2g5gzgu57y3fe3kc3"
			anyToAddress := "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3"
			anyAmount := coin.MustNewCoinFromString("123456")
			anyParams := createmsgsend.Params{
				FromAddress: anyFromAddress,
				ToAddress:   anyToAddress,
				Amount:      anyAmount,
			}
			event := createmsgsend.NewEvent(anyHeight, anyTxHash, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			fmt.Println(encoded)

			decodedEvent, err := createmsgsend.DecodeEvent([]byte(encoded))
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent := decodedEvent.(*createmsgsend.MsgSendCreated)
			Expect(typedEvent.TxHash).To(Equal(anyTxHash))
			Expect(typedEvent.FromAddress).To(Equal(anyFromAddress))
			Expect(typedEvent.ToAddress).To(Equal(anyToAddress))
			Expect(typedEvent.Amount).To(Equal(anyAmount))
		})
	})
})
