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

	outputBytes, readErr := ioutil.ReadAll(rawRespBody)
	if readErr != nil {
		return nil, fmt.Errorf("error reading account info raw response: %v", readErr)
	}

	var rawResp map[string]interface{}

	if err := json.Unmarshal(outputBytes, &rawResp); err != nil {
		return nil, fmt.Errorf("error unmarshalling account info raw response: %v", err)
	}

	var accountType string
	var address string
	var pubKey string
	var accountNumber string
	var sequenceNumber string

	accountKVPair := rawResp["account"].(map[string]interface{})
	accountTypeMeta := accountKVPair["@type"].(string)
	accountName, hasAccountName := accountKVPair["name"].(string)
	if !hasAccountName {
		accountName = ""
	}
	accountType = fmt.Sprintf("%s %s", accountTypeMeta, accountName)
	baseAccount, isBaseAccount := accountKVPair["base_account"].(map[string]interface{})

	if !isBaseAccount {
		// normal account
		address = accountKVPair["address"].(string)
		// FIXME: account may not have pubkey (genesis account?)
		pubKeyKVPair, ok := accountKVPair["pub_key"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf(
				"error asserting account %s public key: %v", accountAddress, accountKVPair["pub_key"],
			)
		}
		pubKey = pubKeyKVPair["key"].(string)
		accountNumber = accountKVPair["account_number"].(string)
		sequenceNumber = accountKVPair["sequence"].(string)

	} else {
		// module account
		address = baseAccount["address"].(string)
		accountNumber = baseAccount["account_number"].(string)
		sequenceNumber = baseAccount["sequence"].(string)

	}

	var accountResp AccountResp
	accountResp.Account.AccountType = accountType
	accountResp.Account.AccountAddress = address
	accountResp.Account.Pubkey = pubKey
	accountResp.Account.AccountNumber = accountNumber
	accountResp.Account.SequenceNumber = sequenceNumber

	return &accountResp.Account, nil
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
	Account cosmosapp_interface.Account
}

type BankBalancesResp struct {
	BankBalanceResponses []cosmosapp_interface.BankBalance `json:"balances"`
	Pagination           cosmosapp_interface.Pagination    `json:"pagination"`
}
