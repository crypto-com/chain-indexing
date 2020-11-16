package tendermint

import (
	"encoding/base64"
	"fmt"
	"io"
	"math/big"
	"strconv"
	"strings"

	"github.com/crypto-com/chainindex/usecase/model"
	jsoniter "github.com/json-iterator/go"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

// Block related parsing functions
func ParseBlockResp(rawRespReader io.Reader) (*model.Block, *model.RawBlock, error) {
	var err error

	var resp RawBlockResp
	jsonDecoder := jsoniter.NewDecoder(rawRespReader)
	jsonDecoder.DisallowUnknownFields()
	if err = jsonDecoder.Decode(&resp); err != nil {
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
	}, &resp.Result, nil
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

// RawBlockResults related parsing functions
func ParseBlockResultsResp(rawRespReader io.Reader) (*model.BlockResults, error) {
	var err error

	var resp RawBlockResultsResp
	jsonDecoder := jsoniter.NewDecoder(rawRespReader)
	jsonDecoder.DisallowUnknownFields()
	if err = jsonDecoder.Decode(&resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling Tendermint block_results response: %v", err)
	}

	rawBlockResults := resp.Result

	height, err := strconv.ParseUint(resp.Result.Height, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error converting block height to unsigned integer: %v", err)
	}

	txsResults, err := parseBlockResultsTxsResults(rawBlockResults.TxsResults)
	if err != nil {
		return nil, fmt.Errorf("error parsing TxsResults: %v", err)
	}
	return &model.BlockResults{
		Height:                int64(height),
		TxsResults:            txsResults,
		BeginBlockEvents:      parseBlockResultsEvents(rawBlockResults.BeginBlockEvents),
		EndBlockEvents:        parseBlockResultsEvents(rawBlockResults.EndBlockEvents),
		ValidatorUpdates:      parseBlockResultsValidatorUpdates(rawBlockResults.ValidatorUpdates),
		ConsensusParamUpdates: parseBlockResultsConsensusParamsUpdates(rawBlockResults.ConsensusParamUpdates),
	}, nil
}

func parseBlockResultsTxsResults(rawTxsResults []RawBlockResultsTxsResult) ([]model.BlockResultsTxsResult, error) {
	txsResults := make([]model.BlockResultsTxsResult, 0, len(rawTxsResults))
	for _, rawTxsResult := range rawTxsResults {
		events := parseBlockResultsEvents(rawTxsResult.Events)

		txsResults = append(txsResults, model.BlockResultsTxsResult{
			Code:      rawTxsResult.Code,
			Data:      mustBase64Decode(rawTxsResult.Data),
			Log:       parseBlockResultsTxsResultLog(rawTxsResult.Log),
			RawLog:    rawTxsResult.Log,
			Info:      rawTxsResult.Info,
			GasWanted: rawTxsResult.GasWanted,
			GasUsed:   rawTxsResult.GasUsed,
			Events:    events,
			Codespace: rawTxsResult.Codespace,
		})
	}

	return txsResults, nil
}

func parseBlockResultsTxsResultLog(rawLog string) []model.BlockResultsTxsResultLog {
	jsonDecoder := jsoniter.NewDecoder(strings.NewReader(rawLog))
	jsonDecoder.DisallowUnknownFields()
	var decodedRawLogs []RawBlockResultsTxsResultLog
	if err := jsonDecoder.Decode(&decodedRawLogs); err != nil {
		// Ignore the JSON decode error because the log could be a string on error
		return []model.BlockResultsTxsResultLog{}
	}

	logs := make([]model.BlockResultsTxsResultLog, 0, len(decodedRawLogs))
	for _, rawLog := range decodedRawLogs {
		var log model.BlockResultsTxsResultLog
		if rawLog.MaybeMsgIndex == nil {
			log.MsgIndex = 0
		} else {
			log.MsgIndex = *rawLog.MaybeMsgIndex
		}

		log.Events = parseBlockResultsTxsResultLogEvents(rawLog.Events)
		logs = append(logs, log)
	}

	return logs
}

func parseBlockResultsTxsResultLogEvents(rawEvents []RawBlockResultsEvent) []model.BlockResultsEvent {
	if rawEvents == nil {
		return []model.BlockResultsEvent{}
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

func parseBlockResultsEvents(rawEvents []RawBlockResultsEvent) []model.BlockResultsEvent {
	if rawEvents == nil {
		return []model.BlockResultsEvent{}
	}

	events := make([]model.BlockResultsEvent, 0, len(rawEvents))
	for _, rawEvent := range rawEvents {
		attributes := make([]model.BlockResultsEventAttribute, 0, len(rawEvent.Attributes))
		for _, rawAttribute := range rawEvent.Attributes {
			attributes = append(attributes, model.BlockResultsEventAttribute{
				Key:   mustBase64Decode(rawAttribute.Key),
				Value: mustBase64Decode(rawAttribute.Value),
			})
		}
		events = append(events, model.BlockResultsEvent{
			Type:       rawEvent.Type,
			Attributes: attributes,
		})
	}

	return events
}

func parseBlockResultsValidatorUpdates(rawUpdates []RawBlockResultsValidatorUpdate) []model.BlockResultsValidatorUpdate {
	if rawUpdates == nil {
		return []model.BlockResultsValidatorUpdate{}
	}

	updates := make([]model.BlockResultsValidatorUpdate, 0, len(rawUpdates))
	for _, rawUpdate := range rawUpdates {
		var power *big.Int
		if rawUpdate.MaybePower != "" {
			power, ok := new(big.Int).SetString(rawUpdate.MaybePower, 10)
			if !ok {
				panic(fmt.Sprintf("invalid block_results power: %s", power))
			}
		}
		updates = append(updates, model.BlockResultsValidatorUpdate{
			PubKey: model.BlockResultsValidatorPubKey{
				Type:    rawUpdate.PubKey.Sum.Type,
				PubKey:  rawUpdate.PubKey.Sum.Value.Ed25519,
				Address: AddressFromPubKey(rawUpdate.PubKey.Sum.Value.Ed25519),
			},
			MaybePower: power,
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

func mustBase64Decode(s string) string {
	decoded, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(fmt.Sprintf("error decoding block_results `%s`: %v", s, err))
	}
	return string(decoded)
}

func parseBlockResultsConsensusParamsUpdates(rawUpdates RawBlockResultsConsensusParamUpdates) model.BlockResultsConsensusParamUpdates {
	var validatorPubKeyTypes []string
	if rawUpdates.Validator.PubKeyTypes == nil {
		validatorPubKeyTypes = []string{}
	} else {
		validatorPubKeyTypes = make([]string, 0, len(rawUpdates.Validator.PubKeyTypes))
		validatorPubKeyTypes = append(validatorPubKeyTypes, rawUpdates.Validator.PubKeyTypes...)
	}
	return model.BlockResultsConsensusParamUpdates{
		Block: model.BlockResultsConsensusParamUpdatesBlock{
			MaxBytes: rawUpdates.Block.MaxBytes,
			MaxGas:   rawUpdates.Block.MaxGas,
		},
		Evidence: model.BlockResultsConsensusParamUpdatesEvidence{
			MaxAgeNumBlocks: rawUpdates.Evidence.MaxAgeNumBlocks,
			MaxAgeDuration:  rawUpdates.Evidence.MaxAgeDuration,
		},
		Validator: model.BlockResultsConsensusParamsUpdatesValidator{
			PubKeyTypes: validatorPubKeyTypes,
		},
	}
}
