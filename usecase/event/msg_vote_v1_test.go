package event_test

import (
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	v1_model "github.com/crypto-com/chain-indexing/usecase/model/v1"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeMsgVoteV1", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyProposalId := "1"
			anyVoter := "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3"
			anyOption := "VOTE_OPTION_YES"
			anyMetadata := "Metadata"
			anyParams := v1_model.MsgVoteParams{
				ProposalId: anyProposalId,
				Voter:      anyVoter,
				Option:     anyOption,
				Metadata:   anyMetadata,
			}
			event := event_usecase.NewMsgVoteV1(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_VOTE_V1_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgVoteV1)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_VOTE_V1_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.ProposalId).To(Equal(anyProposalId))
			Expect(typedEvent.Voter).To(Equal(anyVoter))
			Expect(typedEvent.Option).To(Equal(anyOption))
			Expect(typedEvent.Metadata).To(Equal(anyMetadata))
		})

		It("should able to encode and decode to failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyProposalId := "1"
			anyVoter := "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3"
			anyOption := "VOTE_OPTION_YES"
			anyMetadata := "Metadata"
			anyParams := v1_model.MsgVoteParams{
				ProposalId: anyProposalId,
				Voter:      anyVoter,
				Option:     anyOption,
				Metadata:   anyMetadata,
			}
			event := event_usecase.NewMsgVoteV1(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_VOTE_V1_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgVoteV1)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_VOTE_V1_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.ProposalId).To(Equal(anyProposalId))
			Expect(typedEvent.Voter).To(Equal(anyVoter))
			Expect(typedEvent.Option).To(Equal(anyOption))
			Expect(typedEvent.Metadata).To(Equal(anyMetadata))
		})
	})
})
