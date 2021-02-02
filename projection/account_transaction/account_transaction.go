package account_transaction

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/internal/base64"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	projection_entity "github.com/crypto-com/chain-indexing/entity/projection"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	"github.com/crypto-com/chain-indexing/internal/tmcosmosutils"
	"github.com/crypto-com/chain-indexing/internal/utctime"
	"github.com/crypto-com/chain-indexing/projection/account_transaction/view"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

var _ projection_entity.Projection = &AccountTransaction{}

type AccountTransaction struct {
	*rdbprojectionbase.Base

	accountAddressPrefix string

	rdbConn rdb.Conn
	logger  applogger.Logger
}

func NewAccountTransaction(logger applogger.Logger, rdbConn rdb.Conn, accountAddressPrefix string) *AccountTransaction {
	return &AccountTransaction{
		rdbprojectionbase.NewRDbBase(rdbConn.ToHandle(), "AccountTransaction"),

		accountAddressPrefix,

		rdbConn,
		logger,
	}
}

func (_ *AccountTransaction) GetEventsToListen() []string {
	return append([]string{
		event_usecase.BLOCK_CREATED,
		event_usecase.TRANSACTION_CREATED,
		event_usecase.TRANSACTION_FAILED,
	}, event_usecase.MSG_EVENTS...)
}

func (projection *AccountTransaction) OnInit() error {
	return nil
}

func (projection *AccountTransaction) HandleEvents(height int64, events []event_entity.Event) error {
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

	accountTransactionsView := view.NewAccountTransactions(rdbTxHandle)
	accountTransactionDataView := view.NewAccountTransactionData(rdbTxHandle)
	accountTransactionsTotalView := view.NewAccountTransactionsTotal(rdbTxHandle)

	var blockTime utctime.UTCTime
	var blockHash string

	transactionInfos := make(map[string]*TransactionInfo)

	// Handle and insert a single copy of transaction data
	txs := make([]view.TransactionRow, 0)
	txMsgs := make(map[string][]event_usecase.MsgEvent)
	for _, event := range events {
		if blockCreatedEvent, ok := event.(*event_usecase.BlockCreated); ok {
			blockTime = blockCreatedEvent.Block.Time
			blockHash = blockCreatedEvent.Block.Hash
		} else if transactionCreatedEvent, ok := event.(*event_usecase.TransactionCreated); ok {
			txs = append(txs, view.TransactionRow{
				BlockHeight:   height,
				BlockTime:     utctime.UTCTime{}, // placeholder
				Hash:          transactionCreatedEvent.TxHash,
				Index:         transactionCreatedEvent.Index,
				Success:       true,
				Code:          transactionCreatedEvent.Code,
				Log:           transactionCreatedEvent.Log,
				Fee:           transactionCreatedEvent.Fee,
				FeePayer:      transactionCreatedEvent.FeePayer,
				FeeGranter:    transactionCreatedEvent.FeeGranter,
				GasWanted:     transactionCreatedEvent.GasWanted,
				GasUsed:       transactionCreatedEvent.GasUsed,
				Memo:          transactionCreatedEvent.Memo,
				TimeoutHeight: transactionCreatedEvent.TimeoutHeight,
				Messages:      make([]view.TransactionRowMessage, 0),
			})

			transactionInfos[transactionCreatedEvent.TxHash] = NewTransactionInfo(
				view.AccountTransactionBaseRow{
					Account:      "", // placeholder
					BlockHeight:  height,
					BlockHash:    "",                // placeholder
					BlockTime:    utctime.UTCTime{}, // placeholder
					Hash:         transactionCreatedEvent.TxHash,
					MessageTypes: []string{},
					Success:      true,
				},
			)
			senders := projection.ParseSenderAddresses(transactionCreatedEvent.Senders)
			for _, sender := range senders {
				transactionInfos[transactionCreatedEvent.TxHash].AddAccount(sender)
			}
		} else if transactionFailedEvent, ok := event.(*event_usecase.TransactionFailed); ok {
			row := view.TransactionRow{
				BlockHeight:   height,
				BlockTime:     utctime.UTCTime{}, // placeholder
				Hash:          transactionFailedEvent.TxHash,
				Index:         transactionFailedEvent.Index,
				Success:       false,
				Code:          transactionFailedEvent.Code,
				Log:           transactionFailedEvent.Log,
				Fee:           transactionFailedEvent.Fee,
				FeePayer:      transactionFailedEvent.FeePayer,
				FeeGranter:    transactionFailedEvent.FeeGranter,
				GasWanted:     transactionFailedEvent.GasWanted,
				GasUsed:       transactionFailedEvent.GasUsed,
				Memo:          transactionFailedEvent.Memo,
				TimeoutHeight: transactionFailedEvent.TimeoutHeight,
				Messages:      make([]view.TransactionRowMessage, 0),
			}
			txs = append(txs, row)

			transactionInfos[transactionFailedEvent.TxHash] = NewTransactionInfo(
				view.AccountTransactionBaseRow{
					Account:      "", // placeholder
					BlockHeight:  height,
					BlockHash:    "",                // placeholder
					BlockTime:    utctime.UTCTime{}, // placeholder
					Hash:         transactionFailedEvent.TxHash,
					MessageTypes: []string{},
					Success:      false,
				},
			)
			senders := projection.ParseSenderAddresses(transactionFailedEvent.Senders)
			for _, sender := range senders {
				transactionInfos[transactionFailedEvent.TxHash].AddAccount(sender)
			}
		} else if msgEvent, ok := event.(event_usecase.MsgEvent); ok {
			if _, exist := txMsgs[msgEvent.TxHash()]; !exist {
				txMsgs[msgEvent.TxHash()] = make([]event_usecase.MsgEvent, 0)
			}
			txMsgs[msgEvent.TxHash()] = append(txMsgs[msgEvent.TxHash()], msgEvent)

			transactionInfos[msgEvent.TxHash()].AddMessageTypes(msgEvent.MsgType())
		}
	}

	if len(txs) == 0 {
		if err := projection.UpdateLastHandledEventHeight(rdbTxHandle, height); err != nil {
			return fmt.Errorf("error updating last handled event height: %v", err)
		}

		if err := rdbTx.Commit(); err != nil {
			return fmt.Errorf("error committing changes: %v", err)
		}
		committed = true
		return nil
	}

	for i, tx := range txs {
		txs[i].BlockTime = blockTime
		txs[i].BlockHash = blockHash

		transactionInfos[tx.Hash].FillBlockInfo(blockHash, blockTime)

		for _, msg := range txMsgs[tx.Hash] {
			txs[i].Messages = append(txs[i].Messages, view.TransactionRowMessage{
				Type:    msg.MsgType(),
				Content: msg,
			})
		}
	}
	if insertErr := accountTransactionDataView.InsertAll(txs); insertErr != nil {
		return fmt.Errorf("error inserting account transaction data into view: %v", insertErr)
	}

	// Handle transaction messages
	for _, event := range events {
		if typedEvent, ok := event.(*event_usecase.MsgSend); ok {
			transactionInfos[typedEvent.TxHash()].AddAccount(typedEvent.ToAddress)

		} else if typedEvent, ok := event.(*event_usecase.MsgMultiSend); ok {
			for _, input := range typedEvent.Inputs {
				transactionInfos[typedEvent.TxHash()].AddAccount(input.Address)
			}
			for _, output := range typedEvent.Outputs {
				transactionInfos[typedEvent.TxHash()].AddAccount(output.Address)
			}

		} else if typedEvent, ok := event.(*event_usecase.MsgSetWithdrawAddress); ok {
			transactionInfos[typedEvent.TxHash()].AddAccount(typedEvent.DelegatorAddress)
			transactionInfos[typedEvent.TxHash()].AddAccount(typedEvent.WithdrawAddress)

		} else if typedEvent, ok := event.(*event_usecase.MsgWithdrawDelegatorReward); ok {
			transactionInfos[typedEvent.TxHash()].AddAccount(typedEvent.DelegatorAddress)

		} else if typedEvent, ok := event.(*event_usecase.MsgWithdrawValidatorCommission); ok {
			transactionInfos[typedEvent.TxHash()].AddAccount(typedEvent.RecipientAddress)

		} else if typedEvent, ok := event.(*event_usecase.MsgFundCommunityPool); ok {
			transactionInfos[typedEvent.TxHash()].AddAccount(typedEvent.Depositor)

		} else if typedEvent, ok := event.(*event_usecase.MsgFundCommunityPool); ok {
			transactionInfos[typedEvent.TxHash()].AddAccount(typedEvent.Depositor)

		} else if typedEvent, ok := event.(*event_usecase.MsgSubmitTextProposal); ok {
			transactionInfos[typedEvent.TxHash()].AddAccount(typedEvent.ProposerAddress)

		} else if typedEvent, ok := event.(*event_usecase.MsgSubmitParamChangeProposal); ok {
			transactionInfos[typedEvent.TxHash()].AddAccount(typedEvent.ProposerAddress)

		} else if typedEvent, ok := event.(*event_usecase.MsgSubmitCommunityPoolSpendProposal); ok {
			transactionInfos[typedEvent.TxHash()].AddAccount(typedEvent.ProposerAddress)

		} else if typedEvent, ok := event.(*event_usecase.MsgSubmitSoftwareUpgradeProposal); ok {
			transactionInfos[typedEvent.TxHash()].AddAccount(typedEvent.ProposerAddress)

		} else if typedEvent, ok := event.(*event_usecase.MsgSubmitCancelSoftwareUpgradeProposal); ok {
			transactionInfos[typedEvent.TxHash()].AddAccount(typedEvent.ProposerAddress)

		} else if typedEvent, ok := event.(*event_usecase.MsgDeposit); ok {
			transactionInfos[typedEvent.TxHash()].AddAccount(typedEvent.Depositor)

		} else if typedEvent, ok := event.(*event_usecase.MsgVote); ok {
			transactionInfos[typedEvent.TxHash()].AddAccount(typedEvent.Voter)

		} else if typedEvent, ok := event.(*event_usecase.MsgCreateValidator); ok {
			transactionInfos[typedEvent.TxHash()].AddAccount(typedEvent.DelegatorAddress)

			//} else if _, ok := event.(*event_usecase.MsgEditValidator); ok {

		} else if typedEvent, ok := event.(*event_usecase.MsgDelegate); ok {
			transactionInfos[typedEvent.TxHash()].AddAccount(typedEvent.DelegatorAddress)

		} else if typedEvent, ok := event.(*event_usecase.MsgUndelegate); ok {
			transactionInfos[typedEvent.TxHash()].AddAccount(typedEvent.DelegatorAddress)

		} else if typedEvent, ok := event.(*event_usecase.MsgBeginRedelegate); ok {
			transactionInfos[typedEvent.TxHash()].AddAccount(typedEvent.DelegatorAddress)

			//} else if _, ok := event.(*event_usecase.MsgUnjail); ok {
		}
	}

	accountTransactionRows := make([]view.AccountTransactionBaseRow, 0)
	for _, info := range transactionInfos {
		accountTransactionRows = append(accountTransactionRows, info.ToRows()...)
	}

	for _, row := range accountTransactionRows {
		if err := accountTransactionsTotalView.Increment(fmt.Sprintf("%s:-", row.Account), 1); err != nil {
			return fmt.Errorf("error incremnting total account transaction of account: %w", err)
		}
		for _, messageType := range row.MessageTypes {
			if err := accountTransactionsTotalView.Increment(
				fmt.Sprintf("%s:%s", row.Account, messageType), 1,
			); err != nil {
				return fmt.Errorf("error incremnting total account transaction message type of account: %w", err)
			}
		}

	}
	if err := accountTransactionsView.InsertAll(accountTransactionRows); err != nil {
		return fmt.Errorf("error inserting account message: %w", err)
	}

	if err := projection.UpdateLastHandledEventHeight(rdbTxHandle, height); err != nil {
		return fmt.Errorf("error updating last handled event height: %v", err)
	}

	if err := rdbTx.Commit(); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}
	committed = true
	return nil
}

func (projection *AccountTransaction) ParseSenderAddresses(senders []event_usecase.TransactionSigner) []string {
	addresses := make([]string, 0, len(senders))
	for _, sender := range senders {
		var address string
		if sender.IsMultiSig {
			addrPubKeys := make([][]byte, 0, len(sender.Pubkeys))
			for _, pubKey := range sender.Pubkeys {
				rawPubKey := base64.MustDecodeString(pubKey)
				addrPubKeys = append(addrPubKeys, rawPubKey)
			}
			address = tmcosmosutils.MustMultiSigAddressFromPubKeys(
				projection.accountAddressPrefix,
				addrPubKeys,
				*sender.MaybeThreshold,
				false,
			)
		} else {
			pubKey := base64.MustDecodeString(sender.Pubkeys[0])
			address = tmcosmosutils.MustAccountAddressFromPubKey(projection.accountAddressPrefix, pubKey)
		}
		addresses = append(addresses, address)
	}
	return addresses
}

type TransactionInfo struct {
	row              view.AccountTransactionBaseRow
	involvedAccounts map[string]bool // Set data structure
	messageTypes     map[string]bool // Set data structure
}

func NewTransactionInfo(row view.AccountTransactionBaseRow) *TransactionInfo {
	return &TransactionInfo{
		row: row,

		involvedAccounts: make(map[string]bool),
		messageTypes:     make(map[string]bool),
	}
}

func (info *TransactionInfo) IsAccountExist(address string) bool {
	return info.involvedAccounts[address]
}

func (info *TransactionInfo) AddAccount(address string) {
	info.involvedAccounts[address] = true
}

func (info *TransactionInfo) AddMessageTypes(messageTypes string) {
	info.messageTypes[messageTypes] = true
}

func (info *TransactionInfo) FillBlockInfo(blockHash string, blockTime utctime.UTCTime) {
	info.row.BlockHash = blockHash
	info.row.BlockTime = blockTime
}

func (info *TransactionInfo) ToRows() []view.AccountTransactionBaseRow {
	info.FillMessageTypes()

	rows := make([]view.AccountTransactionBaseRow, 0)
	for account := range info.involvedAccounts {
		clonedRow := info.row
		clonedRow.Account = account
		rows = append(rows, clonedRow)
	}

	return rows
}

func (info *TransactionInfo) FillMessageTypes() {
	info.row.MessageTypes = make([]string, 0)
	for messageType := range info.messageTypes {
		info.row.MessageTypes = append(info.row.MessageTypes, messageType)
	}
}
