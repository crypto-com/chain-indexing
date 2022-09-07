package model

type BlockEvent struct {
	Type       string                `json:"type"`
	Attributes []BlockEventAttribute `json:"attributes"`
}

type BlockEventAttribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
