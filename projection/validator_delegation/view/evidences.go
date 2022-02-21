package view

import (
	"fmt"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/usecase/model"
	jsoniter "github.com/json-iterator/go"
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

	rawEvidenceJSON, err := jsoniter.MarshalToString(row.RawEvidence)
	if err != nil {
		return fmt.Errorf("error JSON marshalling RawEvidence for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Insert(
			"view_validator_delegation_evidences",
		).
		Columns(
			"height",
			"tendermint_address",
			"infraction_height",
			"raw_evidence",
		).
		Values(
			row.Height,
			row.TendermintAddress,
			row.InfractionHeight,
			rawEvidenceJSON,
		).
		ToSql()

	if err != nil {
		return fmt.Errorf("error building evidence insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := view.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error inserting evidence into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting evidence into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (view *EvidencesView) FindBy(height int64, tendermintAddress string) (EvidenceRow, error) {

	sql, sqlArgs, err := view.rdb.StmtBuilder.
		Select(
			"infraction_height",
			"raw_evidence",
		).
		From("view_validator_delegation_evidences").
		Where(
			"tendermint_address = ? AND height = ?",
			tendermintAddress,
			height,
		).
		ToSql()
	if err != nil {
		return EvidenceRow{}, fmt.Errorf("error building select evidence sql: %v: %w", err, rdb.ErrPrepare)
	}

	var evidence EvidenceRow
	evidence.Height = height
	evidence.TendermintAddress = tendermintAddress

	var rawEvidenceJSON string
	if err = view.rdb.QueryRow(sql, sqlArgs...).Scan(
		&evidence.InfractionHeight,
		&rawEvidenceJSON,
	); err != nil {
		return EvidenceRow{}, fmt.Errorf("error scanning evidence: %v: %w", err, rdb.ErrQuery)
	}

	if err = jsoniter.UnmarshalFromString(rawEvidenceJSON, &evidence.RawEvidence); err != nil {
		return EvidenceRow{}, fmt.Errorf("error unmarshalling RawEvidence JSON: %v: %w", err, rdb.ErrQuery)
	}

	return evidence, nil
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
