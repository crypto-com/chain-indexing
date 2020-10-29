package tendermint

import (
	"fmt"
	"github.com/crypto-com/chainindex/appinterface/tendermint/types"
	jsoniter "github.com/json-iterator/go"
	"io"
	"net/http"
	"strconv"
	"time"
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
func (client *HTTPClient) Block(height uint64) (*types.Block, error) {
	var err error

	rawRespBody, err := client.request("block", "height="+strconv.FormatUint(height, 10))
	if err != nil {
		return nil, err
	}
	defer rawRespBody.Close()

	block, err := client.parseBlockResp(rawRespBody)
	if err != nil {
		return nil, err
	}

	return block, nil
}

// parseBlockSignatures parses the BlockResp into Block type
func (client *HTTPClient) parseBlockResp(rawRespReader io.Reader) (*types.Block, error) {
	var err error

	var resp types.BlockResp
	if err = jsoniter.NewDecoder(rawRespReader).Decode(&resp); err != nil {
		return nil, fmt.Errorf("error decoding Tendermint block response: %v", err)
	}

	height, err := strconv.ParseUint(resp.Result.Block.Header.Height, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error converting block height to unsigned integer: %v", err)
	}

	return &types.Block{
		Height:          height,
		Hash:            resp.Result.BlockID.Hash,
		Time:            resp.Result.Block.Header.Time,
		AppHash:         resp.Result.Block.Header.AppHash,
		ProposerAddress: resp.Result.Block.Header.ProposerAddress,
		Txs:             resp.Result.Block.Data.Txs,
		Signatures:      client.parseBlockSignatures(resp.Result.Block.LastCommit.Signatures),
	}, nil
}

// parseBlockSignatures parses the rawSignatures in JSON response into BlockSignature type
func (client *HTTPClient) parseBlockSignatures(rawSignatures []types.RawBlockSignature) []types.BlockSignature {
	if rawSignatures == nil {
		return nil
	}

	signatures := make([]types.BlockSignature, 0, len(rawSignatures))
	for _, rawSignature := range rawSignatures {
		if rawSignature.Signature == nil {
			continue
		}
		signatures = append(signatures, types.BlockSignature{
			BlockIDFlag:      rawSignature.BlockIDFlag,
			ValidatorAddress: rawSignature.ValidatorAddress,
			Timestamp:        rawSignature.Timestamp,
			Signature:        *rawSignature.Signature,
		})
	}

	return signatures
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
