package tendermint

import (
	"encoding/base64"
	"fmt"
	"github.com/crypto-com/chainindex/usecase/model"
	jsoniter "github.com/json-iterator/go"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"io"
	"strconv"
)

// Block related parsing functions
func parseBlockResp(rawRespReader io.Reader) (*model.Block, *model.RawBlock, error) {
	var err error

	var resp model.RawBlock
	if err = jsoniter.NewDecoder(rawRespReader).Decode(&resp); err != nil {
		return nil, nil, fmt.Errorf("error decoding Tendermint block response: %v", err)
	}

	height, err := strconv.ParseInt(resp.Result.Block.Header.Height, 10, 64)
	if err != nil {
		return nil, nil, fmt.Errorf("error converting block height to unsigned integer: %v", err)
	}

	return &model.Block{
		Height:          height,
		Hash:            resp.Result.BlockID.Hash,
		Time:            resp.Result.Block.Header.Time,
		AppHash:         resp.Result.Block.Header.AppHash,
		ProposerAddress: resp.Result.Block.Header.ProposerAddress,
		Txs:             resp.Result.Block.Data.Txs,
		Signatures:      parseBlockSignatures(resp.Result.Block.LastCommit.Signatures),
	}, &resp, nil
}

func parseBlockSignatures(rawSignatures []model.RawBlockSignature) []model.BlockSignature {
	if rawSignatures == nil {
		return nil
	}

	signatures := make([]model.BlockSignature, 0, len(rawSignatures))
	for _, rawSignature := range rawSignatures {
		if rawSignature.Signature == nil {
			continue
		}
		signatures = append(signatures, model.BlockSignature{
			BlockIdFlag:      rawSignature.BlockIDFlag,
			ValidatorAddress: rawSignature.ValidatorAddress,
			Timestamp:        rawSignature.Timestamp,
			Signature:        *rawSignature.Signature,
		})
	}

	return signatures
}

// BlockResults related parsing functions
func parseBlockResultsResp(rawRespReader io.Reader) (*model.BlockResults, error) {
	var err error

	var resp model.RawBlockResults
	if err = jsoniter.NewDecoder(rawRespReader).Decode(&resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling Tendermint block_results response: %v", err)
	}

	var txsResults [][]model.BlockResultsEvent
	if resp.Result.TxsEvents != nil {
		txsResults = parseBlockResultsTxsEvents(resp.Result.TxsEvents)
	}

	var beginBlockEvents []model.BlockResultsEvent
	if resp.Result.BeginBlockEvents != nil {
		beginBlockEvents = parseBlockResultsEvent(resp.Result.BeginBlockEvents)
	}

	height, err := strconv.ParseUint(resp.Result.Height, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error converting block height to unsigned integer: %v", err)
	}
	return &model.BlockResults{
		Height:           int64(height),
		TxsEvents:        txsResults,
		BeginBlockEvents: beginBlockEvents,
		ValidatorUpdates: parseBlockResultsValidatorUpdates(resp.Result.ValidatorUpdates),
	}, nil
}

func parseBlockResultsTxsEvents(rawResults []model.RawBlockResultsTxResult) [][]model.BlockResultsEvent {
	results := make([][]model.BlockResultsEvent, 0, len(rawResults))
	for _, rawResult := range rawResults {
		results = append(results, parseBlockResultsEvent(rawResult.Events))
	}

	return results
}

func parseBlockResultsEvent(rawEvents []model.RawBlockResultsEvent) []model.BlockResultsEvent {
	if rawEvents == nil {
		return nil
	}

	events := make([]model.BlockResultsEvent, 0, len(rawEvents))
	for _, rawEvent := range rawEvents {
		attributes := make([]model.BlockResultsEventAttribute, 0, len(rawEvent.Attributes))
		for _, rawAttribute := range rawEvent.Attributes {
			attributes = append(attributes, model.BlockResultsEventAttribute{
				Key:   rawAttribute.Key,
				Value: rawAttribute.Value,
			})
		}
		events = append(events, model.BlockResultsEvent{
			Type:       rawEvent.Type,
			Attributes: attributes,
		})
	}

	return events
}

func parseBlockResultsValidatorUpdates(rawUpdates []model.RawBlockResultsValidator) []model.BlockResultsValidator {
	if rawUpdates == nil {
		return nil
	}

	updates := make([]model.BlockResultsValidator, 0, len(rawUpdates))
	for _, rawUpdate := range rawUpdates {
		updates = append(updates, model.BlockResultsValidator{
			PubKey: model.BlockResultsValidatorPubKey{
				Type:    rawUpdate.PubKey.Sum.Type,
				PubKey:  rawUpdate.PubKey.Sum.Value.Ed25519,
				Address: AddressFromPubKey(rawUpdate.PubKey.Sum.Value.Ed25519),
			},
			Power: rawUpdate.Power,
		})
	}

	return updates
}

// Tendermint type convert function converts public key to address
func AddressFromPubKey(base64PubKey string) string {
	var key ed25519.PubKeyEd25519
	data, _ := base64.StdEncoding.DecodeString(base64PubKey)
	copy(key[:], data[:])

	return key.Address().String()
}
