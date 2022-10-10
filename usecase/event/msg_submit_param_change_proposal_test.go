package event_test

import (
	"encoding/json"

	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ = Describe("Event", func() {
	Describe("En/DecodeMsgSubmitParamChangeProposal", func() {
		registry := event_entity.NewRegistry()
		event_usecase.RegisterEvents(registry)

		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyProposalId := "1"
			anyProposerAddress := "TCRO1FMPRM0SJY6LZ9LLV7RLTN0V2AZZWCWZVK2LSYN"
			anyContent := model.MsgSubmitParamChangeProposalContent{
				Type:        "/cosmos.params.v1beta1.ParameterChangeProposal",
				Title:       "Param-Change Voting Period",
				Description: "This is to change the voting time on Testnet to be 8 hours.",
				Changes: []model.MsgSubmitParamChangeProposalChange{
					{
						Subspace: "gov",
						Key:      "votingparams",
						Value:    json.RawMessage("\"{ \\\"voting_period\\\": \\\"28800000000000\\\" }\""),
					},
				},
			}
			anyInitialDeposit := coin.MustParseCoinsNormalized("1000basetcro,2000tcro")
			anyParams := model.MsgSubmitParamChangeProposalParams{
				MaybeProposalId: &anyProposalId,
				ProposerAddress: anyProposerAddress,
				Content:         anyContent,
				InitialDeposit:  anyInitialDeposit,
			}
			event := event_usecase.NewMsgSubmitParamChangeProposal(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_SUBMIT_PARAM_CHANGE_PROPOSAL_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgSubmitParamChangeProposal)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_SUBMIT_PARAM_CHANGE_PROPOSAL_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.MaybeProposalId).To(Equal(&anyProposalId))
			Expect(typedEvent.ProposerAddress).To(Equal(anyProposerAddress))
			Expect(typedEvent.Content).To(Equal(anyContent))
			Expect(typedEvent.InitialDeposit).To(Equal(anyInitialDeposit))
		})

		It("should able to encode and decode to failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyProposerAddress := "TCRO1FMPRM0SJY6LZ9LLV7RLTN0V2AZZWCWZVK2LSYN"
			anyContent := model.MsgSubmitParamChangeProposalContent{
				Type:        "/cosmos.params.v1beta1.ParameterChangeProposal",
				Title:       "Param-Change Voting Period",
				Description: "This is to change the voting time on Testnet to be 8 hours.",
				Changes: []model.MsgSubmitParamChangeProposalChange{
					{
						Subspace: "gov",
						Key:      "votingparams",
						Value:    json.RawMessage("\"{ \\\"voting_period\\\": \\\"28800000000000\\\" }\""),
					},
				},
			}
			anyInitialDeposit := coin.MustParseCoinsNormalized("1000basetcro,2000tcro")
			anyParams := model.MsgSubmitParamChangeProposalParams{
				ProposerAddress: anyProposerAddress,
				Content:         anyContent,
				InitialDeposit:  anyInitialDeposit,
			}
			event := event_usecase.NewMsgSubmitParamChangeProposal(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_SUBMIT_PARAM_CHANGE_PROPOSAL_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgSubmitParamChangeProposal)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_SUBMIT_PARAM_CHANGE_PROPOSAL_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.ProposerAddress).To(Equal(anyProposerAddress))
			Expect(typedEvent.Content).To(Equal(anyContent))
			Expect(typedEvent.InitialDeposit).To(Equal(anyInitialDeposit))
		})
	})
})
