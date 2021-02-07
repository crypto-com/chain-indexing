package tendermint

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/crypto-com/chain-indexing/usecase/model/genesis"

	usecase_model "github.com/crypto-com/chain-indexing/usecase/model"
)

type HTTPClient struct {
	httpClient       *http.Client
	tendermintRPCUrl string
}

// NewHTTPClient returns a new HTTPClient for tendermint request
func NewHTTPClient(tendermintRPCUrl string) *HTTPClient {
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	return &HTTPClient{
		httpClient,
		strings.TrimSuffix(tendermintRPCUrl, "/"),
	}
}

func (client *HTTPClient) Genesis() (*genesis.Genesis, error) {
	var err error

	rawRespBody, err := client.request("genesis")
	if err != nil {
		return nil, err
	}
	defer rawRespBody.Close()

	genesis, err := ParseGenesisResp(rawRespBody)
	if err != nil {
		return nil, err
	}

	return genesis, nil
}

// Block gets the block response with target height
func (client *HTTPClient) Block(height int64) (*usecase_model.Block, *usecase_model.RawBlock, error) {
	var err error

	rawRespBody, err := client.request("block", "height="+strconv.FormatInt(height, 10))
	if err != nil {
		return nil, nil, err
	}
	defer rawRespBody.Close()

	block, rawBlock, err := ParseBlockResp(rawRespBody)
	if err != nil {
		return nil, nil, err
	}

	return block, rawBlock, nil
}

func (client *HTTPClient) BlockResults(height int64) (*usecase_model.BlockResults, error) {
	var err error

	rawRespBody, err := client.request("block_results", "height="+strconv.FormatInt(height, 10))
	if err != nil {
		return nil, err
	}
	defer rawRespBody.Close()

	blockResults, err := ParseBlockResultsResp(rawRespBody)
	if err != nil {
		return nil, err
	}

	return blockResults, nil
}

// LatestBlockHeight gets the chain's latest block and return the height
func (client *HTTPClient) LatestBlockHeight() (int64, error) {
	var err error
	rawRespBody, err := client.request("block")
	if err != nil {
		return int64(0), fmt.Errorf("error getting /block: %v", err)
	}
	defer rawRespBody.Close()

	block, _, err := ParseBlockResp(rawRespBody)
	if err != nil {
		return int64(0), fmt.Errorf("error parsing /block response: %v", err)
	}

	return block.Height, nil
}

// request construct tendermint url and issues an HTTP request
// returns the success http Body
func (client *HTTPClient) request(method string, queryString ...string) (io.ReadCloser, error) {
	var err error

	url := client.tendermintRPCUrl + "/" + method
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

func (client *HTTPClient) Status() (*map[string]interface{}, error) {
	rawRespBody, err := client.request("status")
	if err != nil {
		return nil, err
	}
	defer rawRespBody.Close()

	body, _ := ioutil.ReadAll(rawRespBody)
	jsonMap := make(map[string]interface{})
	errread := json.Unmarshal([]byte(body), &jsonMap)
	if errread != nil {
		return nil, fmt.Errorf("error requesting Status : %v", errread)
	}

	return &jsonMap, nil
}

type GenesisResp struct {
	Jsonrpc string            `json:"jsonrpc"`
	ID      int64             `json:"id"`
	Result  GenesisRespResult `json:"result"`
}

type GenesisRespResult struct {
	Genesis genesis.Genesis `json:"genesis"`
}

type RawBlockResp struct {
	Jsonrpc string                 `json:"jsonrpc"`
	ID      int                    `json:"id"`
	Result  usecase_model.RawBlock `json:"result"`
}

type RawBlockResultsResp struct {
	Jsonrpc string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  RawBlockResults `json:"result"`
}
