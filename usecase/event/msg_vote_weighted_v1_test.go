package event_test

import (
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	model_gov_v1 "github.com/crypto-com/chain-indexing/usecase/model/gov/v1"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeMsgVoteWeightedV1", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyProposalId := "1"
			anyVoter := "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3"
			anyVoteOptions := []model_gov_v1.VoteOption{
				{
					Option: "VOTE_OPTION_YES",
					Weight: "0.2",
				},
				{
					Option: "VOTE_OPTION_NO",
					Weight: "0.8",
				},
			}
			anyMetadata := "Metadata"
			anyParams := model_gov_v1.MsgVoteWeightedParams{
				ProposalId:  anyProposalId,
				Voter:       anyVoter,
				VoteOptions: anyVoteOptions,
				Metadata:    anyMetadata,
			}
			event := event_usecase.NewMsgVoteWeightedV1(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_VOTE_WEIGHTED_V1_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgVoteWeightedV1)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_VOTE_WEIGHTED_V1_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.ProposalId).To(Equal(anyProposalId))
			Expect(typedEvent.Voter).To(Equal(anyVoter))
			Expect(typedEvent.VoteOptions).To(Equal(anyVoteOptions))
			Expect(typedEvent.Metadata).To(Equal(anyMetadata))
		})

		It("should able to encode and decode to failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyProposalId := "1"
			anyVoter := "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3"
			anyVoteOptions := []model_gov_v1.VoteOption{
				{
					Option: "VOTE_OPTION_YES",
					Weight: "0.2",
				},
				{
					Option: "VOTE_OPTION_NO",
					Weight: "0.8",
				},
			}
			anyMetadata := "Metadata"
			anyParams := model_gov_v1.MsgVoteWeightedParams{
				ProposalId:  anyProposalId,
				Voter:       anyVoter,
				VoteOptions: anyVoteOptions,
				Metadata:    anyMetadata,
			}
			event := event_usecase.NewMsgVoteWeightedV1(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_VOTE_WEIGHTED_V1_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgVoteWeightedV1)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_VOTE_WEIGHTED_V1_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.ProposalId).To(Equal(anyProposalId))
			Expect(typedEvent.Voter).To(Equal(anyVoter))
			Expect(typedEvent.VoteOptions).To(Equal(anyVoteOptions))
			Expect(typedEvent.Metadata).To(Equal(anyMetadata))
		})
	})
})
