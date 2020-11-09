package model

type BlockResults struct {
	Height           int64
	TxsEvents        [][]BlockResultsEvent
	BeginBlockEvents []BlockResultsEvent
	ValidatorUpdates []BlockResultsValidator
}

type BlockResultsEvent struct {
	Type       string
	Attributes []BlockResultsEventAttribute
}
type BlockResultsEventAttribute struct {
	Key   string
	Value string
}

type BlockResultsValidator struct {
	PubKey BlockResultsValidatorPubKey
	Power  *string
}
type BlockResultsValidatorPubKey struct {
	Type    string
	PubKey  string
	Address string
}
