package tendermint

import (
	"encoding/base64"
	"fmt"
	"io"
	"math/big"
	"strconv"
	"strings"

	tendermint_interface "github.com/crypto-com/chain-indexing/appinterface/tendermint"

	"github.com/crypto-com/chain-indexing/usecase/model/genesis"

	"github.com/crypto-com/chain-indexing/external/tmcosmosutils"

	"github.com/crypto-com/chain-indexing/usecase/model"
	jsoniter "github.com/json-iterator/go"
)

const UNICODE_NULL_VALUE = "\u0000"

// Block related parsing functions
func ParseGenesisResp(rawRespReader io.Reader, strictParsing bool) (*genesis.Genesis, error) {
	var genesisResp GenesisResp
	jsonDecoder := jsoniter.NewDecoder(rawRespReader)
	if strictParsing {
		jsonDecoder.DisallowUnknownFields()
	}
	if err := jsonDecoder.Decode(&genesisResp); err != nil {
		return nil, fmt.Errorf("error decoding Tendermint genesis response: %v", err)
	}

	return &genesisResp.Result.Genesis, nil
}

func ParseGenesisChunkedResp(rawRespReader io.Reader, strictParsing bool) (*GenesisChunkedResp, error) {
	var genesisChunkedResp GenesisChunkedResp

	jsonDecoder := jsoniter.NewDecoder(rawRespReader)
	if strictParsing {
		jsonDecoder.DisallowUnknownFields()
	}
	if err := jsonDecoder.Decode(&genesisChunkedResp); err != nil {
		return nil, fmt.Errorf("error decoding Tendermint genesis response: %v", err)
	}

	genesisDataByte, err := base64.StdEncoding.DecodeString(genesisChunkedResp.Result.Data)
	genesisChunkedResp.Result.Data = string(genesisDataByte)

	if err != nil {
		return nil, fmt.Errorf("error decoding Tendermint genesis chunked from base64 to byte: %v", err)
	}

	return &genesisChunkedResp, nil
}

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
		Evidences:       resp.Result.Block.Evidence.Evidence,
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
func ParseBlockResultsResp(rawRespReader io.Reader, eventAttributeDecoder tendermint_interface.BlockResultEventAttributeDecoder) (*model.BlockResults, error) {
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

	var beginBlockEvents []model.BlockResultsEvent
	var endBlockEvents []model.BlockResultsEvent

	if len(rawBlockResults.FinalizeBlockEvents) > 0 {
		beginBlockEvents, endBlockEvents = ParseEthermintBlockResultFinalizeBlockEvents(rawBlockResults.FinalizeBlockEvents, eventAttributeDecoder)
	} else {
		beginBlockEvents = parseBlockResultsEvents(rawBlockResults.BeginBlockEvents, eventAttributeDecoder)
		endBlockEvents = parseBlockResultsEvents(rawBlockResults.EndBlockEvents, eventAttributeDecoder)
	}

	txsResults := parseBlockResultsTxsResults(rawBlockResults.TxsResults, eventAttributeDecoder)
	return &model.BlockResults{
		//nolint:gosec
		Height:                int64(height),
		TxsResults:            txsResults,
		BeginBlockEvents:      beginBlockEvents,
		EndBlockEvents:        endBlockEvents,
		ValidatorUpdates:      parseBlockResultsValidatorUpdates(rawBlockResults.ValidatorUpdates),
		ConsensusParamUpdates: parseBlockResultsConsensusParamsUpdates(rawBlockResults.ConsensusParamUpdates),
	}, nil
}

func parseBlockResultsTxsResults(rawTxsResults []RawBlockResultsTxsResult, eventAttributeDecoder tendermint_interface.BlockResultEventAttributeDecoder) []model.BlockResultsTxsResult {
	txsResults := make([]model.BlockResultsTxsResult, 0, len(rawTxsResults))
	for _, rawTxsResult := range rawTxsResults {
		events := parseBlockResultsEvents(rawTxsResult.Events, eventAttributeDecoder)

		txsResults = append(txsResults, model.BlockResultsTxsResult{
			Code:      rawTxsResult.Code,
			Data:      mustBase64Decode(rawTxsResult.Data),
			Log:       parseBlockResultsTxsResultLog(rawTxsResult.Log),
			RawLog:    strings.ReplaceAll(rawTxsResult.Log, UNICODE_NULL_VALUE, ""),
			Info:      rawTxsResult.Info,
			GasWanted: rawTxsResult.GasWanted,
			GasUsed:   rawTxsResult.GasUsed,
			Events:    events,
			Codespace: rawTxsResult.Codespace,
		})
	}

	return txsResults
}

func parseBlockResultsTxsResultLog(rawLog string) []model.BlockResultsTxsResultLog {
	rawLog = strings.ReplaceAll(rawLog, UNICODE_NULL_VALUE, "")

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

func parseBlockResultsEvents(rawEvents []RawBlockResultsEvent, eventAttributeDecoder tendermint_interface.BlockResultEventAttributeDecoder) []model.BlockResultsEvent {
	if rawEvents == nil {
		return []model.BlockResultsEvent{}
	}

	events := make([]model.BlockResultsEvent, 0, len(rawEvents))
	for _, rawEvent := range rawEvents {
		attributes := make([]model.BlockResultsEventAttribute, 0, len(rawEvent.Attributes))
		for _, rawAttribute := range rawEvent.Attributes {
			key, err := eventAttributeDecoder.DecodeKey(rawAttribute.Key)
			if err != nil {
				panic(fmt.Sprintf("error block result event %s attribute key (%s): %v", rawEvent.Type, rawAttribute.Key, err))
			}
			value, err := eventAttributeDecoder.DecodeValue(rawAttribute.Value)
			if err != nil {
				panic(fmt.Sprintf("error block result event %s attribute key (%s): %v", rawEvent.Type, rawAttribute.Key, err))
			}
			attributes = append(attributes, model.BlockResultsEventAttribute{
				Key:   key,
				Value: value,
				Index: rawAttribute.Index,
			})
		}
		events = append(events, model.BlockResultsEvent{
			Type:       rawEvent.Type,
			Attributes: attributes,
		})
	}

	return events
}

func ParseEthermintBlockResultFinalizeBlockEvents(rawEvents []RawBlockResultsEvent, eventAttributeDecoder tendermint_interface.BlockResultEventAttributeDecoder) ([]model.BlockResultsEvent, []model.BlockResultsEvent) {
	if rawEvents == nil {
		return []model.BlockResultsEvent{}, []model.BlockResultsEvent{}
	}

	beginBlockEvents := make([]model.BlockResultsEvent, 0, len(rawEvents))
	endBlockEvents := make([]model.BlockResultsEvent, 0, len(rawEvents))
	for _, rawEvent := range rawEvents {
		mode := ""
		attributes := make([]model.BlockResultsEventAttribute, 0, len(rawEvent.Attributes))
		for _, rawAttribute := range rawEvent.Attributes {
			key, err := eventAttributeDecoder.DecodeKey(rawAttribute.Key)
			if err != nil {
				key = rawAttribute.Key
			}
			value, err := eventAttributeDecoder.DecodeValue(rawAttribute.Value)
			if err != nil {
				value = rawAttribute.Value
			}

			if rawAttribute.Key == "mode" {
				mode = rawAttribute.Value
			}

			attributes = append(attributes, model.BlockResultsEventAttribute{
				Key:   key,
				Value: value,
				Index: rawAttribute.Index,
			})
		}

		switch mode {
		case "BeginBlock":
			{
				beginBlockEvents = append(beginBlockEvents, model.BlockResultsEvent{
					Type:       rawEvent.Type,
					Attributes: attributes,
				})
			}
		case "EndBlock":
			{
				endBlockEvents = append(endBlockEvents, model.BlockResultsEvent{
					Type:       rawEvent.Type,
					Attributes: attributes,
				})
			}
		}
	}

	return beginBlockEvents, endBlockEvents
}

func parseBlockResultsValidatorUpdates(rawUpdates []RawBlockResultsValidatorUpdate) []model.BlockResultsValidatorUpdate {
	if rawUpdates == nil {
		return []model.BlockResultsValidatorUpdate{}
	}

	updates := make([]model.BlockResultsValidatorUpdate, 0, len(rawUpdates))
	for _, rawUpdate := range rawUpdates {
		var power *big.Int
		if rawUpdate.MaybePower != "" {
			var ok bool
			power, ok = new(big.Int).SetString(rawUpdate.MaybePower, 10)
			if !ok {
				panic(fmt.Sprintf("invalid block_results power: %s", power))
			}
		}

		ed25519PubKey, err := base64.StdEncoding.DecodeString(rawUpdate.PubKey.Sum.Value.Ed25519)
		if err != nil {
			panic(fmt.Sprintf("invalid tendermint public key: %v", err))
		}
		updates = append(updates, model.BlockResultsValidatorUpdate{
			Pubkey: model.BlockResultsValidatorPubKey{
				Type:   rawUpdate.PubKey.Sum.Type,
				Pubkey: rawUpdate.PubKey.Sum.Value.Ed25519,
			},
			Address:    tmcosmosutils.TmAddressFromTmPubKey(ed25519PubKey),
			MaybePower: power,
		})
	}

	return updates
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
			MaxBytes:        rawUpdates.Evidence.MaxBytes,
		},
		Validator: model.BlockResultsConsensusParamsUpdatesValidator{
			PubKeyTypes: validatorPubKeyTypes,
		},
		Version: model.BlockResultsConsensusParamsUpdatesVersion{
			AppVersion: rawUpdates.Version.AppVersion,
		},
	}
}
