package cosmosapp

import (
	"context"
	"fmt"
	"io"
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

func (client *HTTPClient) Validator(validatorAddress string) (*cosmosapp_interface.Validator, error) {
	rawRespBody, err := client.request(
		fmt.Sprintf("%s/%s", client.url("staking", "validators"), validatorAddress),
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
			url = fmt.Sprintf("%s?pagination.key=%s", url, resp.Pagination.MaybeNextKey)
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
