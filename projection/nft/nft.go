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
	return &NFT{
		rdbprojectionbase.NewRDbBase(rdbConn.ToHandle(), "nft"),

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
	rdbTx, err := nft.rdbConn.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction: %v", err)
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
	tokenTransfersView := view.NewTokenTransfers(rdbTxHandle)

	var blockTime utctime.UTCTime
	for _, event := range events {
		if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
			blockTime = blockCreatedEvent.Block.Time
		}
	}

	for _, event := range events {
		if msgIssueDenom, ok := event.(*event_usecase.MsgNFTIssueDenom); ok {
			row := view.DenomRow{
				DenomId:   msgIssueDenom.DenomId,
				Name:      msgIssueDenom.DenomName,
				Schema:    msgIssueDenom.Schema,
				Creator:   msgIssueDenom.Sender,
				CreatedAt: blockTime,
			}

			if insertDenomErr := denomsView.Insert(&row); insertDenomErr != nil {
				return fmt.Errorf("error inserting NFT denom into view: %v", insertDenomErr)
			}

			identifier := msgIssueDenom.DenomId
			if incrementDenomTotalErr := denomsTotalView.Increment(identifier, 1); incrementDenomTotalErr != nil {
				return fmt.Errorf("error incrementing NFT denom to view: %v", incrementDenomTotalErr)
			}

		} else if msgMintNFT, ok := event.(*event_usecase.MsgNFTMintNFT); ok {
			drop := ""
			if nft.config.EnableDrop {
				rawDrop, getDropErr := utils.GetValueFromJSONData(
					[]byte(msgMintNFT.Data), nft.config.DropDataAccessor,
				)
				if getDropErr == nil {
					rawDropStr, isDropOk := rawDrop.(string)
					if isDropOk {
						drop = rawDropStr
					}
				}
			}
			tokenRow := view.TokenRow{
				TokenId:      msgMintNFT.TokenId,
				DenomId:      msgMintNFT.DenomId,
				Drop:         drop,
				Burned:       false,
				Name:         msgMintNFT.TokenName,
				URI:          msgMintNFT.URI,
				Data:         msgMintNFT.Data,
				Minter:       msgMintNFT.Sender,
				Owner:        msgMintNFT.Recipient,
				MintedAt:     blockTime,
				LastEditedAt: blockTime,
			}
			if insertTokenErr := nft.insertToken(tokensView, tokenRow, tokensTotalView); insertTokenErr != nil {
				return insertTokenErr
			}

			tokenTransferRow := view.TokenTransferRow{
				DenomId:         msgMintNFT.DenomId,
				TokenId:         msgMintNFT.TokenId,
				BlockHeight:     height,
				TransactionHash: msgMintNFT.TxHash(),
				Sender:          msgMintNFT.Sender,
				Recipient:       msgMintNFT.Recipient,
				TransferredAt:   blockTime,
			}
			if insertTokenTransferErr := tokenTransfersView.Insert(
				tokenTransferRow,
			); insertTokenTransferErr != nil {
				return fmt.Errorf(
					"error inserting NFT token transfer into view: %v",
					insertTokenTransferErr,
				)
			}

		} else if msgEditNFT, ok := event.(*event_usecase.MsgNFTEditNFT); ok {
			prevTokenRow, queryPrevTokenRow := tokensView.FindById(msgEditNFT.DenomId, msgEditNFT.TokenId)
			if queryPrevTokenRow != nil {
				return fmt.Errorf("error querying NFT token being edited: %v", queryPrevTokenRow)
			}

			drop := prevTokenRow.Drop
			if nft.config.EnableDrop {
				rawDrop, getDropErr := utils.GetValueFromJSONData(
					[]byte(msgEditNFT.Data), nft.config.DropDataAccessor,
				)
				if getDropErr == nil {
					rawDropStr, isDropOk := rawDrop.(string)
					if isDropOk {
						drop = rawDropStr
					}
				}
			}
			tokenRow := view.TokenRow{
				DenomId:      msgEditNFT.DenomId,
				TokenId:      msgEditNFT.TokenId,
				Drop:         drop,
				Burned:       prevTokenRow.Burned,
				Name:         msgEditNFT.TokenName,
				URI:          msgEditNFT.URI,
				Data:         msgEditNFT.Data,
				Minter:       prevTokenRow.Minter,
				Owner:        prevTokenRow.Owner,
				MintedAt:     prevTokenRow.MintedAt,
				LastEditedAt: blockTime,
			}

			if updateTokenErr := nft.updateToken(
				tokensView,
				prevTokenRow.TokenRow,
				tokenRow,
				tokensTotalView,
			); updateTokenErr != nil {
				return updateTokenErr
			}

		} else if msgBurnNFT, ok := event.(*event_usecase.MsgNFTBurnNFT); ok {
			prevTokenRow, queryPrevTokenRowErr := tokensView.FindById(msgBurnNFT.DenomId, msgBurnNFT.TokenId)
			if queryPrevTokenRowErr != nil {
				return fmt.Errorf("error querying NFT token being edited: %v", queryPrevTokenRowErr)
			}

			tokenRow := view.TokenRow{
				DenomId:      msgBurnNFT.DenomId,
				TokenId:      msgBurnNFT.TokenId,
				Drop:         prevTokenRow.Drop,
				Burned:       true,
				Name:         prevTokenRow.Name,
				URI:          prevTokenRow.URI,
				Data:         prevTokenRow.Data,
				Minter:       prevTokenRow.Minter,
				Owner:        prevTokenRow.Owner,
				MintedAt:     prevTokenRow.MintedAt,
				LastEditedAt: blockTime,
			}

			if updateTokenErr := tokensView.Update(tokenRow); updateTokenErr != nil {
				return fmt.Errorf("error updating NFT token row: %v", updateTokenErr)
			}

			if decrementErr := tokensTotalView.DecrementAll([]string{
				fmt.Sprintf("-:-:-:%s", tokenRow.Owner),
				fmt.Sprintf("%s:-:-:%s", tokenRow.DenomId, tokenRow.Owner),
				fmt.Sprintf("-:%s:-:%s", tokenRow.Drop, tokenRow.Owner),
				fmt.Sprintf("-:-:%s:%s", tokenRow.Minter, tokenRow.Owner),
				fmt.Sprintf("%s:%s:-:%s", tokenRow.DenomId, tokenRow.Drop, tokenRow.Owner),
				fmt.Sprintf("%s:-:%s:%s", tokenRow.DenomId, tokenRow.Minter, tokenRow.Owner),
				fmt.Sprintf("-:%s:%s:%s", tokenRow.Drop, tokenRow.Minter, tokenRow.Owner),
				fmt.Sprintf("%s:%s:%s:%s", tokenRow.DenomId, tokenRow.Drop, tokenRow.Minter, tokenRow.Owner),
			}, 1); decrementErr != nil {
				return fmt.Errorf("error decrementing NFT token total: %v", decrementErr)
			}

		} else if msgTransferNFT, ok := event.(*event_usecase.MsgNFTTransferNFT); ok {
			prevTokenRow, queryPrevTokenRowErr := tokensView.FindById(msgTransferNFT.DenomId, msgTransferNFT.TokenId)
			if queryPrevTokenRowErr != nil {
				return fmt.Errorf("error querying NFT token being edited: %v", queryPrevTokenRowErr)
			}
			tokenRow := view.TokenRow{
				DenomId:      msgTransferNFT.DenomId,
				TokenId:      msgTransferNFT.TokenId,
				Drop:         prevTokenRow.Drop,
				Burned:       prevTokenRow.Burned,
				Name:         prevTokenRow.Name,
				URI:          prevTokenRow.URI,
				Data:         prevTokenRow.Data,
				Minter:       prevTokenRow.Minter,
				Owner:        msgTransferNFT.Recipient,
				MintedAt:     prevTokenRow.MintedAt,
				LastEditedAt: blockTime,
			}
			if updateTokenErr := nft.updateToken(
				tokensView,
				prevTokenRow.TokenRow,
				tokenRow,
				tokensTotalView,
			); updateTokenErr != nil {
				return updateTokenErr
			}

			transferRow := view.TokenTransferRow{
				DenomId:         msgTransferNFT.DenomId,
				TokenId:         msgTransferNFT.TokenId,
				BlockHeight:     height,
				TransactionHash: msgTransferNFT.TxHash(),
				Sender:          msgTransferNFT.Sender,
				Recipient:       msgTransferNFT.Recipient,
				TransferredAt:   blockTime,
			}

			if insertTransferErr := tokenTransfersView.Insert(transferRow); insertTransferErr != nil {
				return fmt.Errorf("error inserting NFT token transfer into view: %v", insertTransferErr)
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
		fmt.Sprintf("-:%s:-:%s", prevTokenRow.Drop, prevTokenRow.Owner),
		fmt.Sprintf("-:-:%s:%s", prevTokenRow.Minter, prevTokenRow.Owner),
		fmt.Sprintf("%s:%s:-:%s", prevTokenRow.DenomId, prevTokenRow.Drop, prevTokenRow.Owner),
		fmt.Sprintf("%s:-:%s:%s", prevTokenRow.DenomId, prevTokenRow.Minter, prevTokenRow.Owner),
		fmt.Sprintf("-:%s:%s:%s", prevTokenRow.Drop, prevTokenRow.Minter, prevTokenRow.Owner),
		fmt.Sprintf("%s:%s:%s:%s", prevTokenRow.DenomId, prevTokenRow.Drop, prevTokenRow.Minter, prevTokenRow.Owner),
	}, 1); decrementErr != nil {
		return fmt.Errorf("error decrementing NFT token total: %v", decrementErr)
	}

	// Increment all total entries with new owner
	if incrementErr := tokensTotalView.IncrementAll([]string{
		fmt.Sprintf("-:-:-:%s", tokenRow.Owner),
		fmt.Sprintf("%s:-:-:%s", tokenRow.DenomId, tokenRow.Owner),
		fmt.Sprintf("-:%s:-:%s", tokenRow.Drop, tokenRow.Owner),
		fmt.Sprintf("-:-:%s:%s", tokenRow.Minter, tokenRow.Owner),
		fmt.Sprintf("%s:%s:-:%s", tokenRow.DenomId, tokenRow.Drop, tokenRow.Owner),
		fmt.Sprintf("%s:-:%s:%s", tokenRow.DenomId, tokenRow.Minter, tokenRow.Owner),
		fmt.Sprintf("-:%s:%s:%s", tokenRow.Drop, tokenRow.Minter, tokenRow.Owner),
		fmt.Sprintf("%s:%s:%s:%s", tokenRow.DenomId, tokenRow.Drop, tokenRow.Minter, tokenRow.Owner),
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
		fmt.Sprintf("-:%s:-:-", tokenRow.Drop),
		fmt.Sprintf("-:-:%s:-", tokenRow.Minter),
		fmt.Sprintf("-:-:-:%s", tokenRow.Owner),
		fmt.Sprintf("%s:%s:-:-", tokenRow.DenomId, tokenRow.Drop),
		fmt.Sprintf("%s:-:%s:-", tokenRow.DenomId, tokenRow.Minter),
		fmt.Sprintf("%s:-:-:%s", tokenRow.DenomId, tokenRow.Owner),
		fmt.Sprintf("-:%s:%s:-", tokenRow.Drop, tokenRow.Minter),
		fmt.Sprintf("-:%s:-:%s", tokenRow.Drop, tokenRow.Owner),
		fmt.Sprintf("-:-:%s:%s", tokenRow.Minter, tokenRow.Owner),
		fmt.Sprintf("%s:%s:%s:-", tokenRow.DenomId, tokenRow.Drop, tokenRow.Minter),
		fmt.Sprintf("%s:%s:-:%s", tokenRow.DenomId, tokenRow.Drop, tokenRow.Owner),
		fmt.Sprintf("%s:-:%s:%s", tokenRow.DenomId, tokenRow.Minter, tokenRow.Owner),
		fmt.Sprintf("-:%s:%s:%s", tokenRow.Drop, tokenRow.Minter, tokenRow.Owner),
		fmt.Sprintf("%s:%s:%s:%s", tokenRow.DenomId, tokenRow.Drop, tokenRow.Minter, tokenRow.Owner),
	}, 1); incrementErr != nil {
		return fmt.Errorf("error incrementing NFT token total: %v", incrementErr)
	}
	return nil
}
