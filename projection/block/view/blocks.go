package view

import (
	"errors"
	"fmt"
	"strings"

	"github.com/crypto-com/chain-indexing/appinterface/projection/view"

	sq "github.com/Masterminds/squirrel"

	"github.com/crypto-com/chain-indexing/appinterface/pagination"

	jsoniter "github.com/json-iterator/go"

	"github.com/crypto-com/chain-indexing/appinterface/rdb"
	"github.com/crypto-com/chain-indexing/internal/utctime"
	_ "github.com/crypto-com/chain-indexing/test/factory"
)

// Block projection view implemented by relational database
type Blocks struct {
	rdb *rdb.Handle
}

func NewBlocks(handle *rdb.Handle) *Blocks {
	return &Blocks{
		handle,
	}
}

func (blocksView *Blocks) Insert(block *Block) error {
	var err error

	var sql string
	sql, _, err = blocksView.rdb.StmtBuilder.Insert(
		"view_blocks",
	).Columns(
		"height",
		"hash",
		"time",
		"app_hash",
		"committed_council_nodes",
		"transaction_count",
	).Values("?", "?", "?", "?", "?", "?").ToSql()
	if err != nil {
		return fmt.Errorf("error building blocks insertion sql: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	var committedCouncilNodesJSON string
	if committedCouncilNodesJSON, err = jsoniter.MarshalToString(block.CommittedCouncilNodes); err != nil {
		return fmt.Errorf("error JSON marshalling blocks committed council nodes for insertion: %v: %w", err, rdb.ErrBuildSQLStmt)
	}

	result, err := blocksView.rdb.Exec(sql,
		block.Height,
		block.Hash,
		blocksView.rdb.Tton(&block.Time),
		block.AppHash,
		committedCouncilNodesJSON,
		block.TransactionCount,
	)
	if err != nil {
		return fmt.Errorf("error inserting block into the table: %v: %w", err, rdb.ErrWrite)
	}
	if result.RowsAffected() != 1 {
		return fmt.Errorf("error inserting block into the table: no rows inserted: %w", rdb.ErrWrite)
	}

	return nil
}

func (blocksView *Blocks) List(order BlocksListOrder, pagination *pagination.Pagination) ([]Block, *pagination.PaginationResult, error) {
	stmtBuilder := blocksView.rdb.StmtBuilder.Select(
		"height",
		"hash",
		"time",
		"app_hash",
		"committed_council_nodes",
		"transaction_count",
	).From(
		"view_blocks",
	)

	if order.Height == view.ORDER_DESC {
		stmtBuilder = stmtBuilder.OrderBy("height DESC")
	} else {
		stmtBuilder = stmtBuilder.OrderBy("height")
	}

	rDbPagination := rdb.NewRDbPaginationBuilder(
		pagination,
		blocksView.rdb,
	).WithCustomTotalQueryFn(
		func(rdbHandle *rdb.Handle, _ sq.SelectBuilder) (int64, error) {
			var total int64
			if err := rdbHandle.QueryRow("SELECT height FROM view_blocks ORDER BY height DESC LIMIT 1").Scan(&total); err != nil {
				return int64(0), err
			}
			return total, nil
		},
	).BuildStmt(stmtBuilder)
	sql, sqlArgs, err := rDbPagination.ToStmtBuilder().ToSql()
	if err != nil {
		return nil, nil, fmt.Errorf("error building blocks select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := blocksView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, nil, fmt.Errorf("error executing blocks select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	blocks := make([]Block, 0)
	for rowsResult.Next() {
		var block Block
		var committedCouncilNodesJSON *string
		timeReader := blocksView.rdb.NtotReader()
		if err = rowsResult.Scan(
			&block.Height,
			&block.Hash,
			timeReader.ScannableArg(),
			&block.AppHash,
			&committedCouncilNodesJSON,
			&block.TransactionCount,
		); err != nil {
			if errors.Is(err, rdb.ErrNoRows) {
				return nil, nil, rdb.ErrNoRows
			}
			return nil, nil, fmt.Errorf("error scanning block row: %v: %w", err, rdb.ErrQuery)
		}
		blockTime, parseErr := timeReader.Parse()
		if parseErr != nil {
			return nil, nil, fmt.Errorf("error parsing block time: %v: %w", parseErr, rdb.ErrQuery)
		}
		block.Time = *blockTime

		var committedCouncilNodes []BlockCommittedCouncilNode
		if unmarshalErr := jsoniter.Unmarshal([]byte(*committedCouncilNodesJSON), &committedCouncilNodes); unmarshalErr != nil {
			return nil, nil, fmt.Errorf("error unmarshalling block council nodes JSON: %v: %w", unmarshalErr, rdb.ErrQuery)
		}

		block.CommittedCouncilNodes = committedCouncilNodes

		blocks = append(blocks, block)
	}

	paginationResult, err := rDbPagination.Result()
	if err != nil {
		return nil, nil, fmt.Errorf("error preparing pagination result: %v", err)
	}

	return blocks, paginationResult, nil
}

func (blocksView *Blocks) FindBy(identity *BlockIdentity) (*Block, error) {
	var err error

	selectStmtBuilder := blocksView.rdb.StmtBuilder.Select(
		"height", "hash", "time", "app_hash", "committed_council_nodes", "transaction_count",
	).From("view_blocks")
	if identity.MaybeHash != nil {
		selectStmtBuilder = selectStmtBuilder.Where("hash = ?", *identity.MaybeHash)
	} else {
		selectStmtBuilder = selectStmtBuilder.Where("height = ?", *identity.MaybeHeight)
	}

	sql, sqlArgs, err := selectStmtBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building block selection sql: %v: %w", err, rdb.ErrPrepare)
	}

	var block Block
	var committedCouncilNodesJSON *string
	timeReader := blocksView.rdb.NtotReader()
	if err = blocksView.rdb.QueryRow(sql, sqlArgs...).Scan(
		&block.Height,
		&block.Hash,
		timeReader.ScannableArg(),
		&block.AppHash,
		&committedCouncilNodesJSON,
		&block.TransactionCount,
	); err != nil {
		if errors.Is(err, rdb.ErrNoRows) {
			return nil, rdb.ErrNoRows
		}
		return nil, fmt.Errorf("error scanning block row: %v: %w", err, rdb.ErrQuery)
	}
	blockTime, err := timeReader.Parse()
	if err != nil {
		return nil, fmt.Errorf("error parsing block time: %v: %w", err, rdb.ErrQuery)
	}
	block.Time = *blockTime

	var committedCouncilNodes []BlockCommittedCouncilNode
	if err = jsoniter.Unmarshal([]byte(*committedCouncilNodesJSON), &committedCouncilNodes); err != nil {
		return nil, fmt.Errorf("error unmarshalling block council nodes JSON: %v: %w", err, rdb.ErrQuery)
	}

	block.CommittedCouncilNodes = committedCouncilNodes

	return &block, nil
}

func (blocksView *Blocks) Search(
	keyword string,
) ([]Block, error) {
	keyword = strings.ToUpper(keyword)
	sql, sqlArgs, err := blocksView.rdb.StmtBuilder.Select(
		"height",
		"hash",
		"time",
		"app_hash",
		"committed_council_nodes",
		"transaction_count",
	).From(
		"view_blocks",
	).Where(
		"height = ? OR hash = ?", keyword, keyword,
	).OrderBy(
		"height",
	).Limit(5).ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building blocks select SQL: %v, %w", err, rdb.ErrBuildSQLStmt)
	}

	rowsResult, err := blocksView.rdb.Query(sql, sqlArgs...)
	if err != nil {
		return nil, fmt.Errorf("error executing blocks select SQL: %v: %w", err, rdb.ErrQuery)
	}
	defer rowsResult.Close()

	blocks := make([]Block, 0)
	for rowsResult.Next() {
		var block Block
		var committedCouncilNodesJSON *string
		timeReader := blocksView.rdb.NtotReader()
		if err = rowsResult.Scan(
			&block.Height,
			&block.Hash,
			timeReader.ScannableArg(),
			&block.AppHash,
			&committedCouncilNodesJSON,
			&block.TransactionCount,
		); err != nil {
			if errors.Is(err, rdb.ErrNoRows) {
				return nil, rdb.ErrNoRows
			}
			return nil, fmt.Errorf("error scanning block row: %v: %w", err, rdb.ErrQuery)
		}
		blockTime, parseErr := timeReader.Parse()
		if parseErr != nil {
			return nil, fmt.Errorf("error parsing block time: %v: %w", parseErr, rdb.ErrQuery)
		}
		block.Time = *blockTime

		var committedCouncilNodes []BlockCommittedCouncilNode
		if unmarshalErr := jsoniter.Unmarshal([]byte(*committedCouncilNodesJSON), &committedCouncilNodes); unmarshalErr != nil {
			return nil, fmt.Errorf("error unmarshalling block council nodes JSON: %v: %w", unmarshalErr, rdb.ErrQuery)
		}

		block.CommittedCouncilNodes = committedCouncilNodes

		blocks = append(blocks, block)
	}

	return blocks, nil
}

func (blocksView *Blocks) Count() (int64, error) {
	sql, _, err := blocksView.rdb.StmtBuilder.Select("MAX(height)").From(
		"view_blocks",
	).ToSql()
	if err != nil {
		return 0, fmt.Errorf("error building blocks count selection sql: %v", err)
	}

	result := blocksView.rdb.QueryRow(sql)
	var count *int64
	if err := result.Scan(&count); err != nil {
		return 0, fmt.Errorf("error scanning blocks count selection query: %v", err)
	}

	if count == nil {
		return 0, nil
	}
	return *count, nil
}

func NewRdbBlockCommittedCouncilNodeFromRaw(raw *BlockCommittedCouncilNode) *RdbBlockCommittedCouncilNode {
	return &RdbBlockCommittedCouncilNode{
		Address:    raw.Address,
		Time:       raw.Time.UnixNano(),
		Signature:  raw.Signature,
		IsProposer: raw.IsProposer,
	}
}

func (node *RdbBlockCommittedCouncilNode) ToRaw() *BlockCommittedCouncilNode {
	return &BlockCommittedCouncilNode{
		Address:    node.Address,
		Time:       utctime.FromUnixNano(node.Time),
		Signature:  node.Signature,
		IsProposer: node.IsProposer,
	}
}

type Block struct {
	Height                int64                       `json:"blockHeight" fake:"{+int64}"`
	Hash                  string                      `json:"blockHash" fake:"{blockhash}"`
	Time                  utctime.UTCTime             `json:"blockTime" fake:"{utctime}"`
	AppHash               string                      `json:"appHash" fake:"{apphash}"`
	TransactionCount      int                         `json:"transactionCount" fake:"{number:0,2147483647}"`
	CommittedCouncilNodes []BlockCommittedCouncilNode `json:"committedCouncilNodes" fakesize:"3"`
}

type BlockCommittedCouncilNode struct {
	Address    string          `json:"address" fake:"{validatoraddress}"`
	Time       utctime.UTCTime `json:"time" fake:"{utctime}"`
	Signature  string          `json:"signature" fake:"{commitsignature}"`
	IsProposer bool            `json:"isProposer" fake:"{bool}"`
}

type RdbBlockCommittedCouncilNode struct {
	Address    string `json:"address"`
	Time       int64  `json:"time"`
	Signature  string `json:"signature"`
	IsProposer bool   `json:"isProposer"`
}

type BlockIdentity struct {
	MaybeHeight *int64
	MaybeHash   *string
}

type BlocksListOrder struct {
	Height view.ORDER
}
