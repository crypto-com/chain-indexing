package account

import (
	"fmt"
	"strconv"

	cosmosapp_interface "github.com/crypto-com/chain-indexing/appinterface/cosmosapp"
	cosmosapp_infrastructure "github.com/crypto-com/chain-indexing/infrastructure/cosmosapp"

	account_view "github.com/crypto-com/chain-indexing/appinterface/projection/account/view"

	"github.com/crypto-com/chain-indexing/appinterface/projection/rdbprojectionbase"
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	event_entity "github.com/crypto-com/chain-indexing/entity/event"
	applogger "github.com/crypto-com/chain-indexing/internal/logger"
	event_usecase "github.com/crypto-com/chain-indexing/usecase/event"
)

func ConvertToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

type Account struct {
	*rdbprojectionbase.Base

	rdbConn    rdb.Conn
	logger     applogger.Logger
	httpclinet *cosmosapp_infrastructure.HTTPClient // cosmos light client deaemon port : 1317 (default)
	baseDenom  string                               // tbasecro, basecro
}

func NewAccount(logger applogger.Logger, rdbConn rdb.Conn, httpclient *cosmosapp_infrastructure.HTTPClient, baseDenom string) *Account {
	return &Account{
		rdbprojectionbase.NewRDbBase(rdbConn.ToHandle(), "Account"),
		rdbConn,
		logger,
		httpclient,
		baseDenom,
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

	var accountInfo, accountInfoError = projection.httpclinet.Account(address)
	if accountInfoError != nil {
		return nil, accountInfoError
	}

	return accountInfo, nil

}

func (projection *Account) getAccountBalance(targetAddress string, targetDenom string) (*cosmosapp_interface.AccountBalance, error) {

	var balanceInfo, balanceInfoError = projection.httpclinet.Balance(targetAddress, targetDenom)
	if balanceInfoError != nil {
		return nil, balanceInfoError
	}

	return balanceInfo, nil

}

func (projection *Account) writeAccountInfo(accountsView *account_view.Accounts, whichaddress string) error {

	accountinfo, accounterror := projection.getAccountInfo(whichaddress)
	if accounterror != nil {
		return accounterror
	}
	atype := accountinfo.AccountType
	aaddress := accountinfo.AccountAddress
	pubkey := accountinfo.Pubkey
	aaccountnumber := accountinfo.AccountNumber
	asequenenumber := accountinfo.SequenceNumber
	//atype, aaddress, pubkey, aaccountnumber, asequenenumber
	// basetcro or basecro
	balanceinfo, balanceinfoerror := projection.getAccountBalance(whichaddress, projection.baseDenom)
	if balanceinfoerror != nil {
		return balanceinfoerror
	}
	abalance := balanceinfo.AccountAmount
	adenom := balanceinfo.AccountDenom
	if err := accountsView.Upsert(&account_view.Account{
		AccountType:    atype,
		AccountAddress: aaddress,
		Pubkey:         pubkey,
		AccountNumber:  ConvertToInt64(aaccountnumber),
		SequenceNumber: ConvertToInt64(asequenenumber),
		AccountBalance: ConvertToInt64(abalance),
		AccountDenom:   adenom,
	}); err != nil {
		return err
	}

	return nil
}
