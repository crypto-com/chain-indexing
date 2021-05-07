package ibc

type MsgConnectionOpenInitParams struct {
	ClientID          string                            `json:"clientId"`
	Counterparty      MsgConnectionOpenInitCounterparty `json:"counterparty"`
	ConnectionVersion MsgConnectionOpenInitVersion      `json:"connectionVersion"`
	DelayPeriod       string                            `json:"delayPeriod"`
	Signer            string                            `json:"signer"`

	ConnectionID string `json:"connectionId"`
}

type MsgConnectionOpenInitCounterparty struct {
	ClientID     string                      `json:"clientId"`
	ConnectionID string                      `json:"connectionId"`
	Prefix       MsgConnectionOpenInitPrefix `json:"prefix"`
}

type MsgConnectionOpenInitPrefix struct {
	KeyPrefix []byte `json:"keyPrefix"`
}

type MsgConnectionOpenInitVersion struct {
	Identifier string   `json:"identifier"`
	Features   []string `json:"features"`
}
