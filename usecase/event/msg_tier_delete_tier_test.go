package event_test

import (
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

var _ = Describe("Event", func() {
	registry := event_entity.NewRegistry()
	event_usecase.RegisterEvents(registry)

	Describe("En/DecodeMsgTierDeleteTier", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyAuthority := "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3"
			anyId := uint32(1)
			anyParams := model.MsgDeleteTierParams{
				RawMsgDeleteTier: model.RawMsgDeleteTier{
					Authority: anyAuthority,
					Id:        anyId,
				},
			}
			event := event_usecase.NewMsgTierDeleteTier(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_TIER_DELETE_TIER_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgTierDeleteTier)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_TIER_DELETE_TIER_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.Authority).To(Equal(anyAuthority))
			Expect(typedEvent.Id).To(Equal(anyId))
		})

		It("should able to encode and decode to failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyParams := model.MsgDeleteTierParams{
				RawMsgDeleteTier: model.RawMsgDeleteTier{
					Authority: "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3",
					Id:        1,
				},
			}
			event := event_usecase.NewMsgTierDeleteTier(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_TIER_DELETE_TIER_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgTierDeleteTier)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_TIER_DELETE_TIER_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
		})
	})
})
