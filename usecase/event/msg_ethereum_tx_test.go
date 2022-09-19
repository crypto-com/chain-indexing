package event_test

import (
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeMsgEthereumTx", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "2678437368AFC7E0E6D891D858F17B9C05CFEE850A786592A11992813D6A89FD"
			anyMsgIndex := 2

			anyParams := model.MsgLegacyTxParams{}

			event := event_usecase.NewMsgEthereumTx(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_ETHEREUM_TX_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgEthereumTx)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_ETHEREUM_TX_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeTrue())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
		})

		It("should able to encode and decode failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "2678437368AFC7E0E6D891D858F17B9C05CFEE850A786592A11992813D6A89FD"
			anyMsgIndex := 2

			anyParams := model.MsgLegacyTxParams{}

			event := event_usecase.NewMsgEthereumTx(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_ETHEREUM_TX_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgEthereumTx)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_ETHEREUM_TX_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
			Expect(typedEvent.TxSuccess()).To(BeFalse())
			Expect(typedEvent.TxHash()).To(Equal(anyTxHash))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
		})
	})
})
