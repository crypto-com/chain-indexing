package model

import (
	"encoding/json"
	"time"

	"github.com/crypto-com/chain-indexing/external/utctime"
	"github.com/crypto-com/chain-indexing/usecase/coin"
)

type MsgSubmitProposalContent struct {
	Type string `json:"@type"`
}

type MsgSubmitCommunityPoolSpendProposalParams struct {
	MaybeProposalId *string                                    `json:"proposalId"`
	Content         MsgSubmitCommunityPoolSpendProposalContent `json:"content"`
	ProposerAddress string                                     `json:"proposerAddress"`
	InitialDeposit  coin.Coins                                 `json:"initialDeposit"`
}
type MsgSubmitCommunityPoolSpendProposalContent struct {
	Type             string     `json:"@type"`
	Title            string     `json:"title"`
	Description      string     `json:"description"`
	RecipientAddress string     `json:"recipientAddress"`
	Amount           coin.Coins `json:"amount"`
}
type RawMsgSubmitCommunityPoolSpendProposalContent struct {
	Type             string        `json:"@type"`
	Title            string        `json:"title"`
	Description      string        `json:"description"`
	RecipientAddress string        `json:"recipient"`
	Amount           []interface{} `json:"amount"`
}

type MsgSubmitParamChangeProposalParams struct {
	MaybeProposalId *string                             `json:"proposalId"`
	Content         MsgSubmitParamChangeProposalContent `json:"content"`
	ProposerAddress string                              `json:"proposerAddress"`
	InitialDeposit  coin.Coins                          `json:"initialDeposit"`
}
type MsgSubmitParamChangeProposalContent struct {
	Type        string                               `json:"@type"`
	Title       string                               `json:"title"`
	Description string                               `json:"description"`
	Changes     []MsgSubmitParamChangeProposalChange `json:"changes"`
}
type MsgSubmitParamChangeProposalChange struct {
	Subspace string          `json:"subspace"`
	Key      string          `json:"key"`
	Value    json.RawMessage `json:"value"`
}

// MsgSubmitSoftwareUpgradeProposal
type MsgSubmitSoftwareUpgradeProposalParams struct {
	MaybeProposalId *string                                 `json:"proposalId"`
	Content         MsgSubmitSoftwareUpgradeProposalContent `json:"content"`
	ProposerAddress string                                  `json:"proposerAddress"`
	InitialDeposit  coin.Coins                              `json:"initialDeposit"`
}
type MsgSubmitSoftwareUpgradeProposalContent struct {
	Type        string                               `json:"@type"`
	Title       string                               `json:"title"`
	Description string                               `json:"description"`
	Plan        MsgSubmitSoftwareUpgradeProposalPlan `json:"plan"`
}
type MsgSubmitSoftwareUpgradeProposalPlan struct {
	Name   string          `json:"name"`
	Time   utctime.UTCTime `json:"time"`
	Height int64           `json:"height"`
	Info   string          `json:"info"`
}

type RawMsgSubmitSoftwareUpgradeProposalContent struct {
	Type        string                                  `json:"@type"`
	Title       string                                  `json:"title"`
	Description string                                  `json:"description"`
	Plan        RawMsgSubmitSoftwareUpgradeProposalPlan `json:"plan"`
}
type RawMsgSubmitSoftwareUpgradeProposalPlan struct {
	Name   string    `json:"name"`
	Time   time.Time `json:"time"`
	Height string    `json:"height"`
	Info   string    `json:"info"`
}

// MsgSubmitSoftwareUpgrade
type MsgSubmitSoftwareUpgradeParams struct {
	MaybeProposalId *string                          `json:"proposalId"`
	Metadata        string                           `json:"metadata"`
	Messages        MsgSubmitSoftwareUpgradeMessages `json:"messages"`
	Proposer        string                           `json:"proposer"`
	InitialDeposit  coin.Coins                       `json:"initialDeposit"`
}

type MsgSubmitSoftwareUpgradeMessages struct {
	Plan MsgSubmitSoftwareUpgradeProposalPlan `json:"plan"`
}

type RawMsgSubmitSoftwareUpgradeContent struct {
	Type string                                  `json:"@type"`
	Plan RawMsgSubmitSoftwareUpgradeProposalPlan `json:"plan"`
}

type MsgSubmitCancelSoftwareUpgradeProposalParams struct {
	MaybeProposalId *string                                       `json:"proposalId"`
	Content         MsgSubmitCancelSoftwareUpgradeProposalContent `json:"content"`
	ProposerAddress string                                        `json:"proposerAddress"`
	InitialDeposit  coin.Coins                                    `json:"initialDeposit"`
}
type MsgSubmitCancelSoftwareUpgradeProposalContent struct {
	Type        string `json:"@type"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type MsgSubmitCancelSoftwareUpgradeParams struct {
	MaybeProposalId *string                                `json:"proposalId"`
	Metadata        string                                 `json:"metadata"`
	Messages        MsgSubmitCancelSoftwareUpgradeMessages `json:"messages"`
	Proposer        string                                 `json:"proposer"`
	InitialDeposit  coin.Coins                             `json:"initialDeposit"`
}
type MsgSubmitCancelSoftwareUpgradeMessages struct {
	Type string `json:"@type"`
}

type MsgSubmitTextProposalParams struct {
	MaybeProposalId *string                      `json:"proposalId"`
	Content         MsgSubmitTextProposalContent `json:"content"`
	ProposerAddress string                       `json:"proposerAddress"`
	InitialDeposit  coin.Coins                   `json:"initialDeposit"`
}
type MsgSubmitTextProposalContent struct {
	Type        string `json:"@type"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type MsgSubmitUnknownProposalParams struct {
	MaybeProposalId *string                         `json:"proposalId"`
	Content         MsgSubmitUnknownProposalContent `json:"content"`
	ProposerAddress string                          `json:"proposerAddress"`
	InitialDeposit  coin.Coins                      `json:"initialDeposit"`
}

type MsgSubmitUnknownProposalContent struct {
	Type        string      `json:"@type"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	RawContent  interface{} `json:"rawContent"`
}
