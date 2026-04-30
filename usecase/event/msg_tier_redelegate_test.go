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

	Describe("En/DecodeMsgTierRedelegate", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyOwner := "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3"
			anyPositionId := "123"
			anyDstValidator := "tcrocncl1kkqxv3szgh099xezt7y38t5anqzue4s3y3kgqe"
			anyParams := model.MsgTierRedelegateParams{
				RawMsgTierRedelegate: model.RawMsgTierRedelegate{
					Owner:        anyOwner,
					PositionId:   anyPositionId,
					DstValidator: anyDstValidator,
				},
			}
			event := event_usecase.NewMsgTierRedelegate(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_TIER_REDELEGATE_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgTierRedelegate)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_TIER_REDELEGATE_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.Owner).To(Equal(anyOwner))
			Expect(typedEvent.PositionId).To(Equal(anyPositionId))
			Expect(typedEvent.DstValidator).To(Equal(anyDstValidator))
		})

		It("should able to encode and decode to failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyParams := model.MsgTierRedelegateParams{
				RawMsgTierRedelegate: model.RawMsgTierRedelegate{
					Owner:        "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3",
					PositionId:   "123",
					DstValidator: "tcrocncl1kkqxv3szgh099xezt7y38t5anqzue4s3y3kgqe",
				},
			}
			event := event_usecase.NewMsgTierRedelegate(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_TIER_REDELEGATE_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgTierRedelegate)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_TIER_REDELEGATE_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
		})
	})
})
