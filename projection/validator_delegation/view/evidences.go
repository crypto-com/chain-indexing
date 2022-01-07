package view

import (
	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/usecase/model"
)

type Evidences interface {
	Insert(row EvidenceRow) error
	FindBy(height int64, tendermintAddress string) (EvidenceRow, error)
}

type EvidencesView struct {
	rdb *rdb.Handle
}

func NewEvidencesView(handle *rdb.Handle) Evidences {
	return &EvidencesView{
		handle,
	}
}

func (view *EvidencesView) Insert(row EvidenceRow) error {

	return nil
}

func (view *EvidencesView) FindBy(height int64, tendermintAddress string) (EvidenceRow, error) {

	// TODO: It seems it is possible that tendermint could put two evidences on a same height.
	//       Maybe in this function, we should SELECT all matched rows and then return the first one in the list.

	return EvidenceRow{}, nil
}

// NOTES: Don't add UNIQUE constraint on (Height, TendermintAddress) now.
//        As we are not sure, in one block, whether two evidences on a same validator will be included.
//        (Might need to check Tendermint source code, or consult other colleagues?)
type EvidenceRow struct {
	// The height the evidence is included
	Height            int64  `json:"height"`
	TendermintAddress string `json:"tendermintAddress"`
	// The height the infraction happened
	InfractionHeight int64               `json:"infractionHeight"`
	RawEvidence      model.BlockEvidence `json:"rawEvidence"`
}
