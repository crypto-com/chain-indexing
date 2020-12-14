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
	rawRespBody, err := client.request(
		fmt.Sprintf("%s/%s", client.url("auth", "accounts"), accountAddress), "",
	)
	if err != nil {
		return nil, err
	}
	defer rawRespBody.Close()

	outputbytes, readerr := ioutil.ReadAll(rawRespBody)
	if readerr != nil {
		return nil, fmt.Errorf("GetAccountInfo error ioutil.ReadAll %v", readerr)
	}

	var myjson map[string]interface{}

	if err := json.Unmarshal(outputbytes, &myjson); err != nil {
		return nil, fmt.Errorf("GetAccountInfo error json.Unmarshal %v", err)
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
		thisAccountNumber = thisBaseAccount["account_number"].(string)
		thisSequenceNumber = thisBaseAccount["sequence"].(string)

	}

	var accountResp AccountResp
	accountResp.Account.AccountType = thisAccountType
	accountResp.Account.AccountAddress = thisAddress
	accountResp.Account.Pubkey = thisPubkey
	accountResp.Account.AccountNumber = thisAccountNumber
	accountResp.Account.SequenceNumber = thisSequenceNumber

	return &accountResp.Account, nil
}

func (client *HTTPClient) Balance(targetAddress string, targetDenom string) (*cosmosapp_interface.AccountBalance, error) {
	rawRespBody, err := client.request(
		fmt.Sprintf("%s/%s/%s", client.url("bank", "balances"), targetAddress, targetDenom), "",
	)
	if err != nil {
		return nil, err
	}
	defer rawRespBody.Close()

	outputbytes, readerr := ioutil.ReadAll(rawRespBody)
	if readerr != nil {
		return nil, fmt.Errorf("GetAccountInfo error ioutil.ReadAll %v", readerr)
	}
	var thisJson map[string]interface{}

	if err := json.Unmarshal(outputbytes, &thisJson); err != nil {
		return nil, err
	}

	thisBalance := thisJson["balance"].(map[string]interface{})
	thisAmount := thisBalance["amount"].(string)
	thisDenom := thisBalance["denom"].(string)

	var accountBalanceResp AccountBalanceResp
	accountBalanceResp.AccountBalace.AccountAmount = thisAmount
	accountBalanceResp.AccountBalace.AccountDenom = thisDenom

	return &accountBalanceResp.AccountBalace, nil

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

type AccountBalanceResp struct {
	AccountBalace cosmosapp_interface.AccountBalance
}
