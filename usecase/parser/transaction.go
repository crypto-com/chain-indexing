package parser

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/crypto-com/chainindex/entity/command"
	command_usecase "github.com/crypto-com/chainindex/usecase/command"
	"github.com/crypto-com/chainindex/usecase/model"
	jsoniter "github.com/json-iterator/go"
)

func ParseTransactionCommands(
	txDecoder *TxDecoder,
	block *model.Block,
	blockResults *model.BlockResults,
) ([]command.Command, error) {
	blockHeight := blockResults.Height
	cmds := make([]command.Command, 0, len(blockResults.TxsResults))
	for i, txHex := range block.Txs {
		txsResult := blockResults.TxsResults[i]
		tx, err := txDecoder.Decode(txHex)
		if err != nil {
			panic(fmt.Sprintf("error decoding transaction: %v", err))
		}

		var log string
		if len(txsResult.Log) == 0 {
			// cater for failed transaction
			log = txsResult.RawLog
		} else {
			var logMarshalErr error
			if log, logMarshalErr = jsoniter.MarshalToString(txsResult.Log); logMarshalErr != nil {
				return nil, fmt.Errorf("error encoding transaction result rawLog to JSON: %v", err)
			}
		}

		fee, err := txDecoder.GetFee(txHex)
		if err != nil {
			return nil, fmt.Errorf("error parsing transaction fee: %v", err)
		}
		cmds = append(cmds, command_usecase.NewCreateTransaction(blockHeight, model.CreateTransactionParams{
			TxHash:    TxHash(txHex),
			Code:      txsResult.Code,
			Log:       log,
			MsgCount:  len(tx.Body.Messages),
			Fee:       fee,
			GasWanted: txsResult.GasWanted,
			GasUsed:   txsResult.GasUsed,
		}))
	}

	return cmds, nil
}

//func getTxFee(feeCollectorAddress string, txsResult model.BlockResultsTxsResult) coin.Coin {
//	for _, event := range txsResult.Events {
//		if event.Type == "transfer" {
//			isFeeEvent := false
//			var amount string
//			for _, attribute := range event.Attributes {
//				if attribute.Key == "recipient" && attribute.Value == feeCollectorAddress {
//					isFeeEvent = true
//				} else if attribute.Key == "amount" {
//					amount = attribute.Value
//				}
//			}
//
//			if isFeeEvent {
//				return coin.MustNewCoinFromString(TrimAmountDenom(amount))
//			}
//		}
//	}
//
//	return coin.MustNewCoinFromInt(int64(0))
//}

func TxHash(base64EncodedTxHex string) string {
	txHexBytes, err := base64.StdEncoding.DecodeString(base64EncodedTxHex)
	if err != nil {
		panic(fmt.Sprintf("invalid transaciton hex %s: %v", base64EncodedTxHex, err))
	}
	sum := sha256.Sum256(txHexBytes)
	return strings.ToUpper(hex.EncodeToString(sum[:]))
}
