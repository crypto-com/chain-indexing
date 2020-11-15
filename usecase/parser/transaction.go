package parser

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	jsoniter "github.com/json-iterator/go"

	"github.com/crypto-com/chainindex/entity/command"
	"github.com/crypto-com/chainindex/usecase/coin"
	command_usecase "github.com/crypto-com/chainindex/usecase/command"
	"github.com/crypto-com/chainindex/usecase/model"
)

func ParseTransactionCommands(
	moduleAccounts *ModuleAccounts,
	block *model.Block,
	blockResults *model.BlockResults,
) ([]command.Command, error) {
	blockHeight := blockResults.Height
	cmds := make([]command.Command, 0, len(blockResults.TxsResults))
	for i, txsResult := range blockResults.TxsResults {
		txHex := block.Txs[i]

		log, err := jsoniter.MarshalToString(txsResult.Log)
		if err != nil {
			return nil, fmt.Errorf("error encoding transaction result rawLog to JSON: %v", err)
		}
		cmds = append(cmds, command_usecase.NewCreateTransaction(blockHeight, model.CreateTransactionParams{
			TxHash:    TxHash(txHex),
			Code:      txsResult.Code,
			Log:       log,
			MsgCount:  len(txsResult.Log),
			Fee:       getTxFee(moduleAccounts.FeeCollector, txsResult),
			GasWanted: txsResult.GasWanted,
			GasUsed:   txsResult.GasUsed,
		}))
	}

	return cmds, nil
}

func getTxFee(feeCollectorAddress string, txsResult model.BlockResultsTxsResult) coin.Coin {
	for _, event := range txsResult.Events {
		if event.Type == "transfer" {
			isFeeEvent := false
			var amount string
			for _, attribute := range event.Attributes {
				if attribute.Key == "recipient" && attribute.Value == feeCollectorAddress {
					isFeeEvent = true
				} else if attribute.Key == "amount" {
					amount = attribute.Value
				}
			}

			if isFeeEvent {
				return coin.MustNewCoinFromString(TrimAmountDenom(amount))
			}
		}
	}

	return coin.MustNewCoinFromInt(int64(0))
}

func TxHash(base64EncodedTxHex string) string {
	txHexBytes, err := base64.StdEncoding.DecodeString(base64EncodedTxHex)
	if err != nil {
		panic(fmt.Sprintf("invalid transaciton hex %s: %v", base64EncodedTxHex, err))
	}
	sum := sha256.Sum256(txHexBytes)
	return strings.ToUpper(hex.EncodeToString(sum[:]))
}
