package model

import "github.com/crypto-com/chain-indexing/external/utctime"

type CreateRawBlockEventParams struct {
	BlockHash  string
	BlockTime  utctime.UTCTime
	FromResult string
	RawData    BlockResultsEvent
}
