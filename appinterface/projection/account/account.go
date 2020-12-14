package account

import (
	"fmt"

	account_view "github.com/crypto-com/chain-indexing/appinterface/projection/account/view"
	usecase_coin "github.com/crypto-com/chain-indexing/usecase/coin"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

type Account struct {
	*rdbprojectionbase.Base

	rdbConn rdb.Conn
	logger  applogger.Logger
}

func NewAccount(logger applogger.Logger, rdbConn rdb.Conn) *Account {
	return &Account{
		rdbprojectionbase.NewRDbBase(rdbConn.ToHandle(), "Account"),

		rdbConn,
		logger,
	}
}

func (_ *Account) GetEventsToListen() []string {
	return []string{event_usecase.ACCOUNT_TRANSFERRED}
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

	accountsView := account_view.NewAccounts(rdbTxHandle)

	for _, event := range events {
		if accountCreatedEvent, ok := event.(*event_usecase.AccountTransferred); ok {
			if handleErr := projection.handleAccountCreatedEvent(accountsView, accountCreatedEvent); handleErr != nil {
				return fmt.Errorf("error handling AccountCreatedEvent: %v", handleErr)
			}
		}
	}

	if err = projection.UpdateLastHandledEventHeight(rdbTxHandle, height); err != nil {
		return fmt.Errorf("error updating last handled event height: %v", err)
	}

	if err = rdbTx.Commit(); err != nil {
		return fmt.Errorf("error committing changes: %v", err)
	}
	committed = true

	return nil
}

func (projection *Account) handleAccountCreatedEvent(accountsView *account_view.Accounts, event *event_usecase.AccountTransferred) error {

	add_err := projection.addAmount(accountsView, event.Recipient, event.Amount.ToBigInt().Int64())
	if add_err != nil {
		return add_err
	}

	subtract_err := projection.addAmount(accountsView, event.Sender, -event.Amount.ToBigInt().Int64())
	if subtract_err != nil {
		return subtract_err
	}

	return nil
}

func (projection *Account) addAmount(accountsView *account_view.Accounts, recipient string, amount int64) error {
	newvalue, coinerr := usecase_coin.NewCoinFromInt(0)
	if coinerr != nil {
		return coinerr
	}
	var foundAccount account_view.Account
	if account, err := accountsView.FindBy(&account_view.AccountIdentity{
		MaybeAddress: recipient,
	}); err == nil {
		foundAccount = *account
		oldcoin, oldcoinerr := usecase_coin.NewCoinFromInt(foundAccount.AccountBalance)
		if oldcoinerr != nil {
			return oldcoinerr
		}

		amountcoin, amountcoinerr := usecase_coin.NewCoinFromInt(amount)
		if amountcoinerr != nil {
			return amountcoinerr
		}
		newcoin, newcoinerr := oldcoin.Add(amountcoin)
		if newcoinerr != nil {
			return newcoinerr
		}
		newvalue = newcoin
	}

	if err := accountsView.Upsert(&account_view.Account{
		AccountAddress: recipient,
		AccountBalance: newvalue.ToBigInt().Int64(),
		AccountDenom:   string("basecro"),
	}); err != nil {
		return err
	}

	return nil
}
