package account_message

import (
	"fmt"

	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/external/tmcosmosutils"
	"github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper"

	"github.com/crypto-com/chain-indexing/projection/account_message/view"

	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	"github.com/crypto-com/chain-indexing/external/utctime"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ projection_entity.Projection = &AccountMessage{}

var (
	NewAccountMessages           = view.NewAccountMessagesView
	NewAccountMessagesTotal      = view.NewAccountMessagesTotalView
	UpdateLastHandledEventHeight = (*AccountMessage).UpdateLastHandledEventHeight
)

type AccountMessage struct {
	*rdbprojectionbase.Base

	rdbConn rdb.Conn
	logger  applogger.Logger

	accountAddressPrefix string

	migrationHelper migrationhelper.MigrationHelper
}

func NewAccountMessage(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	accountAddressPrefix string,
	migrationHelper migrationhelper.MigrationHelper,
) *AccountMessage {
	return &AccountMessage{
		rdbprojectionbase.NewRDbBase(
			rdbConn.ToHandle(),
			"AccountMessage",
		),

		rdbConn,
		logger,

		accountAddressPrefix,
		migrationHelper,
	}
}

func (_ *AccountMessage) GetEventsToListen() []string {
	return append([]string{
		event_usecase.BLOCK_CREATED,
	}, event_usecase.MSG_EVENTS...)
}

func (projection *AccountMessage) OnInit() error {
	if projection.migrationHelper != nil {
		projection.migrationHelper.Migrate()
	}
	return nil
}

func (projection *AccountMessage) HandleEvents(height int64, events []event_entity.Event) error {
	rdbTx, err := projection.rdbConn.Begin()
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

	// TODO: Handle genesis transaction
	if height == int64(0) {
		if err := projection.UpdateLastHandledEventHeight(rdbTxHandle, height); err != nil {
			return fmt.Errorf("error updating last handled event height: %v", err)
		}

		if err := rdbTx.Commit(); err != nil {
			return fmt.Errorf("error committing changes: %v", err)
		}
		committed = true
		return nil
	}

	accountMessagesView := NewAccountMessages(rdbTxHandle)
	accountMessagesTotalView := NewAccountMessagesTotal(rdbTxHandle)

	var blockTime utctime.UTCTime
	var blockHash string
	accountMessages := make([]view.AccountMessageRecord, 0)
	for _, event := range events {
		if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
			blockTime = blockCreatedEvent.Block.Time
			blockHash = blockCreatedEvent.Block.Hash
		}
	}

	for _, event := range events {
		if typedEvent, ok := event.(*event_usecase.MsgSend); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.FromAddress,
					typedEvent.ToAddress,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgMultiSend); ok {
			involvedAccounts := make([]string, 0)
			for _, input := range typedEvent.Inputs {
				involvedAccounts = append(involvedAccounts, input.Address)
			}
			for _, output := range typedEvent.Outputs {
				involvedAccounts = append(involvedAccounts, output.Address)
			}
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: involvedAccounts,
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgSetWithdrawAddress); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.DelegatorAddress,
					typedEvent.WithdrawAddress,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgWithdrawDelegatorReward); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.DelegatorAddress,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgWithdrawValidatorCommission); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.RecipientAddress,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgFundCommunityPool); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Depositor,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgSubmitParamChangeProposal); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.ProposerAddress,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgSubmitCommunityPoolSpendProposal); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.ProposerAddress,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgSubmitSoftwareUpgradeProposal); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.ProposerAddress,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgSubmitCancelSoftwareUpgradeProposal); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.ProposerAddress,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgSubmitProposal); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Proposer,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgDeposit); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Depositor,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgDepositV1); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Depositor,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgVote); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Voter,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgVoteV1); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Voter,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgVoteWeightedV1); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Voter,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgCreateValidator); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.DelegatorAddress,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgEditValidator); ok {
			accountAddress := tmcosmosutils.MustAccountAddressFromValidatorAddress(
				projection.accountAddressPrefix,
				typedEvent.ValidatorAddress,
			)

			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					accountAddress,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgDelegate); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.DelegatorAddress,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgUndelegate); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.DelegatorAddress,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgBeginRedelegate); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.DelegatorAddress,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgUnjail); ok {
			accountAddress := tmcosmosutils.MustAccountAddressFromValidatorAddress(
				projection.accountAddressPrefix,
				typedEvent.ValidatorAddr,
			)

			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					accountAddress,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgNFTIssueDenom); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Sender,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgNFTMintNFT); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Sender,
					typedEvent.Recipient,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgNFTTransferNFT); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Sender,
					typedEvent.Recipient,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgNFTEditNFT); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Sender,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgNFTBurnNFT); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Sender,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCCreateClient); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCUpdateClient); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCConnectionOpenInit); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCConnectionOpenAck); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCConnectionOpenTry); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCConnectionOpenConfirm); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCChannelOpenInit); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCChannelOpenAck); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCChannelOpenTry); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCChannelOpenConfirm); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCAcknowledgement); ok {

			record := view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
			}

			if typedEvent.Params.MaybeFungibleTokenPacketData != nil {
				record.Accounts = append(record.Accounts, typedEvent.Params.MaybeFungibleTokenPacketData.Sender)
			}

			accountMessages = append(accountMessages, record)

		} else if typedEvent, ok := event.(*event_usecase.MsgIBCRecvPacket); ok {

			record := view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
			}

			if typedEvent.Params.MaybeFungibleTokenPacketData != nil {
				record.Accounts = append(record.Accounts, typedEvent.Params.MaybeFungibleTokenPacketData.Receiver)
			}

			accountMessages = append(accountMessages, record)

		} else if typedEvent, ok := event.(*event_usecase.MsgIBCTransferTransfer); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Params.Sender,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCTimeout); ok {

			record := view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
			}

			if typedEvent.Params.MaybeMsgTransfer != nil {
				record.Accounts = append(record.Accounts, typedEvent.Params.MaybeMsgTransfer.RefundReceiver)
			}

			accountMessages = append(accountMessages, record)

		} else if typedEvent, ok := event.(*event_usecase.MsgIBCTimeoutOnClose); ok {

			record := view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
			}

			if typedEvent.Params.MaybeMsgTransfer != nil {
				record.Accounts = append(record.Accounts, typedEvent.Params.MaybeMsgTransfer.RefundReceiver)
			}

			accountMessages = append(accountMessages, record)

		} else if typedEvent, ok := event.(*event_usecase.MsgIBCChannelCloseInit); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgIBCChannelCloseConfirm); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Params.Signer,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgGrant); ok {

			var accounts []string

			if typedEvent.Params.MaybeSendGrant != nil {
				accounts = []string{
					typedEvent.Params.MaybeSendGrant.Granter,
					typedEvent.Params.MaybeSendGrant.Grantee,
				}
			} else if typedEvent.Params.MaybeStakeGrant != nil {
				accounts = []string{
					typedEvent.Params.MaybeStakeGrant.Granter,
					typedEvent.Params.MaybeStakeGrant.Grantee,
				}
			} else if typedEvent.Params.MaybeGenericGrant != nil {
				accounts = []string{
					typedEvent.Params.MaybeGenericGrant.Granter,
					typedEvent.Params.MaybeGenericGrant.Grantee,
				}
			}

			accountMessage := view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: accounts,
			}

			accountMessages = append(accountMessages, accountMessage)

		} else if typedEvent, ok := event.(*event_usecase.MsgRevoke); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Params.Granter,
					typedEvent.Params.Grantee,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgExec); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Params.Grantee,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgGrantAllowance); ok {

			var accounts []string

			if typedEvent.Params.MaybeBasicAllowance != nil {
				accounts = []string{
					typedEvent.Params.MaybeBasicAllowance.Granter,
					typedEvent.Params.MaybeBasicAllowance.Grantee,
				}
			} else if typedEvent.Params.MaybePeriodicAllowance != nil {
				accounts = []string{
					typedEvent.Params.MaybePeriodicAllowance.Granter,
					typedEvent.Params.MaybePeriodicAllowance.Grantee,
				}
			} else if typedEvent.Params.MaybeAllowedMsgAllowance != nil {
				accounts = []string{
					typedEvent.Params.MaybeAllowedMsgAllowance.Granter,
					typedEvent.Params.MaybeAllowedMsgAllowance.Grantee,
				}
			}

			accountMessage := view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: accounts,
			}

			accountMessages = append(accountMessages, accountMessage)

		} else if typedEvent, ok := event.(*event_usecase.MsgRevokeAllowance); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Params.Granter,
					typedEvent.Params.Grantee,
				},
			})
		} else if typedEvent, ok := event.(*event_usecase.MsgCreateVestingAccount); ok {
			accountMessages = append(accountMessages, view.AccountMessageRecord{
				Row: view.AccountMessageRow{
					BlockHeight:     height,
					BlockHash:       "",
					BlockTime:       utctime.UTCTime{},
					TransactionHash: typedEvent.TxHash(),
					Success:         typedEvent.TxSuccess(),
					MessageIndex:    typedEvent.MsgIndex,
					MessageType:     typedEvent.MsgType(),
					Data:            typedEvent,
				},
				Accounts: []string{
					typedEvent.Params.FromAddress,
					typedEvent.Params.ToAddress,
				},
			})
		}
	}

	for i, accountMessage := range accountMessages {
		// TODO: Change to use InsertAll
		accountMessages[i].Row.BlockHash = blockHash
		accountMessages[i].Row.BlockTime = blockTime

		insertedAccounts := make(map[string]bool)
		deduplicatedAccounts := make([]string, 0)
		for _, involvedAccount := range accountMessage.Accounts {
			// Deduplication
			if _, exist := insertedAccounts[involvedAccount]; exist {
				continue
			}

			if err := accountMessagesTotalView.Increment(fmt.Sprintf("%s:-", involvedAccount), 1); err != nil {
				return fmt.Errorf("error incremnting total account message of account: %w", err)
			}
			if err := accountMessagesTotalView.Increment(
				fmt.Sprintf("%s:%s", involvedAccount, accountMessage.Row.MessageType), 1,
			); err != nil {
				return fmt.Errorf("error incremnting total account message of account: %w", err)
			}
			deduplicatedAccounts = append(deduplicatedAccounts, involvedAccount)
			insertedAccounts[involvedAccount] = true
		}

		if err := accountMessagesView.InsertAll(&accountMessages[i].Row, deduplicatedAccounts); err != nil {
			return fmt.Errorf("error inserting account message: %w", err)
		}
	}

	if err := UpdateLastHandledEventHeight(projection, rdbTxHandle, height); err != nil {
		return fmt.Errorf("error updating last handled event height: %v", err)
	}

	if err := rdbTx.Commit(); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}
	committed = true
	return nil
}
