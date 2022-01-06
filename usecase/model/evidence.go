package model

type EvidenceParams struct {
	InfractionHeight  int64         `json:"infractionHeight"`
	TendermintAddress string        `json:"tendermintAddress"`
	RawEvidence       BlockEvidence `json:"rawEvidence"`
}
