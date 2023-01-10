package event_test

import (
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	v1_model "github.com/crypto-com/chain-indexing/usecase/model/v1"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeMsgDepositV1", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(5066)
			anyTxHash := "8F549DCBED33FF8909C83C6FE957D262F833962ACFE70C53E969DE8F1D657B7C"
			anyMsgIndex := 0
			anyProposalId := "3"
			anyDepositor := "crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7"
			anyAmount := coin.MustParseCoinsNormalized("100000basecro")
			anyParams := v1_model.MsgDepositParams{
				ProposalId: anyProposalId,
				Depositor:  anyDepositor,
				Amount:     anyAmount,
			}
			event := event_usecase.NewMsgDepositV1(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_DEPOSIT_V1_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgDepositV1)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_DEPOSIT_V1_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.ProposalId).To(Equal(anyProposalId))
			Expect(typedEvent.Depositor).To(Equal(anyDepositor))
			Expect(typedEvent.Amount).To(Equal(anyAmount))
		})

		It("should able to encode and decode to failed event", func() {
			anyHeight := int64(5066)
			anyTxHash := "8F549DCBED33FF8909C83C6FE957D262F833962ACFE70C53E969DE8F1D657B7C"
			anyMsgIndex := 0
			anyProposalId := "3"
			anyDepositor := "crc18z6q38mhvtsvyr5mak8fj8s8g4gw7kjjtsgrn7"
			anyAmount := coin.MustParseCoinsNormalized("100000basecro, 10cro")
			anyParams := v1_model.MsgDepositParams{
				ProposalId: anyProposalId,
				Depositor:  anyDepositor,
				Amount:     anyAmount,
			}
			event := event_usecase.NewMsgDepositV1(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_DEPOSIT_V1_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgDepositV1)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_DEPOSIT_V1_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.ProposalId).To(Equal(anyProposalId))
			Expect(typedEvent.Depositor).To(Equal(anyDepositor))
			Expect(typedEvent.Amount).To(Equal(anyAmount))
		})
	})
})
