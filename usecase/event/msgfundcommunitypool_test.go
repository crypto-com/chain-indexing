package event_test

import (
	event_entity "github.com/crypto-com/chainindex/entity/event"
	"github.com/crypto-com/chainindex/usecase/coin"
	"github.com/crypto-com/chainindex/usecase/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_usecase "github.com/crypto-com/chainindex/usecase/event"
)

var _ = Describe("Event", func() {
	Describe("En/DecodeMsgFundCommunityPool", func() {
		registry := event_entity.NewRegistry()
		event_usecase.RegisterEvents(registry)

		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyDepositor := "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn"
			anyAmount := coin.MustNewCoinFromString("123456")
			anyParams := model.MsgFundCommunityPoolParams{
				Depositor: anyDepositor,
				Amount:    anyAmount,
			}
			event := event_usecase.NewMsgFundCommunityPool(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_FUND_COMMUNITY_POOL_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgFundCommunityPool)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_FUND_COMMUNITY_POOL_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.Depositor).To(Equal(anyDepositor))
			Expect(typedEvent.Amount).To(Equal(anyAmount))
		})

		It("should able to encode and decode to failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyDepositor := "tcro1fmprm0sjy6lz9llv7rltn0v2azzwcwzvk2lsyn"
			anyAmount := coin.MustNewCoinFromString("123456")
			anyParams := model.MsgFundCommunityPoolParams{
				Depositor: anyDepositor,
				Amount:    anyAmount,
			}
			event := event_usecase.NewMsgFundCommunityPool(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_FUND_COMMUNITY_POOL_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgFundCommunityPool)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_FUND_COMMUNITY_POOL_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.Depositor).To(Equal(anyDepositor))
			Expect(typedEvent.Amount).To(Equal(anyAmount))
		})
	})
})
