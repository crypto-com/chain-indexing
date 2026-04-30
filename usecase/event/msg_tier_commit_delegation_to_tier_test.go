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

	Describe("En/DecodeMsgTierCommitDelegationToTier", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyDelegatorAddress := "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3"
			anyValidatorAddress := "tcrocncl1kkqxv3szgh099xezt7y38t5anqzue4s3y3kgqe"
			anyAmount := "100000basecro"
			anyId := uint32(1)
			anyTriggerExitImmediately := true
			anyParams := model.MsgCommitDelegationToTierParams{
				RawMsgCommitDelegationToTier: model.RawMsgCommitDelegationToTier{
					DelegatorAddress:       anyDelegatorAddress,
					ValidatorAddress:       anyValidatorAddress,
					Amount:                 anyAmount,
					Id:                     anyId,
					TriggerExitImmediately: anyTriggerExitImmediately,
				},
			}
			event := event_usecase.NewMsgTierCommitDelegationToTier(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_TIER_COMMIT_DELEGATION_TO_TIER_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgTierCommitDelegationToTier)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_TIER_COMMIT_DELEGATION_TO_TIER_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.DelegatorAddress).To(Equal(anyDelegatorAddress))
			Expect(typedEvent.ValidatorAddress).To(Equal(anyValidatorAddress))
			Expect(typedEvent.Amount).To(Equal(anyAmount))
			Expect(typedEvent.Id).To(Equal(anyId))
			Expect(typedEvent.TriggerExitImmediately).To(Equal(anyTriggerExitImmediately))
		})

		It("should able to encode and decode to failed event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyParams := model.MsgCommitDelegationToTierParams{
				RawMsgCommitDelegationToTier: model.RawMsgCommitDelegationToTier{
					DelegatorAddress:       "tcro184lta2lsyu47vwyp2e8zmtca3k5yq85p6c4vp3",
					ValidatorAddress:       "tcrocncl1kkqxv3szgh099xezt7y38t5anqzue4s3y3kgqe",
					Amount:                 "100000basecro",
					Id:                     1,
					TriggerExitImmediately: true,
				},
			}
			event := event_usecase.NewMsgTierCommitDelegationToTier(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_TIER_COMMIT_DELEGATION_TO_TIER_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgTierCommitDelegationToTier)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_TIER_COMMIT_DELEGATION_TO_TIER_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))
		})
	})
})
