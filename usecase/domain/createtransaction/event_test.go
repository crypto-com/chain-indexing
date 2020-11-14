package createtransaction_test

import (
	"github.com/crypto-com/chainindex/usecase/coin"

	"github.com/crypto-com/chainindex/test/factory"
	"github.com/crypto-com/chainindex/usecase/domain/createtransaction"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Event", func() {
	Describe("En/DecodeEvent", func() {
		It("should able to encode and decode to the same Event", func() {
			anyHeight := int64(1000)
			anyParams := createtransaction.Params{
				TxHash:    factory.RandomTxHash(),
				Code:      0,
				Log:       "{\"events\":[]}",
				MsgCount:  1,
				Fee:       coin.MustNewCoinFromString("1000"),
				GasWanted: "200000",
				GasUsed:   "10000",
			}
			event := createtransaction.NewEvent(anyHeight, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := createtransaction.DecodeEvent([]byte(encoded))
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
		})
	})
})
