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

	Describe("En/DecodeMsgNFTTransferNFT", func() {
		It("should able to encode and decode to the same event", func() {
			anyHeight := int64(1000)
			anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			anyMsgIndex := 2
			anyTokenId := "tokenId"
			anyDenomId := "denomId"
			anySender := "sender"
			anyRecipient := "recipient"
			anyParams := model.MsgNFTTransferNFTParams{
				DenomId:   anyDenomId,
				TokenId:   anyTokenId,
				Sender:    anySender,
				Recipient: anyRecipient,
			}
			event := event_usecase.NewMsgNFTTransferNFT(event_usecase.MsgCommonParams{
				BlockHeight: anyHeight,
				TxHash:      anyTxHash,
				TxSuccess:   true,
				MsgIndex:    anyMsgIndex,
			}, anyParams)

			encoded, err := event.ToJSON()
			Expect(err).To(BeNil())

			decodedEvent, err := registry.DecodeByType(
				event_usecase.MSG_NFT_TRANSFER_NFT_CREATED, 1, []byte(encoded),
			)
			Expect(err).To(BeNil())
			Expect(decodedEvent).To(Equal(event))
			typedEvent, _ := decodedEvent.(*event_usecase.MsgNFTTransferNFT)
			Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_NFT_TRANSFER_NFT_CREATED))
			Expect(typedEvent.Version()).To(Equal(1))

			Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			Expect(typedEvent.DenomId).To(Equal(anyDenomId))
			Expect(typedEvent.TokenId).To(Equal(anyTokenId))
			Expect(typedEvent.Sender).To(Equal(anySender))
			Expect(typedEvent.Recipient).To(Equal(anyRecipient))
		})

		It("should able to encode and decode to the same event", func() {
			//anyHeight := int64(1000)
			//anyTxHash := "4936522F7391D425F2A93AD47576F8AEC3947DC907113BE8A2FBCFF8E9F2A416"
			//anyMsgIndex := 2
			//anyTokenId := "tokenId"
			//anyDenomId := "denomId"
			//anyTokenName := "name"
			//anyURI := "uri"
			//anyData := "data"
			//anyParams := model.MsgNFTMintNFTParams{
			//	DenomId:   anyDenomId,
			//	TokenId:   anyTokenId,
			//	TokenName: anyTokenName,
			//	URI:       anyURI,
			//	Data:      anyData,
			//}
			//event := event_usecase.NewMsgNFTMintNFT(event_usecase.MsgCommonParams{
			//	BlockHeight: anyHeight,
			//	TxHash:      anyTxHash,
			//	TxSuccess:   false,
			//	MsgIndex:    anyMsgIndex,
			//}, anyParams)
			//
			//encoded, err := event.ToJSON()
			//Expect(err).To(BeNil())
			//
			//decodedEvent, err := registry.DecodeByType(
			//	event_usecase.MSG_NFT_MINT_NFT_FAILED, 1, []byte(encoded),
			//)
			//Expect(err).To(BeNil())
			//Expect(decodedEvent).To(Equal(event))
			//typedEvent, _ := decodedEvent.(*event_usecase.MsgNFTMintNFT)
			//Expect(typedEvent.Name()).To(Equal(event_usecase.MSG_NFT_MINT_NFT_FAILED))
			//Expect(typedEvent.Version()).To(Equal(1))
			//
			//Expect(typedEvent.MsgTxHash).To(Equal(anyTxHash))
			//Expect(typedEvent.MsgIndex).To(Equal(anyMsgIndex))
			//Expect(typedEvent.DenomId).To(Equal(anyDenomId))
			//Expect(typedEvent.TokenId).To(Equal(anyTokenId))
			//Expect(typedEvent.TokenName).To(Equal(anyTokenName))
			//Expect(typedEvent.URI).To(Equal(anyURI))
			//Expect(typedEvent.Data).To(Equal(anyData))
		})
	})
})
