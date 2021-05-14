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

	Describe("En/DecodeMsgNFTBurnNFT", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyTokenId := "tokenId"
			anyDenomId := "denomId"
			anySender := "cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd"
			anyParams := model.MsgNFTBurnNFTParams{
				DenomId: anyDenomId,
				TokenId: anyTokenId,
				Sender:  anySender,
			}
			event := event_usecase.NewMsgNFTBurnNFT(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_NFT_BURN_NFT_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgNFTBurnNFT)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_NFT_BURN_NFT_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.DenomId).To(Equal(anyDenomId))
			Expect(typedEvent.TokenId).To(Equal(anyTokenId))
			Expect(typedEvent.Sender).To(Equal(anySender))
		})

		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyTokenId := "tokenId"
			anyDenomId := "denomId"
			anySender := "cro1nk4rq3q46ltgjghxz80hy385p9uj0tf58apkcd"
			anyParams := model.MsgNFTBurnNFTParams{
				DenomId: anyDenomId,
				TokenId: anyTokenId,
				Sender:  anySender,
			}
			event := event_usecase.NewMsgNFTBurnNFT(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   false,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_NFT_BURN_NFT_FAILED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgNFTBurnNFT)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_NFT_BURN_NFT_FAILED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.DenomId).To(Equal(anyDenomId))
			Expect(typedEvent.TokenId).To(Equal(anyTokenId))
			Expect(typedEvent.Sender).To(Equal(anySender))
		})
	})
})
