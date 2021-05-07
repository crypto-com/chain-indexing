package ibcmsg

import "time"

type TendermintLightClientState struct {
	Type                         string                           `json:"@type" mapstructure:"@type"`
	ChainID                      string                           `json:"chain_id" mapstructure:"chain_id"`
	TrustLevel                   TendermintLightClientTrustLevel  `json:"trust_level" mapstructure:"trust_level"`
	TrustingPeriod               time.Duration                    `json:"trusting_period" mapstructure:"trusting_period"`
	UnbondingPeriod              time.Duration                    `json:"unbonding_period" mapstructure:"unbonding_period"`
	MaxClockDrift                time.Duration                    `json:"max_clock_drift" mapstructure:"max_clock_drift"`
	FrozenHeight                 TendermintLightClientHeight      `json:"frozen_height" mapstructure:"frozen_height"`
	LatestHeight                 TendermintLightClientHeight      `json:"latest_height" mapstructure:"latest_height"`
	ProofSpecs                   []TendermintLightClientProofSpec `json:"proof_specs" mapstructure:"proof_specs"`
	UpgradePath                  []string                         `json:"upgrade_path" mapstructure:"upgrade_path"`
	AllowUpdateAfterExpiry       bool                             `json:"allow_update_after_expiry" mapstructure:"allow_update_after_expiry"`
	AllowUpdateAfterMisbehaviour bool                             `json:"allow_update_after_misbehaviour" mapstructure:"allow_update_after_misbehaviour"`
}

type TendermintLightClientHeight struct {
	RevisionNumber string `json:"revision_number" mapstructure:"revision_number"`
	RevisionHeight string `json:"revision_height" mapstructure:"revision_height"`
}

type TendermintLightClientProofSpec struct {
	MaybeLeafSpec  *TendermintLightClientLeafSpec  `json:"leaf_spec" mapstructure:"leaf_spec"`
	MaybeInnerSpec *TendermintLightClientInnerSpec `json:"inner_spec" mapstructure:"inner_spec"`
	MaxDepth       int64                           `json:"max_depth" mapstructure:"max_depth"`
	MinDepth       int64                           `json:"min_depth" mapstructure:"min_depth"`
}

type TendermintLightClientInnerSpec struct {
	ChildOrder      []int64 `json:"child_order" mapstructure:"child_order"`
	ChildSize       int64   `json:"child_size" mapstructure:"child_size"`
	MinPrefixLength int64   `json:"min_prefix_length" mapstructure:"min_prefix_length"`
	MaxPrefixLength int64   `json:"max_prefix_length" mapstructure:"max_prefix_length"`
	EmptyChild      []byte  `json:"empty_child" mapstructure:"empty_child"`
	Hash            string  `json:"hash" mapstructure:"hash"`
}

type TendermintLightClientLeafSpec struct {
	Hash         string `json:"hash" mapstructure:"hash"`
	PrehashKey   string `json:"prehash_key" mapstructure:"prehash_key"`
	PrehashValue string `json:"prehash_value" mapstructure:"prehash_value"`
	Length       string `json:"length" mapstructure:"length"`
	Prefix       string `json:"prefix" mapstructure:"prefix"`
}

type TendermintLightClientTrustLevel struct {
	Numerator   string `json:"numerator" mapstructure:"numerator"`
	Denominator string `json:"denominator" mapstructure:"denominator"`
}

type TendermintLightClientConsensusState struct {
	Type               string                    `json:"@type" mapstructure:"@type"`
	Timestamp          string                    `json:"timestamp" mapstructure:"timestamp"`
	Root               TendermintLightClientRoot `json:"root" mapstructure:"root"`
	NextValidatorsHash string                    `json:"next_validators_hash" mapstructure:"next_validators_hash"`
}

type TendermintLightClientRoot struct {
	Hash string `json:"hash" mapstructure:"hash"`
}

type MsgConnectionOpenInitCounterparty struct {
	ClientID     string                      `json:"client_id" mapstructure:"client_id"`
	ConnectionID string                      `json:"connection_id" mapstructure:"connection_id"`
	Prefix       MsgConnectionOpenInitPrefix `json:"prefix" mapstructure:"prefix"`
}

type MsgConnectionOpenInitPrefix struct {
	KeyPrefix string `json:"key_prefix" mapstructure:"key_prefix"`
}

type MsgConnectionOpenInitVersion struct {
	Identifier string   `json:"identifier" mapstructure:"identifier"`
	Features   []string `json:"features" mapstructure:"features"`
}
