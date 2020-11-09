package tendermint

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/crypto-com/chainindex/usecase/model"
)

type HTTPClient struct {
	httpClient       *http.Client
	tendermintRPCUrl string
}

// NewHTTPClient returns a new HTTPClient for tendermint request
func NewHTTPClient(tendermintRPCUrl string) *HTTPClient {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &HTTPClient{
		httpClient,
		tendermintRPCUrl,
	}
}

// Block gets the block response with target height
func (client *HTTPClient) Block(height int64) (*model.Block, *model.RawBlock, error) {
	var err error

	rawRespBody, err := client.request("block", "height="+strconv.FormatInt(height, 10))
	if err != nil {
		return nil, nil, err
	}
	defer rawRespBody.Close()

	block, rawBlock, err := parseBlockResp(rawRespBody)
	if err != nil {
		return nil, nil, err
	}

	return block, rawBlock, nil
}

func (client *HTTPClient) BlockResults(height int64) (*model.BlockResults, error) {
	var err error

	rawRespBody, err := client.request("block_results", "height="+strconv.FormatInt(height, 10))
	if err != nil {
		return nil, err
	}
	defer rawRespBody.Close()

	blockResults, err := parseBlockResultsResp(rawRespBody)
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

	block, _, err := parseBlockResp(rawRespBody)
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
	rawResp, err := client.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error requesting Tendermint %s endpoint: %v", url, err)
	}

	if rawResp.StatusCode != 200 {
		rawResp.Body.Close()
		return nil, fmt.Errorf("error requesting Tendermint %s endpoint: %s", method, rawResp.Status)
	}

	return rawResp.Body, nil
}
