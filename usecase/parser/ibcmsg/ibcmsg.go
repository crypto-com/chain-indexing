package ibcmsg

import (
	"fmt"
	"time"

	"github.com/crypto-com/chain-indexing/usecase/parser/utils"

	"github.com/crypto-com/chain-indexing/usecase/model"

	"github.com/mitchellh/mapstructure"

	"github.com/crypto-com/chain-indexing/entity/command"
	"github.com/crypto-com/chain-indexing/internal/utctime"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/event"
	ibc_model "github.com/crypto-com/chain-indexing/usecase/model/ibc"
)

func ParseMsgCreateClient(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) []command.Command {
	clientStateType := msg["client_state"].(map[string]interface{})["@type"]
	if clientStateType != "/ibc.lightclients.tendermint.v1.ClientState" {
		// TODO: SoloMachine and Localhost LightClient
		panic(fmt.Sprintf("Unsupported Light Client Type: %s", clientStateType))
	}

	var clientState TendermintLightClientState
	clientStateDecoderConfig := &mapstructure.DecoderConfig{
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
		),
		Result: &clientState,
	}
	clientStateDecoder, clientStateDecoderErr := mapstructure.NewDecoder(clientStateDecoderConfig)
	if clientStateDecoderErr != nil {
		panic(fmt.Errorf("error creating client state decoder: %v", clientStateDecoderErr))
	}
	if err := clientStateDecoder.Decode(msg["client_state"]); err != nil {
		panic(fmt.Errorf("error decoding client state: %v", err))
	}

	var proofSpecs []ibc_model.TendermintLightClientProofSpec
	if clientState.ProofSpecs != nil {
		proofSpecs = make([]ibc_model.TendermintLightClientProofSpec, 0, len(clientState.ProofSpecs))
		for _, proofSpec := range clientState.ProofSpecs {
			var maybeLeafSpec *ibc_model.TendermintLightClientLeafSpec
			if proofSpec.MaybeLeafSpec != nil {
				maybeLeafSpec = &ibc_model.TendermintLightClientLeafSpec{
					Hash:         proofSpec.MaybeLeafSpec.Hash,
					PrehashKey:   proofSpec.MaybeLeafSpec.PrehashKey,
					PrehashValue: proofSpec.MaybeLeafSpec.PrehashValue,
					Length:       proofSpec.MaybeLeafSpec.Length,
					Prefix:       proofSpec.MaybeLeafSpec.Prefix,
				}
			}
			var maybeInnerSpec *ibc_model.TendermintLightClientInnerSpec
			if proofSpec.MaybeInnerSpec != nil {
				var childOrder []int64
				if proofSpec.MaybeInnerSpec.ChildOrder != nil {
					childOrder = append(childOrder, proofSpec.MaybeInnerSpec.ChildOrder...)
				}

				maybeInnerSpec = &ibc_model.TendermintLightClientInnerSpec{
					ChildOrder:      childOrder,
					ChildSize:       proofSpec.MaybeInnerSpec.ChildSize,
					MinPrefixLength: proofSpec.MaybeInnerSpec.MinPrefixLength,
					MaxPrefixLength: proofSpec.MaybeInnerSpec.MaxPrefixLength,
					EmptyChild:      proofSpec.MaybeInnerSpec.EmptyChild,
					Hash:            proofSpec.MaybeInnerSpec.Hash,
				}
			}

			proofSpecs = append(proofSpecs, ibc_model.TendermintLightClientProofSpec{
				MaybeLeafSpec:  maybeLeafSpec,
				MaybeInnerSpec: maybeInnerSpec,
				MaxDepth:       proofSpec.MaxDepth,
				MinDepth:       proofSpec.MinDepth,
			})
		}

	}

	var upgradePath []string
	if clientState.UpgradePath != nil {
		upgradePath = make([]string, 0, len(clientState.UpgradePath))
		upgradePath = append(upgradePath, clientState.UpgradePath...)
	}

	var consensusState TendermintLightClientConsensusState
	if err := mapstructure.Decode(msg["consensus_state"], &consensusState); err != nil {
		panic(fmt.Errorf("error decoding consensus state: %v", err))
	}

	timestamp, parseTimestampErr := utctime.Parse(time.RFC3339, consensusState.Timestamp)
	if parseTimestampErr != nil {
		panic(fmt.Errorf("error parsing consensus state timestamp: %v", parseTimestampErr))
	}

	log := utils.NewParsedTxsResultLog(&txsResult.Log[msgIndex])
	// When there is no reward withdrew, `transfer` event would not exist
	event := log.GetEventByType("create_client")
	if event == nil {
		panic("missing `create_client` event in TxsResult log")
	}
	params := ibc_model.MsgCreateClientParams{
		MaybeTendermintLightClient: &ibc_model.TendermintLightClient{
			TendermintClientState: ibc_model.TendermintLightClientState{
				Type:    clientState.Type,
				ChainID: clientState.ChainID,
				TrustLevel: ibc_model.TendermintLightClientTrustLevel{
					Numerator:   clientState.TrustLevel.Numerator,
					Denominator: clientState.TrustLevel.Denominator,
				},
				TrustingPeriod:  clientState.TrustingPeriod,
				UnbondingPeriod: clientState.UnbondingPeriod,
				MaxClockDrift:   clientState.MaxClockDrift,
				FrozenHeight: ibc_model.TendermintLightClientHeight{
					RevisionNumber: clientState.FrozenHeight.RevisionNumber,
					RevisionHeight: clientState.FrozenHeight.RevisionHeight,
				},
				LatestHeight: ibc_model.TendermintLightClientHeight{
					RevisionNumber: clientState.LatestHeight.RevisionNumber,
					RevisionHeight: clientState.LatestHeight.RevisionHeight,
				},
				ProofSpecs:                   proofSpecs,
				UpgradePath:                  upgradePath,
				AllowUpdateAfterExpiry:       clientState.AllowUpdateAfterExpiry,
				AllowUpdateAfterMisbehaviour: clientState.AllowUpdateAfterMisbehaviour,
			},
			TendermintLightClientConsensusState: ibc_model.TendermintLightClientConsensusState{
				Type:      consensusState.Type,
				Timestamp: timestamp,
				Root: ibc_model.TendermintLightClientRoot{
					Hash: consensusState.Root.Hash,
				},
				NextValidatorsHash: consensusState.NextValidatorsHash,
			},
		},
		Signer:     msg["signer"].(string),
		ClientID:   event.MustGetAttributeByKey("client_id"),
		ClientType: event.MustGetAttributeByKey("client_type"),
	}

	return []command.Command{command_usecase.NewCreateMsgIBCCreateClient(
		msgCommonParams,

		params,
	)}
}

func ParseMsgConnectionOpenInit(
	msgCommonParams event.MsgCommonParams,
	txsResult model.BlockResultsTxsResult,
	msgIndex int,
	msg map[string]interface{},
) []command.Command {
	var rawMessage ibc_model.RawMsgConnectionOpenInit
	decoder, decoderErr := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &rawMessage,
	})
	if decoderErr != nil {
		panic(fmt.Errorf("error creating mapstructure decoder: %v", decoderErr))
	}
	if err := decoder.Decode(msg); err != nil {
		panic(fmt.Errorf("error decoding message: %v", err))
	}

	log := utils.NewParsedTxsResultLog(&txsResult.Log[msgIndex])
	event := log.GetEventByType("connection_open_init")
	if event == nil {
		panic("missing `connection_open_init` event in TxsResult log")
	}
	params := ibc_model.MsgConnectionOpenInitParams{
		RawMsgConnectionOpenInit: rawMessage,

		ConnectionID: event.MustGetAttributeByKey("connection_id"),
	}

	return []command.Command{command_usecase.NewCreateMsgIBCConnectionOpenInit(
		msgCommonParams,

		params,
	)}
}

//func ParseMsgUpdateClient(
//	msgCommonParams event.MsgCommonParams,
//	txsResult model.BlockResultsTxsResult,
//	msgIndex int,
//	msg map[string]interface{},
//) []command.Command {
//	header := msg["header"].(map[string]interface{})
//	if headerx["@type"]
//    params := ibc_model.MsgUpdateClientParams{
//		MaybeTendermintLightClientUpdate: nil,
//		ClientID:                         "",
//		Signer:                           "",
//	}
//}
