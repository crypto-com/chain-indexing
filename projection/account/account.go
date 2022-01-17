package account

import (
	"fmt"
	"strconv"

	"github.com/crypto-com/chain-indexing/external/tmcosmosutils"
	"github.com/crypto-com/chain-indexing/internal/base64"
	"github.com/crypto-com/chain-indexing/usecase/model"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"

	cosmosapp_interface "github.com/crypto-com/chain-indexing/appinterface/cosmosapp"
	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper"
	account_view "github.com/crypto-com/chain-indexing/projection/account/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

// Account number, sequence number, balances are fetched from the latest state (regardless of current replaying height)
type Account struct {
	*rdbprojectionbase.Base

	accountAddressPrefix string

	rdbConn      rdb.Conn
	logger       applogger.Logger
	cosmosClient cosmosapp_interface.Client // cosmos light client deamon port : 1317 (default)

	migrationHelper migrationhelper.MigrationHelper
}

func NewAccount(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	accountAddressPrefix string,
	cosmosClient cosmosapp_interface.Client,
	migrationHelper migrationhelper.MigrationHelper,
) *Account {
	return &Account{
		rdbprojectionbase.NewRDbBaseWithOptions(
			rdbConn.ToHandle(),
			"Account",
			rdbprojectionbase.Options{
				MaybeConfigPtr: nil,
				MaybeTable:     nil,
			},
		),

		accountAddressPrefix,

		rdbConn,
		logger,
		cosmosClient,

		migrationHelper,
	}
}

var (
	NewAccountsView              = account_view.NewAccountsView
	UpdateLastHandledEventHeight = (*Account).UpdateLastHandledEventHeight
)

func (_ *Account) GetEventsToListen() []string {
	return []string{
		event_usecase.GENESIS_CREATED,
		event_usecase.COIN_SPENT,
		event_usecase.COIN_RECEIVED,
		event_usecase.COIN_MINT,
		event_usecase.COIN_BURN,
		event_usecase.TRANSACTION_FAILED,
	}
}

func (projection *Account) OnInit() error {
	if projection.migrationHelper != nil {
		projection.migrationHelper.Migrate()
	}
	return nil
}

func (projection *Account) HandleEvents(height int64, events []event_entity.Event) error {
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

	accountsView := NewAccountsView(rdbTxHandle)

	for _, event := range events {
		if genesisCreatedEvent, ok := event.(*event_usecase.GenesisCreated); ok {
			if handleErr := projection.handleGenesisCreatedEvent(accountsView, genesisCreatedEvent); handleErr != nil {
				return fmt.Errorf("error handling GenesisCreatedEvent: %v", handleErr)
			}
		} else if coinSpentEvent, ok := event.(*event_usecase.CoinSpent); ok {
			if handleErr := projection.decrementUsableBalance(accountsView, coinSpentEvent.Address, coinSpentEvent.Amount); handleErr != nil {
				return fmt.Errorf("error handling CoinSpentEvent: %v", handleErr)
			}
			if handleErr := projection.addAccountEvent(
				accountsView,
				coinSpentEvent.Address,
				coinSpentEvent,
				coinSpentEvent.EventName,
				coinSpentEvent.BlockHeight,
			); handleErr != nil {
				return fmt.Errorf("error handling CoinSpentEvent: %v", handleErr)
			}
		} else if coinReceivedEvent, ok := event.(*event_usecase.CoinReceived); ok {
			if handleErr := projection.incrementUsableBalance(accountsView, coinReceivedEvent.Address, coinReceivedEvent.Amount); handleErr != nil {
				return fmt.Errorf("error handling CoinReceivedEvent: %v", handleErr)
			}

			if handleErr := projection.addAccountEvent(
				accountsView,
				coinReceivedEvent.Address,
				coinReceivedEvent,
				coinReceivedEvent.EventName,
				coinReceivedEvent.BlockHeight,
			); handleErr != nil {
				return fmt.Errorf("error handling CoinReceivedEvent: %v", handleErr)
			}
		} else if coinMintEvent, ok := event.(*event_usecase.CoinMint); ok {
			if handleErr := projection.incrementUsableBalance(accountsView, coinMintEvent.Address, coinMintEvent.Amount); handleErr != nil {
				return fmt.Errorf("error handling coinMintEvent: %v", handleErr)
			}
			if handleErr := projection.addAccountEvent(
				accountsView,
				coinMintEvent.Address,
				coinMintEvent,
				coinMintEvent.EventName,
				coinMintEvent.BlockHeight,
			); handleErr != nil {
				return fmt.Errorf("error handling coinMintEvent: %v", handleErr)
			}
		} else if coinBurnEvent, ok := event.(*event_usecase.CoinBurn); ok {
			if handleErr := projection.decrementUsableBalance(accountsView, coinBurnEvent.Address, coinBurnEvent.Amount); handleErr != nil {
				return fmt.Errorf("error handling coinBurnEvent: %v", handleErr)
			}
			if handleErr := projection.addAccountEvent(
				accountsView,
				coinBurnEvent.Address,
				coinBurnEvent,
				coinBurnEvent.EventName,
				coinBurnEvent.BlockHeight,
			); handleErr != nil {
				return fmt.Errorf("error handling coinBurnEvent: %v", handleErr)
			}
		} else if transactionFailedEvent, ok := event.(*event_usecase.TransactionFailed); ok {
			feePayer := ""
			if transactionFailedEvent.FeeGranter != "" {
				// if set, the fee payer (either the first signer or the value of the payer field) requests that a fee grant be used
				// to pay fees instead of the fee payer's own balance. If an appropriate fee grant does not exist or the chain does
				// not support fee grants, this will fail
				feePayer = transactionFailedEvent.FeeGranter
			} else if transactionFailedEvent.FeePayer != "" {
				// if unset, the first signer is responsible for paying the fees. If set, the specified account must pay the fees.
				// the payer must be a tx signer (and thus have signed this field in AuthInfo).
				feePayer = transactionFailedEvent.FeePayer
			} else {
				signers := projection.ParseSenderAddresses(transactionFailedEvent.Signers)
				feePayer = signers[0]
			}
			if feePayer != "" {
				for i := range transactionFailedEvent.Fee {
					if handleErr := projection.decrementUsableBalance(
						accountsView,
						feePayer,
						coin.Coins{transactionFailedEvent.Fee[i]},
					); handleErr != nil {
						return fmt.Errorf("error handling TransactionFaileddEvent: %v", handleErr)
					}
					if handleErr := projection.addAccountEvent(
						accountsView,
						feePayer,
						transactionFailedEvent,
						transactionFailedEvent.EventName,
						transactionFailedEvent.BlockHeight,
					); handleErr != nil {
						return fmt.Errorf("error handling TransactionFailedEvent: %v", handleErr)
					}
				}
			}
		}
	}

	if err = UpdateLastHandledEventHeight(projection, rdbTxHandle, height); err != nil {
		return fmt.Errorf("error updating last handled event height: %v", err)
	}

	if err = rdbTx.Commit(); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}
	committed = true

	return nil
}

func (projection *Account) handleGenesisCreatedEvent(accountsView account_view.Accounts, event *event_usecase.GenesisCreated) error {

	for _, account := range event.Genesis.AppState.Bank.Balances {
		for _, accountBalance := range account.Coins {
			amount, err := strconv.ParseInt(accountBalance.Amount, 10, 0)
			if err != nil {
				return err
			}
			if err := accountsView.IncrementUsableBalance(
				account.Address, accountBalance.Denom, amount,
			); err != nil {
				return err
			}

			if handleErr := projection.addAccountEvent(
				accountsView,
				account.Address,
				event,
				event.EventName,
				event.BlockHeight,
			); handleErr != nil {
				return fmt.Errorf("error handling CoinSpentEvent: %v", handleErr)
			}
		}
	}

	return nil
}

func (projection *Account) incrementUsableBalance(accountsView account_view.Accounts, address string, amount coin.Coins) error {
	for _, increment := range amount {
		if err := accountsView.IncrementUsableBalance(
			address, increment.Denom, increment.Amount.Int64(),
		); err != nil {
			return err
		}
	}

	return nil
}

func (projection *Account) decrementUsableBalance(accountsView account_view.Accounts, address string, amount coin.Coins) error {
	for _, decrement := range amount {
		if err := accountsView.DecrementUsableBalance(
			address, decrement.Denom, decrement.Amount.Int64(),
		); err != nil {
			return err
		}
	}

	return nil
}

func (projection *Account) addAccountEvent(
	accountsView account_view.Accounts,
	address string,
	data interface{},
	eventType string,
	blockHeight int64,
) error {

	if err := accountsView.InsertAccountEvent(
		account_view.AccountEvent{
			Address:     address,
			Type:        eventType,
			Data:        data,
			BlockHeight: blockHeight,
		},
	); err != nil {
		return err
	}

	return nil
}

func (projection *Account) ParseSenderAddresses(senders []model.TransactionSigner) []string {
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
