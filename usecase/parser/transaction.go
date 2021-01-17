package parser

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/crypto-com/chain-indexing/entity/command"
	command_usecase "github.com/crypto-com/chain-indexing/usecase/command"
	"github.com/crypto-com/chain-indexing/usecase/model"
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

		gasWanted, err := strconv.Atoi(txsResult.GasWanted)
		if err != nil {
			return nil, fmt.Errorf("error parsing gas wanted: %v", err)
		}
		gasUsed, err := strconv.Atoi(txsResult.GasUsed)
		if err != nil {
			return nil, fmt.Errorf("error parsing gas wanted: %v", err)
		}
		timeoutHeight, err := strconv.ParseInt(tx.Body.TimeoutHeight, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing timeout height: %v", err)
		}

		signers := make([]model.TransactionSigner, 0)
		for _, signer := range tx.AuthInfo.SignerInfos {
			sequence, parseErr := strconv.ParseUint(signer.Sequence, 10, 64)
			if parseErr != nil {
				return nil, fmt.Errorf("error parsing account sequence: %v", parseErr)
			}
			if signer.ModeInfo.MaybeSingle != nil {
				signers = append(signers, model.TransactionSigner{
					Type:            signer.PublicKey.Type,
					IsMultiSig:      false,
					Pubkeys:         []string{*signer.PublicKey.MaybeKey},
					AccountSequence: sequence,
				})
			} else {
				pubkeys := make([]string, 0, len(signer.PublicKey.MaybePublicKeys))
				for _, pubkey := range signer.PublicKey.MaybePublicKeys {
					pubkeys = append(pubkeys, pubkey.Key)
				}
				signers = append(signers, model.TransactionSigner{
					Type:            signer.PublicKey.Type,
					IsMultiSig:      true,
					Pubkeys:         pubkeys,
					AccountSequence: sequence,
				})
			}
		}

		cmds = append(cmds, command_usecase.NewCreateTransaction(blockHeight, model.CreateTransactionParams{
			TxHash:        TxHash(txHex),
			Index:         i,
			Code:          txsResult.Code,
			Log:           log,
			MsgCount:      len(tx.Body.Messages),
			Signers:       signers,
			Fee:           fee,
			FeePayer:      tx.AuthInfo.Fee.Payer,
			FeeGranter:    tx.AuthInfo.Fee.Granter,
			GasWanted:     gasWanted,
			GasUsed:       gasUsed,
			Memo:          tx.Body.Memo,
			TimeoutHeight: timeoutHeight,
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
//				if attribute.Key == "recipient" && attribute.Amount == feeCollectorAddress {
//					isFeeEvent = true
//				} else if attribute.Key == "amount" {
//					amount = attribute.Amount
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
