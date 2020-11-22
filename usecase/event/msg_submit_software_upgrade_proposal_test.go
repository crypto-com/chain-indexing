package event_test

import (
	event_entity "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/internal/utctime"
	"github.com/crypto-com/chainindex/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/crypto-com/chainindex/usecase/coin"
	event_usecase "github.com/crypto-com/chainindex/usecase/event"
)

var _ = Describe("Event", func() {
	Describe("En/DecodeMsgSubmitSoftwareUpgradeProposal", func() {
		registry := event_entity.NewRegistry()
		event_usecase.RegisterEvents(registry)

		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyProposerAddress := "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn"
			anyContent := model.MsgSubmitSoftwareUpgradeProposalContent{
				Type:        "/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal",
				Title:       "any time",
				Description: "any description",
				Plan: model.MsgSubmitSoftwareUpgradeProposalPlan{
					Name:   "any name",
					Time:   utctime.FromUnixNano(int64(1000)),
					Height: 0,
					Info:   "any info",
				},
			}
			anyInitialDeposit := coin.MustNewCoinFromString("1000")
			anyParams := model.MsgSubmitSoftwareUpgradeProposalParams{
				ProposerAddress: anyProposerAddress,
				Content:         anyContent,
				InitialDeposit:  anyInitialDeposit,
			}
			event := event_usecase.NewMsgSubmitSoftwareUpgradeProposal(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgSubmitSoftwareUpgradeProposal)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.ProposerAddress).To(Equal(anyProposerAddress))
			Expect(typedEvent.Content).To(Equal(anyContent))
			Expect(typedEvent.InitialDeposit).To(Equal(anyInitialDeposit))
		})

		It("should able to encode and decode to failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyProposerAddress := "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn"
			anyContent := model.MsgSubmitSoftwareUpgradeProposalContent{
				Type:        "/cosmos.upgrade.v1beta1.SoftwareUpgradeProposal",
				Title:       "any time",
				Description: "any description",
				Plan: model.MsgSubmitSoftwareUpgradeProposalPlan{
					Name:   "any name",
					Time:   utctime.FromUnixNano(int64(1000)),
					Height: 0,
					Info:   "any info",
				},
			}
			anyInitialDeposit := coin.MustNewCoinFromString("1000")
			anyParams := model.MsgSubmitSoftwareUpgradeProposalParams{
				ProposerAddress: anyProposerAddress,
				Content:         anyContent,
				InitialDeposit:  anyInitialDeposit,
			}
			event := event_usecase.NewMsgSubmitSoftwareUpgradeProposal(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgSubmitSoftwareUpgradeProposal)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_SUBMIT_SOFTWARE_UPGRADE_PROPOSAL_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.ProposerAddress).To(Equal(anyProposerAddress))
			Expect(typedEvent.Content).To(Equal(anyContent))
			Expect(typedEvent.InitialDeposit).To(Equal(anyInitialDeposit))
		})
	})
})
