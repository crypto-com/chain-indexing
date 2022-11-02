package view

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
)

type ValidatorActiveBlocks interface {
	UpdateValidatorsActiveBlocks(
		operatorAddressToSignedBlockFlagMap OperatorAddressToSignedBlockFlagMap,
		height int64,
		maxRecentUpTimeInBlocks int64,
	) error
	ClearValidatorActiveBLocks(operatorAddress string) error
}

const VALIDATOR_ACTIVE_BLOCKS_TABLE_NAME = "view_validator_active_blocks"

// ValidatorActiveBlocksView projection view implemented by relational database
type ValidatorActiveBlocksView struct {
	rdb *rdb.Handle
}

func NewValidatorActiveBlocksView(handle *rdb.Handle) ValidatorActiveBlocks {
	return &ValidatorActiveBlocksView{
		handle,
	}
}

func (validatorActiveBlocksView *ValidatorActiveBlocksView) UpdateValidatorsActiveBlocks(
	operatorAddressToSignedBlockFlagMap OperatorAddressToSignedBlockFlagMap,
	height int64,
	maxRecentActiveBlocks int64,
) error {
	expiredRecentUpTimeBlock := int64(0)

	if height-maxRecentActiveBlocks > 0 {
		expiredRecentUpTimeBlock = height - maxRecentActiveBlocks
	}

	pendingRowCount := 0
	var stmtBuilder sq.InsertBuilder

	validatorsCount := len(operatorAddressToSignedBlockFlagMap)
	i := 0
	for operatorAddress, signed := range operatorAddressToSignedBlockFlagMap {
		if pendingRowCount == 0 {
			stmtBuilder = validatorActiveBlocksView.rdb.StmtBuilder.Insert(
				VALIDATOR_ACTIVE_BLOCKS_TABLE_NAME,
			).Columns(
				"block_height",
				"operator_address",
				"signed",
			)
		}

		stmtBuilder = stmtBuilder.Values(
			height,
			operatorAddress,
			signed,
		)
		pendingRowCount += 1

		if pendingRowCount == 500 || i+1 == validatorsCount {
			sql, sqlArgs, err := stmtBuilder.ToSql()
			if err != nil {
				return fmt.Errorf("error building validator active block insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
			}

			result, err := validatorActiveBlocksView.rdb.Exec(sql, sqlArgs...)
			if err != nil {
				return fmt.Errorf("error inserting validator active block into the table: %v: %w", err, rdb.ErrWrite)
			}
			if result.RowsAffected() != int64(pendingRowCount) {
				return fmt.Errorf("error inserting validator active block into the table: mismatch rows inserted: %w", rdb.ErrWrite)
			}

			pendingRowCount = 0
		}

		i += 1
	}

	sql, sqlArgs, err := validatorActiveBlocksView.rdb.StmtBuilder.Delete(
		VALIDATOR_ACTIVE_BLOCKS_TABLE_NAME,
	).Where(
		"block_height <= ?", expiredRecentUpTimeBlock,
	).ToSql()
	if err != nil {
		return fmt.Errorf("error building validator active block deletion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	_, err = validatorActiveBlocksView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error deleting validator active block from the table: %v: %w", err, rdb.ErrWrite)
	}

	return nil
}

func (validatorActiveBlocksView *ValidatorActiveBlocksView) ClearValidatorActiveBLocks(operatorAddress string) error {

	sql, sqlArgs, err := validatorActiveBlocksView.rdb.StmtBuilder.Delete(
		VALIDATOR_ACTIVE_BLOCKS_TABLE_NAME,
	).Where(
		"operator_address = ?", operatorAddress,
	).ToSql()
	if err != nil {
		return fmt.Errorf("error building validator active block deletion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	_, err = validatorActiveBlocksView.rdb.Exec(sql, sqlArgs...)
	if err != nil {
		return fmt.Errorf("error deleting validator active block from the table: %v: %w", err, rdb.ErrWrite)
	}

	return nil
}

type OperatorAddressToSignedBlockFlagMap map[string]bool

type ValidatorActiveBlockCountFilter struct {
	MaybeOperatorAddresses []string
}

type ValidatorActiveBlockRow struct {
	OperatorAddress string `json:"operatorAddress"`
	BlockHeight     int64  `json:"blockHeight"`
	Signed          bool   `json:"signed"`
}
