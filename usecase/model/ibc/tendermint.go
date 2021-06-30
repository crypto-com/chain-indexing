package ibc

import "github.com/crypto-com/chain-indexing/usecase/model"

type TendermintLightClientState struct {
	Type                         string                           `mapstructure:"@type" json:"@type"`
	ChainID                      string                           `mapstructure:"chain_id" json:"chainId"`
	TrustLevel                   TendermintLightClientTrustLevel  `mapstructure:"trust_level" json:"trustLevel"`
	TrustingPeriod               model.Duration                   `mapstructure:"trusting_period" json:"trustingPeriod"`
	UnbondingPeriod              model.Duration                   `mapstructure:"unbonding_period" json:"unbondingPeriod"`
	MaxClockDrift                model.Duration                   `mapstructure:"max_clock_drift" json:"maxClockDrift"`
	FrozenHeight                 TendermintLightClientHeight      `mapstructure:"frozen_height" json:"frozenHeight"`
	LatestHeight                 TendermintLightClientHeight      `mapstructure:"latest_height" json:"latestHeight"`
	ProofSpecs                   []TendermintLightClientProofSpec `mapstructure:"proof_specs" json:"proofSpecs"`
	UpgradePath                  []string                         `mapstructure:"upgrade_path" json:"upgradePath"`
	AllowUpdateAfterExpiry       bool                             `mapstructure:"allow_update_after_expiry" json:"allowUpdateAfterExpiry"`
	AllowUpdateAfterMisbehaviour bool                             `mapstructure:"allow_update_after_misbehaviour" json:"allowUpdateAfterMisbehaviour"`
}

type TendermintLightClientHeight struct {
	RevisionNumber uint64 `mapstructure:"revision_number" json:"revisionNumber,string"`
	RevisionHeight uint64 `mapstructure:"revision_height" json:"revisionHeight,string"`
}

type TendermintLightClientProofSpec struct {
	LeafSpec  TendermintLightClientLeafSpec  `mapstructure:"leaf_spec" json:"leafSpec"`
	InnerSpec TendermintLightClientInnerSpec `mapstructure:"inner_spec" json:"innerSpec"`
	MaxDepth  int32                          `mapstructure:"max_depth" json:"maxDepth"`
	MinDepth  int32                          `mapstructure:"min_depth" json:"minDepth"`
}

type TendermintLightClientInnerSpec struct {
	ChildOrder      []int32                     `mapstructure:"child_order" json:"childOrder"`
	ChildSize       int32                       `mapstructure:"child_size" json:"childSize"`
	MinPrefixLength int32                       `mapstructure:"min_prefix_length" json:"minPrefixLength"`
	MaxPrefixLength int32                       `mapstructure:"max_prefix_length" json:"maxPrefixLength"`
	EmptyChild      []byte                      `mapstructure:"empty_child" json:"emptyChild"`
	Hash            TendermintLightClientHashOp `mapstructure:"hash" json:"hash"`
}

type TendermintLightClientLeafSpec struct {
	Hash         TendermintLightClientHashOp   `mapstructure:"hash" json:"hash"`
	PrehashKey   TendermintLightClientHashOp   `mapstructure:"prehash_key" json:"prehashKey"`
	PrehashValue TendermintLightClientHashOp   `mapstructure:"prehash_value" json:"prehashValue"`
	Length       TendermintLightClientLengthOp `mapstructure:"length" json:"length"`
	Prefix       []byte                        `mapstructure:"prefix" json:"prefix"`
}

type TendermintLightClientHashOp string

type TendermintLightClientLengthOp string

type TendermintLightClientTrustLevel struct {
	Numerator   uint64 `mapstructure:"numerator" json:"numerator,string"`
	Denominator uint64 `mapstructure:"denominator" json:"denominator,string"`
}

type TendermintLightClientConsensusState struct {
	Type               string                    `mapstructure:"@type" json:"@type"`
	Timestamp          string                    `mapstructure:"timestamp" json:"timestamp"`
	Root               TendermintLightClientRoot `mapstructure:"root" json:"root"`
	NextValidatorsHash string                    `mapstructure:"next_validators_hash" json:"nextValidatorsHash"`
}

type TendermintLightClientRoot struct {
	Hash string `mapstructure:"hash" json:"hash"`
}
