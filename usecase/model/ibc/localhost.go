package ibc

type LocalhostLightClientState struct {
	Type    string                          `mapstructure:"@type" json:"@type"`
	ChainID string                          `mapstructure:"chain_id" json:"chainId"`
	Height  LocalhostLightClientStateHeight `mapstructure:"height" json:"height"`
}
type LocalhostLightClientStateHeight struct {
	RevisionNumber uint64 `mapstructure:"revision_number" json:"revisionNumber,string"`
	RevisionHeight uint64 `mapstructure:"revision_height" json:"revisionHeight,string"`
}
