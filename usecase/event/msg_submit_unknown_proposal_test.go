package event_test

import (
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ = Describe("Event", func() {
	Describe("En/DecodeMsgSubmitProposal", func() {
		registry := event_entity.NewRegistry()
		event_usecase.RegisterEvents(registry)

		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyProposerAddress := "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn"
			anyContent := model.MsgSubmitUnknownProposalContent{
				Type:        "/cosmos.gov.v1beta1.UnknownProposal",
				Title:       "any title",
				Description: "any description",
				RawContent:  "{\"@type\":\"/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal\",\"title\":\"anytitle\",\"description\":\"any description\",\"plan\":{\"name\":\"testing\",\"time\":\"0001-01-01T00:00:00Z\",\"height\":\"10020\",\"info\":\"test\",\"upgraded_client_state\":null}}",
			}
			anyInitialDeposit := coin.MustParseCoinsNormalized("1000basetcro,2000tcro")
			anyParams := model.MsgSubmitUnknownProposalParams{
				ProposerAddress: anyProposerAddress,
				Content:         anyContent,
				InitialDeposit:  anyInitialDeposit,
			}
			event := event_usecase.NewMsgSubmitUnknownProposal(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_SUBMIT_UNKNOWN_PROPOSAL_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgSubmitUnknownProposal)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_SUBMIT_UNKNOWN_PROPOSAL_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.ProposerAddress).To(Equal(anyProposerAddress))
			Expect(typedEvent.InitialDeposit).To(Equal(anyInitialDeposit))
		})

		It("should able to encode and decode to failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyProposerAddress := "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn"
			anyContent := model.MsgSubmitUnknownProposalContent{
				Type:        "/cosmos.gov.v1beta1.UnknownProposal",
				Title:       "any title",
				Description: "any description",
				RawContent:  "{\"@type\":\"/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal\",\"title\":\"anytitle\",\"description\":\"any description\",\"plan\":{\"name\":\"testing\",\"time\":\"0001-01-01T00:00:00Z\",\"height\":\"10020\",\"info\":\"test\",\"upgraded_client_state\":null}}",
			}
			anyInitialDeposit := coin.MustParseCoinsNormalized("1000basetcro,2000tcro")
			anyParams := model.MsgSubmitUnknownProposalParams{
				ProposerAddress: anyProposerAddress,
				Content:         anyContent,
				InitialDeposit:  anyInitialDeposit,
			}
			event := event_usecase.NewMsgSubmitUnknownProposal(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_SUBMIT_UNKNOWN_PROPOSAL_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgSubmitUnknownProposal)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_SUBMIT_UNKNOWN_PROPOSAL_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.ProposerAddress).To(Equal(anyProposerAddress))
			Expect(typedEvent.Content).To(Equal(anyContent))
			Expect(typedEvent.InitialDeposit).To(Equal(anyInitialDeposit))
		})
	})
})
