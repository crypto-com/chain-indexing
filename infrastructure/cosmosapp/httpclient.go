package cosmosapp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/crypto-com/chain-indexing/usecase/coin"

	jsoniter "github.com/json-iterator/go"

	cosmosapp_interface "github.com/crypto-com/chain-indexing/appinterface/cosmosapp"
)

var _ cosmosapp_interface.Client = &HTTPClient{}

type HTTPClient struct {
	httpClient *http.Client
	rpcUrl     string
}

// NewHTTPClient returns a new HTTPClient for tendermint request
func NewHTTPClient(rpcUrl string) *HTTPClient {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &HTTPClient{
		httpClient,
		strings.TrimSuffix(rpcUrl, "/"),
	}
}

func (client *HTTPClient) Account(accountAddress string) (*cosmosapp_interface.Account, error) {
	fmt.Println(accountAddress)
	rawRespBody, err := client.request(
		fmt.Sprintf("%s/%s", client.url("auth", "accounts"), accountAddress), "",
	)
	if err != nil {
		return nil, err
	}
	defer rawRespBody.Close()

	var accountResp AccountResp
	if err := jsoniter.NewDecoder(rawRespBody).Decode(&accountResp); err != nil {
		return nil, err
	}
	rawAccount := accountResp.Account

	var account cosmosapp_interface.Account
	switch rawAccount.Type {
	case cosmosapp_interface.ACCOUNT_BASE:
		account = cosmosapp_interface.Account{
			Type:          rawAccount.Type,
			Address:       *rawAccount.MaybeAddress,
			MaybePubkey:   rawAccount.MaybePubKey,
			AccountNumber: *rawAccount.MaybeAccountNumber,
			Sequence:      *rawAccount.MaybeSequence,

			MaybeModuleAccount:            nil,
			MaybeDelayedVestingAccount:    nil,
			MaybeContinuousVestingAccount: nil,
			MaybePeriodicVestingAccount:   nil,
		}
	case cosmosapp_interface.ACCOUNT_MODULE:
		account = cosmosapp_interface.Account{
			Type:          rawAccount.Type,
			Address:       rawAccount.MaybeBaseAccount.Address,
			MaybePubkey:   nil,
			AccountNumber: rawAccount.MaybeBaseAccount.AccountNumber,
			Sequence:      rawAccount.MaybeBaseAccount.Sequence,
			MaybeModuleAccount: &cosmosapp_interface.ModuleAccount{
				Name:        *rawAccount.MaybeName,
				Permissions: rawAccount.MaybePermissions,
			},

			MaybeDelayedVestingAccount:    nil,
			MaybeContinuousVestingAccount: nil,
			MaybePeriodicVestingAccount:   nil,
		}
	case "/cosmos.vesting.v1beta1.DelayedVestingAccount":
		account = cosmosapp_interface.Account{
			Type:          rawAccount.Type,
			Address:       rawAccount.MaybeBaseVestingAccount.BaseAccount.Address,
			MaybePubkey:   rawAccount.MaybeBaseVestingAccount.BaseAccount.MaybePubKey,
			AccountNumber: rawAccount.MaybeBaseVestingAccount.BaseAccount.AccountNumber,
			Sequence:      rawAccount.MaybeBaseVestingAccount.BaseAccount.Sequence,

			MaybeModuleAccount: nil,
			MaybeDelayedVestingAccount: &cosmosapp_interface.DelayedVestingAccount{
				OriginalVesting:  rawAccount.MaybeBaseVestingAccount.OriginalVesting,
				DelegatedFree:    rawAccount.MaybeBaseVestingAccount.DelegatedFree,
				DelegatedVesting: rawAccount.MaybeBaseVestingAccount.DelegatedVesting,
				EndTime:          rawAccount.MaybeBaseVestingAccount.EndTime,
			},
			MaybeContinuousVestingAccount: nil,
			MaybePeriodicVestingAccount:   nil,
		}

	case "/cosmos.vesting.v1beta1.ContinuousVestingAccount":
		account = cosmosapp_interface.Account{
			Type:          rawAccount.Type,
			Address:       rawAccount.MaybeBaseVestingAccount.BaseAccount.Address,
			MaybePubkey:   rawAccount.MaybeBaseVestingAccount.BaseAccount.MaybePubKey,
			AccountNumber: rawAccount.MaybeBaseVestingAccount.BaseAccount.AccountNumber,
			Sequence:      rawAccount.MaybeBaseVestingAccount.BaseAccount.Sequence,

			MaybeModuleAccount:         nil,
			MaybeDelayedVestingAccount: nil,
			MaybeContinuousVestingAccount: &cosmosapp_interface.ContinuousVestingAccount{
				OriginalVesting:  rawAccount.MaybeBaseVestingAccount.OriginalVesting,
				DelegatedFree:    rawAccount.MaybeBaseVestingAccount.DelegatedFree,
				DelegatedVesting: rawAccount.MaybeBaseVestingAccount.DelegatedVesting,
				StartTime:        *rawAccount.MaybeStartTime,
				EndTime:          rawAccount.MaybeBaseVestingAccount.EndTime,
			},
			MaybePeriodicVestingAccount: nil,
		}
	case cosmosapp_interface.ACCOUNT_VESTING_PERIODIC:
		account = cosmosapp_interface.Account{
			Type:          rawAccount.Type,
			Address:       rawAccount.MaybeBaseVestingAccount.BaseAccount.Address,
			MaybePubkey:   rawAccount.MaybeBaseVestingAccount.BaseAccount.MaybePubKey,
			AccountNumber: rawAccount.MaybeBaseVestingAccount.BaseAccount.AccountNumber,
			Sequence:      rawAccount.MaybeBaseVestingAccount.BaseAccount.Sequence,

			MaybeModuleAccount:            nil,
			MaybeDelayedVestingAccount:    nil,
			MaybeContinuousVestingAccount: nil,
			MaybePeriodicVestingAccount: &cosmosapp_interface.PeriodicVestingAccount{
				OriginalVesting:  rawAccount.MaybeBaseVestingAccount.OriginalVesting,
				DelegatedFree:    rawAccount.MaybeBaseVestingAccount.DelegatedFree,
				DelegatedVesting: rawAccount.MaybeBaseVestingAccount.DelegatedVesting,
				StartTime:        *rawAccount.MaybeStartTime,
				EndTime:          rawAccount.MaybeBaseVestingAccount.EndTime,
				VestingPeriods:   rawAccount.MaybeVestingPeriods,
			},
		}
	default:
		return nil, fmt.Errorf("unrecognized account type: %s", rawAccount.Type)
	}

	return &account, nil
}

func (client *HTTPClient) Balances(accountAddress string) (coin.Coins, error) {
	resp := &BankBalancesResp{
		Pagination: cosmosapp_interface.Pagination{
			MaybeNextKey: nil,
			Total:        "",
		},
	}
	balances := coin.NewEmptyCoins()
	for {
		url := fmt.Sprintf("%s/%s", client.url("bank", "balances"), accountAddress)
		if resp.Pagination.MaybeNextKey != nil {
			url = fmt.Sprintf("%s?pagination.key=%s", url, *resp.Pagination.MaybeNextKey)
		}

		rawRespBody, err := client.request(url)
		if err != nil {
			return nil, err
		}
		defer rawRespBody.Close()

		if err := jsoniter.NewDecoder(rawRespBody).Decode(&resp); err != nil {
			return nil, err
		}
		for _, balanceKVPair := range resp.BankBalanceResponses {
			balance, coinErr := coin.NewCoinFromString(balanceKVPair.Denom, balanceKVPair.Amount)
			if coinErr != nil {
				return nil, coinErr
			}
			balances = balances.Add(balance)
		}

		if resp.Pagination.MaybeNextKey == nil {
			break
		}
	}

	return balances, nil
}

func (client *HTTPClient) BalanceByDenom(accountAddress string, denom string) (*coin.Coin, error) {
	rawRespBody, err := client.request(
		fmt.Sprintf("%s/%s/%s", client.url("bank", "balances"), accountAddress, denom), "",
	)
	if err != nil {
		return nil, err
	}
	defer rawRespBody.Close()

	outputBytes, readErr := ioutil.ReadAll(rawRespBody)
	if readErr != nil {
		return nil, fmt.Errorf("GetAccountInfo error ioutil.ReadAll %v", readErr)
	}
	var rawResp map[string]interface{}

	if unmarshalErr := json.Unmarshal(outputBytes, &rawResp); unmarshalErr != nil {
		return nil, unmarshalErr
	}

	rawBalance := rawResp["balance"].(map[string]interface{})
	amount := rawBalance["amount"].(string)

	balance, err := coin.NewCoinFromString(denom, amount)
	if err != nil {
		return nil, err
	}
	return &balance, nil
}

func (client *HTTPClient) Validator(validatorAddress string) (*cosmosapp_interface.Validator, error) {
	rawRespBody, err := client.request(
		fmt.Sprintf("%s/%s", client.url("staking", "validators"), validatorAddress), "",
	)
	if err != nil {
		return nil, err
	}
	defer rawRespBody.Close()

	var validatorResp ValidatorResp
	if err := jsoniter.NewDecoder(rawRespBody).Decode(&validatorResp); err != nil {
		return nil, err
	}

	return &validatorResp.Validator, nil
}

func (client *HTTPClient) Delegation(
	delegator string, validator string,
) (*cosmosapp_interface.DelegationResponse, error) {
	resp := &DelegationResp{
		Pagination: cosmosapp_interface.Pagination{
			MaybeNextKey: nil,
			Total:        "",
		},
	}
	for {
		url := fmt.Sprintf("%s/%s", client.url("staking", "delegations"), delegator)
		if resp.Pagination.MaybeNextKey != nil {
			url = fmt.Sprintf("%s?pagination.key=%s", url, *resp.Pagination.MaybeNextKey)
		}

		rawRespBody, err := client.request(url)
		if err != nil {
			return nil, err
		}
		defer rawRespBody.Close()

		if err := jsoniter.NewDecoder(rawRespBody).Decode(&resp); err != nil {
			return nil, err
		}
		for _, delegation := range resp.DelegationResponses {
			if delegation.Delegation.DelegatorAddress == delegator && delegation.Delegation.ValidatorAddress == validator {
				return &delegation, nil
			}
		}

		if resp.Pagination.MaybeNextKey == nil {
			break
		}
	}

	return nil, nil
}

func (client *HTTPClient) url(module string, method string) string {
	return fmt.Sprintf("cosmos/%s/v1beta1/%s", module, method)
}

// request construct tendermint url and issues an HTTP request
// returns the success http Body
func (client *HTTPClient) request(method string, queryString ...string) (io.ReadCloser, error) {
	var err error

	url := client.rpcUrl + "/" + method
	if len(queryString) > 0 {
		url += "?" + queryString[0]
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request with context: %v", err)
	}
	rawResp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error requesting Tendermint %s endpoint: %v", url, err)
	}

	if rawResp.StatusCode != 200 {
		rawResp.Body.Close()
		return nil, fmt.Errorf("error requesting Tendermint %s endpoint: %s", method, rawResp.Status)
	}

	return rawResp.Body, nil
}

type ValidatorResp struct {
	Validator cosmosapp_interface.Validator `json:"validator"`
}

type DelegationResp struct {
	DelegationResponses []cosmosapp_interface.DelegationResponse `json:"delegation_responses"`
	Pagination          cosmosapp_interface.Pagination           `json:"pagination"`
}

type AccountResp struct {
	Account Account
}

type Account struct {
	// Common fields
	Type string `json:"@type"`

	// Module account
	MaybeName        *string      `json:"name"`
	MaybeBaseAccount *BaseAccount `json:"base_account"`
	MaybePermissions []string     `json:"permissions"`

	// Vesting account common fields
	MaybeBaseVestingAccount *BaseVestingAccount `json:"base_vesting_account"`
	// Continuous vesting account
	MaybeStartTime *string `json:"start_time"`
	// Periodic vesting account
	MaybeVestingPeriods []cosmosapp_interface.VestingPeriod `json:"vesting_periods"`

	// User account
	MaybeAddress       *string                     `json:"address"`
	MaybePubKey        *cosmosapp_interface.PubKey `json:"pub_key"`
	MaybeAccountNumber *string                     `json:"account_number"`
	MaybeSequence      *string                     `json:"sequence"`
}

type BaseVestingAccount struct {
	BaseAccount      BaseAccount                          `json:"base_account"`
	OriginalVesting  []cosmosapp_interface.VestingBalance `json:"original_vesting"`
	DelegatedFree    []cosmosapp_interface.VestingBalance `json:"delegated_free"`
	DelegatedVesting []cosmosapp_interface.VestingBalance `json:"delegated_vesting"`
	EndTime          string                               `json:"end_time"`
}

type BaseAccount struct {
	Address       string                      `json:"address"`
	MaybePubKey   *cosmosapp_interface.PubKey `json:"pub_key"`
	AccountNumber string                      `json:"account_number"`
	Sequence      string                      `json:"sequence"`
}

type BankBalancesResp struct {
	BankBalanceResponses []cosmosapp_interface.BankBalance `json:"balances"`
	Pagination           cosmosapp_interface.Pagination    `json:"pagination"`
}
