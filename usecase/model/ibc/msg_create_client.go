package ibc

import (
	"time"

	"github.com/crypto-com/chain-indexing/internal/utctime"
)

type MsgCreateClientParams struct {
	MaybeTendermintLightClient *TendermintLightClient `json:"maybeTendermintLightClient"`
	// TODO: SoloMachine and Localhost LightClient
	Signer string `json:"signer"`

	ClientID   string `json:"clientId"`
	ClientType string `json:"clientType"`
}

type TendermintLightClient struct {
	TendermintClientState               TendermintLightClientState          `json:"clientState"`
	TendermintLightClientConsensusState TendermintLightClientConsensusState `json:"consensusState"`
}

type TendermintLightClientState struct {
	Type                         string                           `json:"@type"`
	ChainID                      string                           `json:"chainId"`
	TrustLevel                   TendermintLightClientTrustLevel  `json:"trustLevel"`
	TrustingPeriod               time.Duration                    `json:"trustingPeriod"`
	UnbondingPeriod              time.Duration                    `json:"unbondingPeriod"`
	MaxClockDrift                time.Duration                    `json:"maxClockDrift"`
	FrozenHeight                 TendermintLightClientHeight      `json:"frozenHeight"`
	LatestHeight                 TendermintLightClientHeight      `json:"latestHeight"`
	ProofSpecs                   []TendermintLightClientProofSpec `json:"proofSpecs"`
	UpgradePath                  []string                         `json:"upgradePath"`
	AllowUpdateAfterExpiry       bool                             `json:"allowUpdateAfterExpiry"`
	AllowUpdateAfterMisbehaviour bool                             `json:"allowUpdateAfterMisbehaviour"`
}

type TendermintLightClientHeight struct {
	RevisionNumber string `json:"revisionNumber"`
	RevisionHeight string `json:"revisionHeight"`
}

type TendermintLightClientProofSpec struct {
	MaybeLeafSpec  *TendermintLightClientLeafSpec  `json:"leafSpec"`
	MaybeInnerSpec *TendermintLightClientInnerSpec `json:"innerSpec"`
	MaxDepth       int64                           `json:"maxDepth"`
	MinDepth       int64                           `json:"minDepth"`
}

type TendermintLightClientInnerSpec struct {
	ChildOrder      []int64 `json:"childOrder"`
	ChildSize       int64   `json:"childSize"`
	MinPrefixLength int64   `json:"minPrefixLength"`
	MaxPrefixLength int64   `json:"maxPrefixLength"`
	EmptyChild      []byte  `json:"emptyChild"`
	Hash            string  `json:"hash"`
}

type TendermintLightClientLeafSpec struct {
	Hash         string `json:"hash"`
	PrehashKey   string `json:"prehashKey"`
	PrehashValue string `json:"prehashValue"`
	Length       string `json:"length"`
	Prefix       string `json:"prefix"`
}

type TendermintLightClientTrustLevel struct {
	Numerator   string `json:"numerator"`
	Denominator string `json:"denominator"`
}

type TendermintLightClientConsensusState struct {
	Type               string                    `json:"@type"`
	Timestamp          utctime.UTCTime           `json:"timestamp"`
	Root               TendermintLightClientRoot `json:"root"`
	NextValidatorsHash string                    `json:"nextValidatorsHash"`
}

type TendermintLightClientRoot struct {
	Hash string `json:"hash"`
}
