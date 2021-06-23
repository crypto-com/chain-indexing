package nft

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/projection/nft/utils"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	"github.com/crypto-com/chain-indexing/internal/utctime"
	"github.com/crypto-com/chain-indexing/projection/nft/view"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ projection_entity.Projection = &NFT{}

const DO_NOT_MODIFY = "[do-not-modify]"

type NFT struct {
	*rdbprojectionbase.Base

	rdbConn rdb.Conn
	logger  applogger.Logger

	config Config
}

type Config struct {
	EnableDrop       bool
	DropDataAccessor string
}

func NewNFT(logger applogger.Logger, rdbConn rdb.Conn, config Config) *NFT {
	projectionId := "NFT"
	if config.EnableDrop {
		projectionId = "CryptoComNFT"
	}
	return &NFT{
		rdbprojectionbase.NewRDbBase(rdbConn.ToHandle(), projectionId),

		rdbConn,
		logger,
		config,
	}
}

func (nft *NFT) GetEventsToListen() []string {
	return []string{
		event_usecase.BLOCK_CREATED,
		event_usecase.MSG_NFT_ISSUE_DENOM_CREATED,
		event_usecase.MSG_NFT_MINT_NFT_CREATED,
		event_usecase.MSG_NFT_TRANSFER_NFT_CREATED,
		event_usecase.MSG_NFT_EDIT_NFT_CREATED,
		event_usecase.MSG_NFT_BURN_NFT_CREATED,
	}
}

func (nft *NFT) OnInit() error {
	return nil
}

func (nft *NFT) HandleEvents(height int64, events []event_entity.Event) error {
	rdbTx, rdbTxErr := nft.rdbConn.Begin()
	if rdbTxErr != nil {
		return fmt.Errorf("error beginning transaction: %v", rdbTxErr)
	}

	committed := false
	defer func() {
		if !committed {
			_ = rdbTx.Rollback()
		}
	}()

	rdbTxHandle := rdbTx.ToHandle()

	denomsView := view.NewDenoms(rdbTxHandle)
	denomsTotalView := view.NewDenomsTotal(rdbTxHandle)
	tokensView := view.NewTokens(rdbTxHandle)
	tokensTotalView := view.NewTokensTotal(rdbTxHandle)
	nftMessagesView := view.NewMessages(rdbTxHandle)
	nftMessagesTotalView := view.NewMessagesTotal(rdbTxHandle)

	var blockTime utctime.UTCTime
	var blockHash string
	for _, event := range events {
		if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
			blockTime = blockCreatedEvent.Block.Time
			blockHash = blockCreatedEvent.Block.Hash
		}
	}

	for _, event := range events {
		if msgIssueDenom, ok := event.(*event_usecase.MsgNFTIssueDenom); ok {
			row := view.DenomRow{
				DenomId:              msgIssueDenom.DenomId,
				Name:                 msgIssueDenom.DenomName,
				Schema:               msgIssueDenom.Schema,
				Creator:              msgIssueDenom.Sender,
				CreatedAt:            blockTime,
				CreatedAtBlockHeight: height,
			}

			if insertDenomErr := denomsView.Insert(&row); insertDenomErr != nil {
				return fmt.Errorf("error inserting NFT denom into view: %v", insertDenomErr)
			}

			identifier := msgIssueDenom.DenomId
			if incrementDenomTotalErr := denomsTotalView.Increment("-", 1); incrementDenomTotalErr != nil {
				return fmt.Errorf("error incrementing total NFT denom to view: %v", incrementDenomTotalErr)
			}
			if incrementDenomTotalErr := denomsTotalView.Increment(identifier, 1); incrementDenomTotalErr != nil {
				return fmt.Errorf("error incrementing NFT denom to view: %v", incrementDenomTotalErr)
			}

			if insertMessageErr := nft.insertMessage(nftMessagesView, nftMessagesTotalView, view.MessageRow{
				BlockHeight:     height,
				BlockHash:       blockHash,
				BlockTime:       blockTime,
				DenomId:         msgIssueDenom.DenomId,
				MaybeTokenId:    nil,
				MaybeDrop:       nil,
				TransactionHash: msgIssueDenom.TxHash(),
				Success:         msgIssueDenom.TxSuccess(),
				MessageIndex:    msgIssueDenom.MsgIndex,
				MessageType:     msgIssueDenom.MsgType(),
				Data:            msgIssueDenom,
			}); insertMessageErr != nil {
				return insertMessageErr
			}

		} else if msgMintNFT, ok := event.(*event_usecase.MsgNFTMintNFT); ok {
			var maybeDrop *string
			if nft.config.EnableDrop {
				rawDrop, getDropErr := utils.GetValueFromJSONData(
					[]byte(msgMintNFT.Data), nft.config.DropDataAccessor,
				)
				if getDropErr == nil {
					rawDropStr, isDropOk := rawDrop.(string)
					if isDropOk {
						maybeDrop = &rawDropStr
					}
				}
			}
			tokenRow := view.TokenRow{
				TokenId:                 msgMintNFT.TokenId,
				DenomId:                 msgMintNFT.DenomId,
				MaybeDrop:               maybeDrop,
				Name:                    msgMintNFT.TokenName,
				URI:                     msgMintNFT.URI,
				Data:                    msgMintNFT.Data,
				Minter:                  msgMintNFT.Sender,
				Owner:                   msgMintNFT.Recipient,
				MintedAt:                blockTime,
				MintedAtBlockHeight:     height,
				LastEditedAt:            blockTime,
				LastEditedAtBlockHeight: height,
			}
			if insertTokenErr := nft.insertToken(tokensView, tokenRow, tokensTotalView); insertTokenErr != nil {
				return insertTokenErr
			}

			if insertMessageErr := nft.insertMessage(nftMessagesView, nftMessagesTotalView, view.MessageRow{
				BlockHeight:     height,
				BlockHash:       blockHash,
				BlockTime:       blockTime,
				DenomId:         msgMintNFT.DenomId,
				MaybeTokenId:    &msgMintNFT.TokenId,
				MaybeDrop:       maybeDrop,
				TransactionHash: msgMintNFT.TxHash(),
				Success:         msgMintNFT.TxSuccess(),
				MessageIndex:    msgMintNFT.MsgIndex,
				MessageType:     msgMintNFT.MsgType(),
				Data:            msgMintNFT,
			}); insertMessageErr != nil {
				return insertMessageErr
			}

		} else if msgEditNFT, ok := event.(*event_usecase.MsgNFTEditNFT); ok {
			mutPrevTokenRow, queryPrevTokenRow := tokensView.FindById(msgEditNFT.DenomId, msgEditNFT.TokenId)
			if queryPrevTokenRow != nil {
				return fmt.Errorf("error querying NFT token being edited: %v", queryPrevTokenRow)
			}

			drop := mutPrevTokenRow.MaybeDrop
			if nft.config.EnableDrop {
				rawDrop, getDropErr := utils.GetValueFromJSONData(
					[]byte(msgEditNFT.Data), nft.config.DropDataAccessor,
				)
				if getDropErr == nil {
					rawDropStr, isDropOk := rawDrop.(string)
					if isDropOk {
						drop = &rawDropStr
					}
				}
			}

			if msgEditNFT.TokenName != DO_NOT_MODIFY {
				mutPrevTokenRow.Name = msgEditNFT.TokenName
			}
			if msgEditNFT.URI != DO_NOT_MODIFY {
				mutPrevTokenRow.URI = msgEditNFT.URI
			}
			if msgEditNFT.Data != DO_NOT_MODIFY {
				mutPrevTokenRow.Data = msgEditNFT.Data
			}

			tokenRow := view.TokenRow{
				DenomId:                 msgEditNFT.DenomId,
				TokenId:                 msgEditNFT.TokenId,
				MaybeDrop:               drop,
				Name:                    mutPrevTokenRow.Name,
				URI:                     mutPrevTokenRow.URI,
				Data:                    mutPrevTokenRow.Data,
				Minter:                  mutPrevTokenRow.Minter,
				Owner:                   mutPrevTokenRow.Owner,
				MintedAt:                mutPrevTokenRow.MintedAt,
				MintedAtBlockHeight:     mutPrevTokenRow.MintedAtBlockHeight,
				LastEditedAt:            blockTime,
				LastEditedAtBlockHeight: height,
			}

			if updateTokenErr := nft.updateToken(
				tokensView,
				mutPrevTokenRow.TokenRow,
				tokenRow,
				tokensTotalView,
			); updateTokenErr != nil {
				return updateTokenErr
			}

			if insertMessageErr := nft.insertMessage(nftMessagesView, nftMessagesTotalView, view.MessageRow{
				BlockHeight:     height,
				BlockHash:       blockHash,
				BlockTime:       blockTime,
				DenomId:         msgEditNFT.DenomId,
				MaybeTokenId:    &msgEditNFT.TokenId,
				MaybeDrop:       drop,
				TransactionHash: msgEditNFT.TxHash(),
				Success:         msgEditNFT.TxSuccess(),
				MessageIndex:    msgEditNFT.MsgIndex,
				MessageType:     msgEditNFT.MsgType(),
				Data:            msgEditNFT,
			}); insertMessageErr != nil {
				return insertMessageErr
			}

		} else if msgBurnNFT, ok := event.(*event_usecase.MsgNFTBurnNFT); ok {
			prevTokenRow, queryPrevTokenRowErr := tokensView.FindById(msgBurnNFT.DenomId, msgBurnNFT.TokenId)
			if queryPrevTokenRowErr != nil {
				return fmt.Errorf("error querying NFT token being edited: %v", queryPrevTokenRowErr)
			}

			if deleteTokenErr := nft.deleteToken(
				tokensView, tokensTotalView, nftMessagesView, prevTokenRow.TokenRow,
			); deleteTokenErr != nil {
				return fmt.Errorf("error deleteing burnt NFT token: %v", deleteTokenErr)
			}

			// Burn does not need to record the message because the token and all the records will
			// be deleted

		} else if msgTransferNFT, ok := event.(*event_usecase.MsgNFTTransferNFT); ok {
			prevTokenRow, queryPrevTokenRowErr := tokensView.FindById(msgTransferNFT.DenomId, msgTransferNFT.TokenId)
			if queryPrevTokenRowErr != nil {
				return fmt.Errorf("error querying NFT token being edited: %v", queryPrevTokenRowErr)
			}
			tokenRow := view.TokenRow{
				DenomId:                 msgTransferNFT.DenomId,
				TokenId:                 msgTransferNFT.TokenId,
				MaybeDrop:               prevTokenRow.MaybeDrop,
				Name:                    prevTokenRow.Name,
				URI:                     prevTokenRow.URI,
				Data:                    prevTokenRow.Data,
				Minter:                  prevTokenRow.Minter,
				Owner:                   msgTransferNFT.Recipient,
				MintedAt:                prevTokenRow.MintedAt,
				MintedAtBlockHeight:     prevTokenRow.MintedAtBlockHeight,
				LastEditedAt:            blockTime,
				LastEditedAtBlockHeight: height,
			}
			if updateTokenErr := nft.updateToken(
				tokensView,
				prevTokenRow.TokenRow,
				tokenRow,
				tokensTotalView,
			); updateTokenErr != nil {
				return updateTokenErr
			}

			if insertMessageErr := nft.insertMessage(nftMessagesView, nftMessagesTotalView, view.MessageRow{
				BlockHeight:     height,
				BlockHash:       blockHash,
				BlockTime:       blockTime,
				DenomId:         msgTransferNFT.DenomId,
				MaybeTokenId:    &msgTransferNFT.TokenId,
				MaybeDrop:       prevTokenRow.MaybeDrop,
				TransactionHash: msgTransferNFT.TxHash(),
				Success:         msgTransferNFT.TxSuccess(),
				MessageIndex:    msgTransferNFT.MsgIndex,
				MessageType:     msgTransferNFT.MsgType(),
				Data:            msgTransferNFT,
			}); insertMessageErr != nil {
				return insertMessageErr
			}
		}
	}

	if err := nft.UpdateLastHandledEventHeight(rdbTxHandle, height); err != nil {
		return fmt.Errorf("error updating last handled event height: %v", err)
	}

	if err := rdbTx.Commit(); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}
	committed = true
	return nil
}

func (nft *NFT) insertMessage(
	messagesView *view.Messages,
	messagesTotalView *view.MessagesTotal,
	message view.MessageRow,
) error {
	totalIdentities := []string{
		"-:-:-:-",
		fmt.Sprintf("%s:-:-:-", nilIdentifier(message.MaybeDrop)),
		fmt.Sprintf("-:%s:-:-", message.DenomId),
		fmt.Sprintf("-:-:%s:-", nilIdentifier(message.MaybeTokenId)),
		fmt.Sprintf("-:-:-:%s", message.MessageType),
		fmt.Sprintf("%s:%s:-:-", nilIdentifier(message.MaybeDrop), message.DenomId),
		fmt.Sprintf("%s:-:%s:-", nilIdentifier(message.MaybeDrop), nilIdentifier(message.MaybeTokenId)),
		fmt.Sprintf("%s:-:-:%s", nilIdentifier(message.MaybeDrop), message.MessageType),
		fmt.Sprintf("-:%s:%s:-", message.DenomId, nilIdentifier(message.MaybeTokenId)),
		fmt.Sprintf("-:%s:-:%s", message.DenomId, message.MessageType),
		fmt.Sprintf("-:-:%s:%s", nilIdentifier(message.MaybeTokenId), message.MessageType),
		fmt.Sprintf("%s:%s:%s:-", nilIdentifier(message.MaybeDrop), message.DenomId, nilIdentifier(message.MaybeTokenId)),
		fmt.Sprintf("%s:%s:-:%s", nilIdentifier(message.MaybeDrop), message.DenomId, message.MessageType),
		fmt.Sprintf("%s:-:%s:%s", nilIdentifier(message.MaybeDrop), nilIdentifier(message.MaybeTokenId), message.MessageType),
		fmt.Sprintf("-:%s:%s:%s", message.DenomId, nilIdentifier(message.MaybeTokenId), message.MessageType),
		fmt.Sprintf("%s:%s:%s:%s", nilIdentifier(message.MaybeDrop), message.DenomId, nilIdentifier(message.MaybeTokenId), message.MessageType),
	}
	if err := messagesTotalView.IncrementAll(totalIdentities, 1); err != nil {
		return fmt.Errorf("error incremnting NFT message total: %w", err)
	}

	if err := messagesView.Insert(&message); err != nil {
		return fmt.Errorf("error inserting NFT message: %w", err)
	}

	return nil
}

func (nft *NFT) updateToken(
	tokensView *view.Tokens,
	prevTokenRow view.TokenRow,
	tokenRow view.TokenRow,
	tokensTotalView *view.TokensTotal,
) error {
	if updateTokenErr := tokensView.Update(tokenRow); updateTokenErr != nil {
		return fmt.Errorf("error updating NFT token row: %v", updateTokenErr)
	}

	// Decrement all total entries with previous owner
	if decrementErr := tokensTotalView.DecrementAll([]string{
		fmt.Sprintf("-:-:-:%s", prevTokenRow.Owner),
		fmt.Sprintf("%s:-:-:%s", prevTokenRow.DenomId, prevTokenRow.Owner),
		fmt.Sprintf("-:%s:-:%s", nilIdentifier(prevTokenRow.MaybeDrop), prevTokenRow.Owner),
		fmt.Sprintf("-:-:%s:%s", prevTokenRow.Minter, prevTokenRow.Owner),
		fmt.Sprintf("%s:%s:-:%s", prevTokenRow.DenomId, nilIdentifier(prevTokenRow.MaybeDrop), prevTokenRow.Owner),
		fmt.Sprintf("%s:-:%s:%s", prevTokenRow.DenomId, prevTokenRow.Minter, prevTokenRow.Owner),
		fmt.Sprintf("-:%s:%s:%s", nilIdentifier(prevTokenRow.MaybeDrop), prevTokenRow.Minter, prevTokenRow.Owner),
		fmt.Sprintf("%s:%s:%s:%s", prevTokenRow.DenomId, nilIdentifier(prevTokenRow.MaybeDrop), prevTokenRow.Minter, prevTokenRow.Owner),
	}, 1); decrementErr != nil {
		return fmt.Errorf("error decrementing NFT token total: %v", decrementErr)
	}

	// Increment all total entries with new owner
	if incrementErr := tokensTotalView.IncrementAll([]string{
		fmt.Sprintf("-:-:-:%s", tokenRow.Owner),
		fmt.Sprintf("%s:-:-:%s", tokenRow.DenomId, tokenRow.Owner),
		fmt.Sprintf("-:%s:-:%s", nilIdentifier(tokenRow.MaybeDrop), tokenRow.Owner),
		fmt.Sprintf("-:-:%s:%s", tokenRow.Minter, tokenRow.Owner),
		fmt.Sprintf("%s:%s:-:%s", tokenRow.DenomId, nilIdentifier(tokenRow.MaybeDrop), tokenRow.Owner),
		fmt.Sprintf("%s:-:%s:%s", tokenRow.DenomId, tokenRow.Minter, tokenRow.Owner),
		fmt.Sprintf("-:%s:%s:%s", nilIdentifier(tokenRow.MaybeDrop), tokenRow.Minter, tokenRow.Owner),
		fmt.Sprintf("%s:%s:%s:%s", tokenRow.DenomId, nilIdentifier(tokenRow.MaybeDrop), tokenRow.Minter, tokenRow.Owner),
	}, 1); incrementErr != nil {
		return fmt.Errorf("error incrementing NFT token total: %v", incrementErr)
	}
	return nil
}

func (nft *NFT) insertToken(
	tokensView *view.Tokens, tokenRow view.TokenRow, tokensTotalView *view.TokensTotal,
) error {
	if insertTokenErr := tokensView.Insert(&tokenRow); insertTokenErr != nil {
		return fmt.Errorf("error inserting NFT token into view: %v", insertTokenErr)
	}
	if incrementErr := tokensTotalView.IncrementAll([]string{
		"-:-:-:-",
		fmt.Sprintf("%s:-:-:-", tokenRow.DenomId),
		fmt.Sprintf("-:%s:-:-", nilIdentifier(tokenRow.MaybeDrop)),
		fmt.Sprintf("-:-:%s:-", tokenRow.Minter),
		fmt.Sprintf("-:-:-:%s", tokenRow.Owner),
		fmt.Sprintf("%s:%s:-:-", tokenRow.DenomId, nilIdentifier(tokenRow.MaybeDrop)),
		fmt.Sprintf("%s:-:%s:-", tokenRow.DenomId, tokenRow.Minter),
		fmt.Sprintf("%s:-:-:%s", tokenRow.DenomId, tokenRow.Owner),
		fmt.Sprintf("-:%s:%s:-", nilIdentifier(tokenRow.MaybeDrop), tokenRow.Minter),
		fmt.Sprintf("-:%s:-:%s", nilIdentifier(tokenRow.MaybeDrop), tokenRow.Owner),
		fmt.Sprintf("-:-:%s:%s", tokenRow.Minter, tokenRow.Owner),
		fmt.Sprintf("%s:%s:%s:-", tokenRow.DenomId, nilIdentifier(tokenRow.MaybeDrop), tokenRow.Minter),
		fmt.Sprintf("%s:%s:-:%s", tokenRow.DenomId, nilIdentifier(tokenRow.MaybeDrop), tokenRow.Owner),
		fmt.Sprintf("%s:-:%s:%s", tokenRow.DenomId, tokenRow.Minter, tokenRow.Owner),
		fmt.Sprintf("-:%s:%s:%s", nilIdentifier(tokenRow.MaybeDrop), tokenRow.Minter, tokenRow.Owner),
		fmt.Sprintf("%s:%s:%s:%s", tokenRow.DenomId, nilIdentifier(tokenRow.MaybeDrop), tokenRow.Minter, tokenRow.Owner),
	}, 1); incrementErr != nil {
		return fmt.Errorf("error incrementing NFT token total: %v", incrementErr)
	}
	return nil
}

func (nft *NFT) deleteToken(
	tokensView *view.Tokens,
	tokensTotalView *view.TokensTotal,
	nftMessagesView *view.Messages,
	tokenRow view.TokenRow,
) error {
	deletedRowCount, deleteTokenErr := tokensView.Delete(tokenRow.DenomId, tokenRow.TokenId)
	if deleteTokenErr != nil {
		return fmt.Errorf("error deleting NFT token from view: %v", deleteTokenErr)
	}
	if deletedRowCount != 1 {
		return fmt.Errorf("error deleting NFT token from view: %d rows deleted", deletedRowCount)
	}

	deleteMessagesCount, deleteMessagesErr := nftMessagesView.DeleteAllByDenomTokenIds(
		tokenRow.DenomId, tokenRow.TokenId,
	)
	if deleteMessagesErr != nil {
		return fmt.Errorf("error deleting NFT messages from view: %v", deleteMessagesErr)
	}
	if deleteMessagesCount == 0 {
		return fmt.Errorf("error deleting NFT messages from view: no rows deleted")
	}

	if decrementErr := tokensTotalView.DecrementAll([]string{
		"-:-:-:-",
		fmt.Sprintf("%s:-:-:-", tokenRow.DenomId),
		fmt.Sprintf("-:%s:-:-", nilIdentifier(tokenRow.MaybeDrop)),
		fmt.Sprintf("-:-:%s:-", tokenRow.Minter),
		fmt.Sprintf("-:-:-:%s", tokenRow.Owner),
		fmt.Sprintf("%s:%s:-:-", tokenRow.DenomId, nilIdentifier(tokenRow.MaybeDrop)),
		fmt.Sprintf("%s:-:%s:-", tokenRow.DenomId, tokenRow.Minter),
		fmt.Sprintf("%s:-:-:%s", tokenRow.DenomId, tokenRow.Owner),
		fmt.Sprintf("-:%s:%s:-", nilIdentifier(tokenRow.MaybeDrop), tokenRow.Minter),
		fmt.Sprintf("-:%s:-:%s", nilIdentifier(tokenRow.MaybeDrop), tokenRow.Owner),
		fmt.Sprintf("-:-:%s:%s", tokenRow.Minter, tokenRow.Owner),
		fmt.Sprintf("%s:%s:%s:-", tokenRow.DenomId, nilIdentifier(tokenRow.MaybeDrop), tokenRow.Minter),
		fmt.Sprintf("%s:%s:-:%s", tokenRow.DenomId, nilIdentifier(tokenRow.MaybeDrop), tokenRow.Owner),
		fmt.Sprintf("%s:-:%s:%s", tokenRow.DenomId, tokenRow.Minter, tokenRow.Owner),
		fmt.Sprintf("-:%s:%s:%s", nilIdentifier(tokenRow.MaybeDrop), tokenRow.Minter, tokenRow.Owner),
		fmt.Sprintf("%s:%s:%s:%s", tokenRow.DenomId, nilIdentifier(tokenRow.MaybeDrop), tokenRow.Minter, tokenRow.Owner),
	}, 1); decrementErr != nil {
		return fmt.Errorf("error decrementing NFT token total: %v", decrementErr)
	}
	return nil
}

func nilIdentifier(maybeValue *string) string {
	if maybeValue == nil {
		return ""
	}
	return *maybeValue
}
