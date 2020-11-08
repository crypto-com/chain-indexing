package chain

import (
	"fmt"
	"github.com/crypto-com/chainindex/infrastructure/tendermint"
	"time"
)

const POLLING_INTERVAL = 5

func LatestBlockHeightGenerator(client tendermint.HTTPClient) chan int64 {
	latestBlockHeightChan := make(chan int64)

	go func() {
		for {
			// get latest block height of the chain
			height, err := client.LatestBlockHeight()
			if err != nil {
				// TODO: is there global logger?
				fmt.Println("error getting chain's latest block height", err)
			}
			latestBlockHeightChan <- height

			<-time.After(POLLING_INTERVAL * time.Second)
		}
	}()

	return latestBlockHeightChan
}
