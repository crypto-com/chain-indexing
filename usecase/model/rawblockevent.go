package model

import "github.com/crypto-com/chain-indexing/external/utctime"

type CreateRawBlockEventParams struct {
	BlockHash  string
	BlockTime  utctime.UTCTime
	FromResult string
	Data       DataParams
}

type DataParams struct {
	Type    string            `json:"type"`
	Content BlockResultsEvent `json:"content"`
}
