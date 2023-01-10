package event_test

import (
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/external/primptr"
	v1_model "github.com/crypto-com/chain-indexing/usecase/model/v1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ = Describe("Event", func() {
	Describe("En/DecodeMsgCancelUpgrade", func() {
		registry := event_entity.NewRegistry()
		event_usecase.RegisterEvents(registry)

		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			maybeProposalId := primptr.String("1")
			anyAuthority := "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn"
			anyProposer := "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn"
			anyInitialDeposit := coin.MustParseCoinsNormalized("1000basetcro,2000tcro")
			anyMetadata := "Metadata"

			anyParams := v1_model.MsgCancelUpgradeParams{
				MaybeProposalId: maybeProposalId,
				Authority:       anyAuthority,
				Proposer:        anyProposer,
				InitialDeposit:  anyInitialDeposit,
				Metadata:        anyMetadata,
			}

			event := event_usecase.NewMsgCancelUpgrade(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_CANCEL_UPGRADE_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgCancelUpgrade)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_CANCEL_UPGRADE_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.MaybeProposalId).To(Equal(maybeProposalId))
			Expect(typedEvent.Authority).To(Equal(anyAuthority))
			Expect(typedEvent.Proposer).To(Equal(anyProposer))
			Expect(typedEvent.InitialDeposit).To(Equal(anyInitialDeposit))
			Expect(typedEvent.Metadata).To(Equal(anyMetadata))
		})

		It("should able to encode and decode to failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			maybeProposalId := primptr.String("1")
			anyAuthority := "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn"
			anyProposer := "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn"
			anyInitialDeposit := coin.MustParseCoinsNormalized("1000basetcro,2000tcro")
			anyMetadata := "Metadata"

			anyParams := v1_model.MsgCancelUpgradeParams{
				MaybeProposalId: maybeProposalId,
				Authority:       anyAuthority,
				Proposer:        anyProposer,
				InitialDeposit:  anyInitialDeposit,
				Metadata:        anyMetadata,
			}

			event := event_usecase.NewMsgCancelUpgrade(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_CANCEL_UPGRADE_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgCancelUpgrade)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_CANCEL_UPGRADE_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.MaybeProposalId).To(Equal(maybeProposalId))
			Expect(typedEvent.Authority).To(Equal(anyAuthority))
			Expect(typedEvent.Proposer).To(Equal(anyProposer))
			Expect(typedEvent.InitialDeposit).To(Equal(anyInitialDeposit))
			Expect(typedEvent.Metadata).To(Equal(anyMetadata))
		})
	})
})
