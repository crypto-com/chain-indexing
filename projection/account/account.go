package account

import (
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"

	cosmosapp_interface "github.com/crypto-com/chain-indexing/appinterface/cosmosapp"
	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/infrastructure/pg"
	github_mh "github.com/crypto-com/chain-indexing/infrastructure/pg/migrationhelper/github"
	appprojection "github.com/crypto-com/chain-indexing/projection"
	account_view "github.com/crypto-com/chain-indexing/projection/account/view"
	"github.com/crypto-com/chain-indexing/usecase/coin"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

// Account number, sequence number, balances are fetched from the latest state (regardless of current replaying height)
type Account struct {
	*rdbprojectionbase.Base

	rdbConn      rdb.Conn
	logger       applogger.Logger
	cosmosClient cosmosapp_interface.Client // cosmos light client deamon port : 1317 (default)

	config *appprojection.Config
}

func NewAccount(
	logger applogger.Logger,
	rdbConn rdb.Conn,
	cosmosClient cosmosapp_interface.Client,
	config *appprojection.Config,
) *Account {
	return &Account{
		rdbprojectionbase.NewRDbBaseWithOptions(
			rdbConn.ToHandle(),
			"Account",

			rdbprojectionbase.Options{
				MaybeConfigPtr: nil,
				MaybeTable:     nil,
				MaybeMigrationHelper: github_mh.NewGithubMigrationHelper(
					github_mh.Config{
						Config:     *config,
						ConnString: rdbConn.(*pg.PgxConn).ConnString(),
					},
					// MaybeSourceURL and MaybeDatabaseURL are empty, will generate these fields for the projection
					// in usecase/projection/base.go
					nil,
					nil,
				),
			},
		),

		rdbConn,
		logger,
		cosmosClient,
		config,
	}
}

var (
	NewAccountsView              = account_view.NewAccountsView
	UpdateLastHandledEventHeight = (*Account).UpdateLastHandledEventHeight
)

func (_ *Account) GetEventsToListen() []string {
	return []string{
		// TODO: Genesis account
		event_usecase.ACCOUNT_TRANSFERRED,
	}
}

func (projection *Account) OnInit() error {
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
		if accountCreatedEvent, ok := event.(*event_usecase.AccountTransferred); ok {
			if handleErr := projection.handleAccountCreatedEvent(accountsView, accountCreatedEvent); handleErr != nil {
				return fmt.Errorf("error handling AccountCreatedEvent: %v", handleErr)
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

func (projection *Account) handleAccountCreatedEvent(accountsView account_view.Accounts, event *event_usecase.AccountTransferred) error {

	recipienterr := projection.writeAccountInfo(accountsView, event.Recipient)
	if recipienterr != nil {
		return recipienterr
	}

	sendererr := projection.writeAccountInfo(accountsView, event.Sender)
	if sendererr != nil {
		return sendererr
	}

	return nil
}

func (projection *Account) getAccountInfo(address string) (*cosmosapp_interface.Account, error) {
	var accountInfo, accountInfoError = projection.cosmosClient.Account(address)
	if accountInfoError != nil {
		return nil, accountInfoError
	}

	return accountInfo, nil
}

func (projection *Account) getAccountBalances(targetAddress string) (coin.Coins, error) {
	var balanceInfo, balanceInfoError = projection.cosmosClient.Balances(targetAddress)
	if balanceInfoError != nil {
		return nil, balanceInfoError
	}

	return balanceInfo, nil
}

func (projection *Account) writeAccountInfo(accountsView account_view.Accounts, address string) error {
	accountInfo, err := projection.getAccountInfo(address)
	if err != nil {
		return err
	}

	accountType := accountInfo.Type
	var name *string
	if accountInfo.Type == cosmosapp_interface.ACCOUNT_MODULE {
		name = &accountInfo.MaybeModuleAccount.Name
	}
	var pubkey *string
	if accountInfo.MaybePubkey != nil {
		pubkey = &accountInfo.MaybePubkey.Key
	}
	accountNumber := accountInfo.AccountNumber
	sequenceNumber := accountInfo.Sequence

	balances, err := projection.getAccountBalances(address)
	if err != nil {
		return err
	}
	if err := accountsView.Upsert(&account_view.AccountRow{
		Type:           accountType,
		Address:        address,
		MaybeName:      name,
		MaybePubkey:    pubkey,
		AccountNumber:  accountNumber,
		SequenceNumber: sequenceNumber,
		Balance:        balances,
	}); err != nil {
		return err
	}

	return nil
}
