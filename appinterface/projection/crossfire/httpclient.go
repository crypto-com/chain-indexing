package crossfire

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
)

type HTTPClient struct {
	httpClient *http.Client
	url     string
}

// NewHTTPClient returns a new HTTPClient for external request
func NewHTTPClient(participantsListURL string) *HTTPClient {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &HTTPClient{
		httpClient,
		participantsListURL,
	}
}

func (client *HTTPClient) Participants() (*[]Participant, error) {
	rawRespBody, err := client.request(client.url)
	if err != nil {
		return nil, err
	}
	defer rawRespBody.Close()

	var participants []Participant
	if err := jsoniter.NewDecoder(rawRespBody).Decode(&participants); err != nil {
		return nil, err
	}

	return &participants, nil
}

// request construct participants list url and issues an HTTP request
// returns the success http Body
func (client *HTTPClient) request(method string, queryString ...string) (io.ReadCloser, error) {
	var err error

	url := client.url
	if len(queryString) > 0 {
		url += "?" + queryString[0]
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating HTTP request with context: %v", err)
	}
	rawResp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error requesting url %s endpoint: %v", url, err)
	}

	if rawResp.StatusCode != 200 {
		rawResp.Body.Close()
		return nil, fmt.Errorf("error requesting url %s endpoint: %s", method, rawResp.Status)
	}

	return rawResp.Body, nil
}

type Participant struct {
	OperatorAddress string `json:"operatorAddress"`
	PrimaryAddress  string `json:"primaryAddress"`
	Moniker         string `json:"moniker"`
}
