package parser

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	cosmosapp_interface "github.com/crypto-com/chain-indexing/appinterface/cosmosapp"
	applogger "github.com/crypto-com/chain-indexing/external/logger"
	"github.com/crypto-com/chain-indexing/usecase/parser/utils"

	"github.com/crypto-com/chain-indexing/entity/command"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/model"
	jsoniter "github.com/json-iterator/go"
)

func ParseTransactionCommands(
	logger applogger.Logger,
	txs []model.Tx,
	cosmosClient cosmosapp_interface.Client,
	blockResults *model.BlockResults,
	accountAddressPrefix string,
	possibleSignerAddresses []string,
) ([]command.Command, error) {
	blockHeight := blockResults.Height
	cmds := make([]command.Command, 0, len(blockResults.TxsResults))
	for i, tx := range txs {
		txHash := tx.TxResponse.TxHash
		txsResult := blockResults.TxsResults[i]

		var log string
		if len(txsResult.Log) == 0 {
			// cater for failed transaction
			log = txsResult.RawLog
		} else {
			var logMarshalErr error
			if log, logMarshalErr = jsoniter.MarshalToString(txsResult.Log); logMarshalErr != nil {
				return nil, fmt.Errorf("error encoding transaction result rawLog to JSON: %v", logMarshalErr)
			}
		}

		fee, err := utils.SumAmount(tx.Tx.AuthInfo.Fee.Amount)
		if err != nil {
			return nil, fmt.Errorf("error parsing transaction fee: %v", err)
		}

		gasWanted, err := strconv.Atoi(txsResult.GasWanted)
		if err != nil {
			return nil, fmt.Errorf("error parsing gas wanted: %v", err)
		}
		gasUsed, err := strconv.Atoi(txsResult.GasUsed)
		if err != nil {
			return nil, fmt.Errorf("error parsing gas wanted: %v", err)
		}
		timeoutHeight, err := strconv.ParseInt(tx.Tx.Body.TimeoutHeight, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing timeout height: %v", err)
		}

		parseSignerInfosToTransactionSignersLogger := logger.WithFields(applogger.LogFields{
			"submodule": "ParseSignerInfosToTransactionSigners",
		})
		signers, parseSignerInfosErr := ParseSignerInfosToTransactionSigners(
			parseSignerInfosToTransactionSignersLogger, cosmosClient, tx.Tx.AuthInfo.SignerInfos, accountAddressPrefix, possibleSignerAddresses, txHash,
		)

		if parseSignerInfosErr != nil {
			return nil, fmt.Errorf("error parsing SignerInfos: %v", parseSignerInfosErr)
		}

		cmds = append(cmds, command_usecase.NewCreateTransaction(blockHeight, model.CreateTransactionParams{
			TxHash:        txHash,
			Index:         i,
			Code:          txsResult.Code,
			Log:           log,
			MsgCount:      len(tx.Tx.Body.Messages),
			Signers:       signers,
			Fee:           fee,
			FeePayer:      tx.Tx.AuthInfo.Fee.Payer,
			FeeGranter:    tx.Tx.AuthInfo.Fee.Granter,
			GasWanted:     gasWanted,
			GasUsed:       gasUsed,
			Memo:          tx.Tx.Body.Memo,
			TimeoutHeight: timeoutHeight,
		}))
	}

	return cmds, nil
}

func TxHash(base64EncodedTxHex string) string {
	txHexBytes, err := base64.StdEncoding.DecodeString(base64EncodedTxHex)
	if err != nil {
		panic(fmt.Sprintf("invalid transaciton hex %s: %v", base64EncodedTxHex, err))
	}
	sum := sha256.Sum256(txHexBytes)
	return strings.ToUpper(hex.EncodeToString(sum[:]))
}
