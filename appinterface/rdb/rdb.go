package rdb

import sq "github.com/Masterminds/squirrel"

// Relational database interface

type Conn interface {
	Begin() (Tx, error)
	Exec(sql string, args ...interface{}) (ExecResult, error)
	Query(sql string, args ...interface{}) (RowsResult, error)
	QueryRow(sql string, args ...interface{}) RowResult
	ToHandle() *Handle
}

type Tx interface {
	Exec(sql string, args ...interface{}) (ExecResult, error)
	Query(sql string, args ...interface{}) (RowsResult, error)
	QueryRow(sql string, args ...interface{}) RowResult
	Commit() error
	Rollback() error
	ToHandle() *Handle
}

type Handle struct {
	Runner
	TypeConv

	StmtBuilder *StatementBuilder
}

type StatementBuilder struct {
	sq.StatementBuilderType
	sq.PlaceholderFormat
}

func NewStatementBuilder(builder sq.StatementBuilderType, placeholderFormat sq.PlaceholderFormat) *StatementBuilder {
	return &StatementBuilder{
		builder,
		placeholderFormat,
	}
}

// Implementing Conn and Tx interface automatically fulfills Runner
type Runner interface {
	Exec(sql string, args ...interface{}) (ExecResult, error)
	Query(sql string, args ...interface{}) (RowsResult, error)
	QueryRow(sql string, args ...interface{}) RowResult
}

type ExecResult interface {
	RowsAffected() int64
	IsInsert() bool
	IsUpdate() bool
	IsDelete() bool
	IsSelect() bool
	String() string
}

type RowsResult interface {
	Close()
	Err() error
	ExecResult() ExecResult
	Next() bool
	// When no row was found it should return ErrNoRows
	Scan(dest ...interface{}) error
}

type RowResult interface {
	Scan(dest ...interface{}) error
}
