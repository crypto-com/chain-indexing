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

	bondingDenom string
}

// NewHTTPClient returns a new HTTPClient for tendermint request
func NewHTTPClient(rpcUrl string, bondingDenom string) *HTTPClient {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &HTTPClient{
		httpClient,
		strings.TrimSuffix(rpcUrl, "/"),

		bondingDenom,
	}
}

func (client *HTTPClient) Account(accountAddress string) (*cosmosapp_interface.Account, error) {
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
		Pagination: Pagination{
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

func (client *HTTPClient) BondedBalance(accountAddress string) (coin.Coins, error) {
	resp := &DelegationResp{
		MaybePagination: &Pagination{
			MaybeNextKey: nil,
			Total:        "",
		},
	}
	balance := coin.NewEmptyCoins()
	for {
		url := fmt.Sprintf("%s/%s", client.url("staking", "delegations"), accountAddress)
		if resp.MaybePagination.MaybeNextKey != nil {
			url = fmt.Sprintf("%s?pagination.key=%s", url, *resp.MaybePagination.MaybeNextKey)
		}

		rawRespBody, statusCode, err := client.rawRequest(url)
		if err != nil {
			return nil, err
		}
		defer rawRespBody.Close()

		if decodeErr := jsoniter.NewDecoder(rawRespBody).Decode(&resp); decodeErr != nil {
			return nil, decodeErr
		}
		if resp.MaybeCode != nil && *resp.MaybeCode == 2 {
			return nil, cosmosapp_interface.ErrAccountNotFound
		}
		if statusCode != 200 {
			return nil, fmt.Errorf("error requesting Cosmos %s endpoint: %v", url, err)
		}
		for _, delegation := range resp.MaybeDelegationResponses {
			delegatedCoin, coinErr := coin.NewCoinFromString(delegation.Balance.Denom, delegation.Balance.Amount)
			if coinErr != nil {
				return nil, fmt.Errorf("error parsing Coin from delegation balance: %v", coinErr)
			}
			balance = balance.Add(delegatedCoin)
		}

		if resp.MaybePagination.MaybeNextKey == nil {
			break
		}
	}

	return balance, nil
}

func (client *HTTPClient) RedelegatingBalance(accountAddress string) (coin.Coins, error) {
	resp := &UnbondingResp{
		Pagination: Pagination{
			MaybeNextKey: nil,
			Total:        "",
		},
	}
	balance := coin.NewEmptyCoins()
	for {
		url := fmt.Sprintf(
			"%s/%s/redelegations", client.url("staking", "delegators"), accountAddress,
		)
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
		for _, unbonding := range resp.UnbondingResponses {
			for _, entry := range unbonding.Entries {
				unbondingCoin, coinErr := coin.NewCoinFromString(client.bondingDenom, entry.Balance)
				if coinErr != nil {
					return nil, fmt.Errorf("error parsing Coin from unbonding balance: %v", coinErr)
				}
				balance = balance.Add(unbondingCoin)
			}
		}

		if resp.Pagination.MaybeNextKey == nil {
			break
		}
	}

	return balance, nil
}

func (client *HTTPClient) UnbondingBalance(accountAddress string) (coin.Coins, error) {
	resp := &UnbondingResp{
		Pagination: Pagination{
			MaybeNextKey: nil,
			Total:        "",
		},
	}
	balance := coin.NewEmptyCoins()
	for {
		url := fmt.Sprintf(
			"%s/%s/unbonding_delegations", client.url("staking", "delegators"), accountAddress,
		)
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
		for _, unbonding := range resp.UnbondingResponses {
			for _, entry := range unbonding.Entries {
				unbondingCoin, coinErr := coin.NewCoinFromString(client.bondingDenom, entry.Balance)
				if coinErr != nil {
					return nil, fmt.Errorf("error parsing Coin from unbonding balance: %v", coinErr)
				}
				balance = balance.Add(unbondingCoin)
			}
		}

		if resp.Pagination.MaybeNextKey == nil {
			break
		}
	}

	return balance, nil
}

func (client *HTTPClient) TotalRewards(accountAddress string) (coin.DecCoins, error) {
	rawRespBody, err := client.request(
		fmt.Sprintf(
			"%s/%s/rewards",
			client.url("distribution", "delegators"),
			accountAddress,
		), "",
	)
	if err != nil {
		return nil, err
	}
	defer rawRespBody.Close()

	var delegatorRewardResp DelegatorRewardResp
	if err := jsoniter.NewDecoder(rawRespBody).Decode(&delegatorRewardResp); err != nil {
		return nil, err
	}

	rewards := coin.NewEmptyDecCoins()
	for _, total := range delegatorRewardResp.Total {
		rewardCoin, coinErr := coin.NewDecCoinFromString(total.Denom, total.Amount)
		if coinErr != nil {
			return nil, fmt.Errorf("error parsing Coin from total reward: %v", coinErr)
		}
		rewards = rewards.Add(rewardCoin)
	}
	return rewards, nil
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

func (client *HTTPClient) Commission(validatorAddress string) (coin.DecCoins, error) {
	rawRespBody, err := client.request(
		fmt.Sprintf("%s/%s/commission",
			client.url("distribution", "validators"), validatorAddress,
		), "",
	)
	if err != nil {
		return nil, err
	}
	defer rawRespBody.Close()

	var commissionResp ValidatorCommissionResp
	if err := jsoniter.NewDecoder(rawRespBody).Decode(&commissionResp); err != nil {
		return nil, err
	}

	totalCommission := coin.NewEmptyDecCoins()
	for _, commission := range commissionResp.Commissions.Commission {
		commissionCoin, coinErr := coin.NewDecCoinFromString(commission.Denom, commission.Amount)
		if coinErr != nil {
			return nil, fmt.Errorf("error parsing Coin from commission: %v", coinErr)
		}
		totalCommission = totalCommission.Add(commissionCoin)
	}
	return totalCommission, nil
}

func (client *HTTPClient) Delegation(
	delegator string, validator string,
) (*cosmosapp_interface.DelegationResponse, error) {
	resp := &DelegationResp{
		MaybePagination: &Pagination{
			MaybeNextKey: nil,
			Total:        "",
		},
	}
	for {
		url := fmt.Sprintf("%s/%s", client.url("staking", "delegations"), delegator)
		if resp.MaybePagination.MaybeNextKey != nil {
			url = fmt.Sprintf("%s?pagination.key=%s", url, *resp.MaybePagination.MaybeNextKey)
		}

		rawRespBody, statusCode, err := client.rawRequest(url)
		if err != nil {
			return nil, err
		}
		defer rawRespBody.Close()

		if decodeErr := jsoniter.NewDecoder(rawRespBody).Decode(&resp); decodeErr != nil {
			return nil, decodeErr
		}
		if resp.MaybeCode != nil && *resp.MaybeCode == 2 {
			return nil, cosmosapp_interface.ErrAccountNotFound
		}
		if statusCode != 200 {
			return nil, fmt.Errorf("error requesting Cosmos %s endpoint: %v", url, err)
		}

		for _, delegation := range resp.MaybeDelegationResponses {
			if delegation.Delegation.DelegatorAddress == delegator && delegation.Delegation.ValidatorAddress == validator {
				return &delegation, nil
			}
		}

		if resp.MaybePagination.MaybeNextKey == nil {
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
		return nil, fmt.Errorf("error requesting Cosmos %s endpoint: %v", url, err)
	}

	if rawResp.StatusCode != 200 {
		rawResp.Body.Close()
		return nil, fmt.Errorf("error requesting Cosmos %s endpoint: %s", method, rawResp.Status)
	}

	return rawResp.Body, nil
}

// rawRequest construct tendermint url and issues an HTTP request
// returns the http Body with any status code
func (client *HTTPClient) rawRequest(method string, queryString ...string) (io.ReadCloser, int, error) {
	var err error

	url := client.rpcUrl + "/" + method
	if len(queryString) > 0 {
		url += "?" + queryString[0]
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return nil, 0, fmt.Errorf("error creating HTTP request with context: %v", err)
	}
	// nolint:bodyclose
	rawResp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("error requesting Cosmos %s endpoint: %v", url, err)
	}

	return rawResp.Body, rawResp.StatusCode, nil
}

type Pagination struct {
	MaybeNextKey *string `json:"next_key"`
	Total        string  `json:"total"`
}
