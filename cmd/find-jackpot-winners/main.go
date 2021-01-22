package main

import (
	"fmt"
	"os"
	"strconv"

	jsoniter "github.com/json-iterator/go"

	"github.com/crypto-com/chain-indexing/appinterface/projection/crossfire"
	"github.com/crypto-com/chain-indexing/infrastructure/tendermint"
	"github.com/crypto-com/chain-indexing/internal/base64"
	"github.com/crypto-com/chain-indexing/internal/tmcosmosutils"
	"github.com/crypto-com/chain-indexing/usecase/parser"
)

func main() {
	blockHeightStr := os.Args[1]
	outputFile := os.Args[2]

	blockHeight, err := strconv.ParseInt(blockHeightStr, 10, 64)
	if err != nil {
		panic(fmt.Errorf("error parsing block height argument: %v", err))
	}

	crossfireClient := crossfire.NewHTTPClient("https://chain.crypto.com/testnet-participants.json")
	participants, err := crossfireClient.Participants()
	if err != nil {
		panic(fmt.Errorf("error getting testnet participants: %v", err))
	}

	tendermintClient := tendermint.NewHTTPClient("https://crossfire.crypto.com")
	block, _, err := tendermintClient.Block(blockHeight)
	if err != nil {
		panic(fmt.Errorf("error getting block response: %v", err))
	}
	blockResults, err := tendermintClient.BlockResults(blockHeight)
	if err != nil {
		panic(fmt.Errorf("error getting block results response: %v", err))
	}

	txDecoder := parser.NewTxDecoder()
	winnerSet := make(map[string]bool)
	for i, tx := range block.Txs {
		blockResult := blockResults.TxsResults[i]
		decodedTx, decodeErr := txDecoder.Decode(tx)
		if decodeErr != nil {
			panic(fmt.Errorf("error decoding raw tranaction %d: %v", i, err))
		}

		if blockResult.Code != 0 {
			fmt.Printf("skipped %d transaction because it is a failed transaction: %s\n", i, blockResult.RawLog)
		}

		for _, signer := range decodedTx.AuthInfo.SignerInfos {
			var address string
			if signer.ModeInfo.MaybeSingle != nil {
				pubkey := base64.MustDecodeString(*signer.PublicKey.MaybeKey)
				address = tmcosmosutils.MustAccountAddressFromPubKey("cro", pubkey)
			} else {
				pubkeys := make([][]byte, 0, len(signer.PublicKey.MaybePublicKeys))
				for _, pubkey := range signer.PublicKey.MaybePublicKeys {
					pubkeys = append(pubkeys, base64.MustDecodeString(pubkey.Key))
				}
				sort := false
				address = tmcosmosutils.MustMultiSigAddressFromPubKeys(
					"cro", pubkeys, *signer.PublicKey.MaybeThreshold, sort,
				)
			}

			winnerSet[address] = true
		}
	}

	participantMap := make(map[string]crossfire.Participant)
	for _, participant := range *participants {
		participantMap[participant.PrimaryAddress] = participant
	}

	winners := make([]crossfire.Participant, 0)
	for address := range winnerSet {
		if participant, exist := participantMap[address]; exist {
			winners = append(winners, participant)
		}
	}

	fd, err := os.Create(outputFile)
	if err != nil {
		panic(fmt.Errorf("error creating output file to write: %v", err))
	}
	if err := jsoniter.NewEncoder(fd).Encode(winners); err != nil {
		panic(fmt.Errorf("error encoding JSON to file: %v", err))
	}
	fmt.Println("Done")
}
