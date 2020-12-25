package account

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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

	rdbConn   rdb.Conn
	logger    applogger.Logger
	lcdUrl    string // cosmos light client deaemon port : 1317 (default)
	baseDenom string // tbasecro, basecro
}

func NewAccount(logger applogger.Logger, rdbConn rdb.Conn, lcdUrl string, baseDenom string) *Account {
	return &Account{
		rdbprojectionbase.NewRDbBase(rdbConn.ToHandle(), "Account"),
		rdbConn,
		logger,
		lcdUrl,
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

func (projection *Account) GetAccountInfo(address string) (accountType string, accountAddress string, pubkey string, accountNumber string, sequenceNumber string, err error) {
	resp, err := http.Get(fmt.Sprintf("%s/cosmos/auth/v1beta1/accounts/%s", projection.lcdUrl, address))
	if err != nil {
		return "", "", "", "", "", fmt.Errorf("GetAccountInfo error http.Get %v", err)
	}
	defer resp.Body.Close()
	outputbytes, readerr := ioutil.ReadAll(resp.Body)
	if readerr != nil {
		return "", "", "", "", "", fmt.Errorf("GetAccountInfo error ioutil.ReadAll %v", readerr)
	}

	var myjson map[string]interface{}

	if err := json.Unmarshal(outputbytes, &myjson); err != nil {
		return "", "", "", "", "", fmt.Errorf("GetAccountInfo error json.Unmarshal %v", err)
	}

	var thisAccountType string
	var thisAddress string
	var thisPubkeyContainer map[string]interface{}
	var thisPubkey string
	var thisAccountNumber string
	var thisSequenceNumber string

	thisAccount := myjson["account"].(map[string]interface{})
	thisAccountTypeMeta := thisAccount["@type"].(string)
	thisAccountTypeName, thisAccountTypeNameOk := thisAccount["name"].(string)
	if !thisAccountTypeNameOk {
		thisAccountTypeName = ""
	}
	thisAccountType = fmt.Sprintf("%s %s", thisAccountTypeMeta, thisAccountTypeName)
	thisBaseAccount, thisBaseAccountOk := thisAccount["base_account"].(map[string]interface{})

	if !thisBaseAccountOk {
		// normal account
		thisAddress = thisAccount["address"].(string)
		thisPubkeyContainer = thisAccount["pub_key"].(map[string]interface{})
		thisPubkey = thisPubkeyContainer["key"].(string)
		thisAccountNumber = thisAccount["account_number"].(string)
		thisSequenceNumber = thisAccount["sequence"].(string)

	} else {
		// module account
		thisAddress = thisBaseAccount["address"].(string)
		thisPubkeyContainer, _ = thisBaseAccount["pub_key"].(map[string]interface{})
		thisAccountNumber = thisBaseAccount["account_number"].(string)
		thisSequenceNumber = thisBaseAccount["sequence"].(string)

	}
	return thisAccountType, thisAddress, thisPubkey, thisAccountNumber, thisSequenceNumber, nil

}

func (projection *Account) GetAccountBalance(targetAddress string, targetDenom string) (returnBalance string, returnDenom string, returnErr error) {
	resp, returnErr := http.Get(fmt.Sprintf("%s/cosmos/bank/v1beta1/balances/%s/%s", projection.lcdUrl, targetAddress, targetDenom))
	if returnErr != nil {
		return "", "", fmt.Errorf("GetAccountBalance error http.Get %v", returnErr)
	}
	defer resp.Body.Close()
	outputbytes, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return "", "", fmt.Errorf("GetAccountBalance error ioutil.ReadAll %v", returnErr)
	}

	var thisJson map[string]interface{}

	if err := json.Unmarshal(outputbytes, &thisJson); err != nil {
		return "", "", fmt.Errorf("GetAccountBalance error json.Unmarshal %v", err)
	}

	thisBalance := thisJson["balance"].(map[string]interface{})
	thisAmount := thisBalance["amount"].(string)
	thisDenom := thisBalance["denom"].(string)

	return thisAmount, thisDenom, nil
}

func (projection *Account) writeAccountInfo(accountsView *account_view.Accounts, whichaddress string) error {

	atype, aaddress, pubkey, aaccountnumber, asequenenumber, _ := projection.GetAccountInfo(whichaddress)
	// basetcro or basecro
	abalance, adenom, _ := projection.GetAccountBalance(whichaddress, projection.baseDenom)

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
