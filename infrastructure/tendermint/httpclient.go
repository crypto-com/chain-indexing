package tendermint

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/crypto-com/chain-indexing/appinterface/tendermint"

	"github.com/crypto-com/chain-indexing/usecase/model/genesis"

	usecase_model "github.com/crypto-com/chain-indexing/usecase/model"
)

var _ tendermint.Client = &HTTPClient{}

type HTTPClient struct {
	httpClient           *http.Client
	tendermintRPCUrl     string
	strictGenesisParsing bool
}

// NewHTTPClient returns a new HTTPClient for tendermint request
func NewHTTPClient(tendermintRPCUrl string, strictGenesisParsing bool) *HTTPClient {
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	return &HTTPClient{
		httpClient,
		strings.TrimSuffix(tendermintRPCUrl, "/"),
		strictGenesisParsing,
	}
}

// NewHTTPClient returns a new HTTPClient for tendermint request
func NewInsecureHTTPClient(tendermintRPCUrl string, strictGenesisParsing bool) *HTTPClient {
	// nolint:gosec
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{
		Timeout:   30 * time.Second,
		Transport: transport,
	}

	return &HTTPClient{
		httpClient,
		strings.TrimSuffix(tendermintRPCUrl, "/"),
		strictGenesisParsing,
	}
}

func (client *HTTPClient) Genesis() (*genesis.Genesis, error) {
	var err error

	rawRespBody, err := client.request("genesis")
	if err != nil {
		return nil, err
	}
	defer rawRespBody.Close()

	genesis, err := ParseGenesisResp(rawRespBody, client.strictGenesisParsing)
	if err != nil {
		return nil, err
	}

	return genesis, nil
}

func (client *HTTPClient) GenesisChunked() (*genesis.Genesis, error) {
	var err error

	// get the total number of genesis chunks
	firstRawRespBody, err := client.request("genesis_chunked", "chunk="+strconv.FormatInt(0, 10))
	if err != nil {
		return nil, err
	}
	defer firstRawRespBody.Close()

	firstGenesisData, err := ParseGenesisChunkedResp(firstRawRespBody, client.strictGenesisParsing)
	if err != nil {
		return nil, err
	}

	total, err := strconv.Atoi(firstGenesisData.Result.Total)
	if err != nil {
		return nil, err
	}
	// loop through the genesis chunks
	decoded := make([]string, 0, total)
	decoded = append(decoded, firstGenesisData.Result.Data)

	for i := 1; i < total; i++ {
		rawRespBody, rawRespBodyErr := client.request("genesis_chunked", "chunk="+strconv.FormatInt(int64(i), 10))
		if rawRespBodyErr != nil {
			return nil, rawRespBodyErr
		}
		defer rawRespBody.Close()

		genesisData, genesisDataErr := ParseGenesisChunkedResp(rawRespBody, client.strictGenesisParsing)
		if genesisDataErr != nil {
			return nil, genesisDataErr
		}
		decoded = append(decoded, genesisData.Result.Data)
	}

	// combine the genesis chunks
	genesisDataByte := []byte(strings.Join(decoded, ""))

	var genesis genesis.Genesis
	genesisErr := json.Unmarshal(genesisDataByte, &genesis)
	if genesisErr != nil {
		return nil, fmt.Errorf("error decoding Tendermint genesis chunked from string to json: %v", err)
	}

	return &genesis, nil
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
		defer rawResp.Body.Close()

		body, errReadBody := ioutil.ReadAll(rawResp.Body)
		if errReadBody != nil {
			return nil, fmt.Errorf("error reading Body : %v", errReadBody)
		}

		bodyJsonMap := make(map[string]interface{})
		errRead := json.Unmarshal([]byte(body), &bodyJsonMap)
		if errRead != nil {
			return nil, fmt.Errorf("error unmarshalling Body : %v", errRead)
		}

		return nil, fmt.Errorf("error requesting Tendermint %s %s endpoint: %s Body: %v Header: %v", method, url, rawResp.Status, bodyJsonMap, rawResp.Header)
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
	errRead := json.Unmarshal([]byte(body), &jsonMap)
	if errRead != nil {
		return nil, fmt.Errorf("error requesting Status : %v", errRead)
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

type GenesisChunkedResp struct {
	Jsonrpc string                   `json:"jsonrpc"`
	ID      int64                    `json:"id"`
	Result  GenesisChunkedRespResult `json:"result"`
}

type GenesisChunkedRespResult struct {
	Chunk string `json:"chunk"`
	Total string `json:"total"`
	Data  string `json:"data"`
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
