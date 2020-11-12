package parser

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	jsoniter "github.com/json-iterator/go"

	"github.com/crypto-com/chainindex/entity/command"
	"github.com/crypto-com/chainindex/usecase/domain/createtransaction"
	"github.com/crypto-com/chainindex/usecase/model"
)

func ParseTransactionCommands(block *model.Block, blockResults *model.BlockResults) ([]command.Command, error) {
	blockHeight := blockResults.Height
	cmds := make([]command.Command, 0, len(blockResults.TxsResults))
	for i, txsResult := range blockResults.TxsResults {
		txHex := block.Txs[i]

		log, err := jsoniter.MarshalToString(txsResult.Log)
		if err != nil {
			return nil, fmt.Errorf("error encoding transaction result log to JSON: %v", err)
		}
		cmds = append(cmds, createtransaction.NewCommand(blockHeight, createtransaction.Params{
			TxHash:    TxHash(txHex),
			Code:      txsResult.Code,
			Log:       log,
			MsgCount:  len(txsResult.Log),
			GasWanted: txsResult.GasWanted,
			GasUsed:   txsResult.GasUsed,
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
