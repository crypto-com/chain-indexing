package ibc

type MsgConnectionOpenConfirmParams struct {
	RawMsgConnectionOpenConfirm

	ClientID                 string `json:"clientId"`
	CounterpartyClientID     string `json:"counterpartyClientId"`
	CounterpartyConnectionID string `json:"counterpartyConnectionId"`
}

type RawMsgConnectionOpenConfirm struct {
	ConnectionID string `mapstructure:"connection_id" json:"connectionId"`
	ProofACK     []byte `mapstructure:"proof_ack" json:"proofAck"`
	ProofHeight  Height `mapstructure:"proof_height" json:"proofHeight"`
	Signer       string `mapstructure:"signer" json:"signer"`
}
