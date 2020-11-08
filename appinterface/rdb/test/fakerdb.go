package test

import (
	"math/big"

	"github.com/crypto-com/chainindex/appinterface/rdb"
)

type FakeRDbConn struct{}

func NewFakeRDbConn() *FakeRDbConn {
	return &FakeRDbConn{}
}

func (conn *FakeRDbConn) Begin() (rdb.Tx, error) {
	return &FakeRDbTx{}, nil
}
func (conn *FakeRDbConn) Exec(sql string, args ...interface{}) (rdb.ExecResult, error) {
	return &FakeRDbExecResult{}, nil
}
func (conn *FakeRDbConn) Query(sql string, args ...interface{}) (rdb.RowsResult, error) {
	return &FakeRDbRowsResult{}, nil
}
func (conn *FakeRDbConn) QueryRow(sql string, args ...interface{}) rdb.RowResult {
	return &FakeRDbRowResult{}
}
func (conn *FakeRDbConn) ToHandle() *rdb.Handle { return nil }

type FakeRDbTx struct{}

func (tx *FakeRDbTx) Exec(sql string, args ...interface{}) (rdb.ExecResult, error) {
	return &FakeRDbExecResult{}, nil
}
func (tx *FakeRDbTx) Query(sql string, args ...interface{}) (rdb.RowsResult, error) {
	return &FakeRDbRowsResult{}, nil
}
func (tx *FakeRDbTx) QueryRow(sql string, args ...interface{}) rdb.RowResult {
	return &FakeRDbRowResult{}
}
func (tx *FakeRDbTx) Commit() error         { return nil }
func (tx *FakeRDbTx) Rollback() error       { return nil }
func (tx *FakeRDbTx) ToHandle() *rdb.Handle { return nil }

type FakeRDbRowsResult struct{}

func (result *FakeRDbRowsResult) Close()     {}
func (result *FakeRDbRowsResult) Err() error { return nil }
func (result *FakeRDbRowsResult) ExecResult() rdb.ExecResult {
	return &FakeRDbExecResult{}
}
func (result *FakeRDbRowsResult) Next() bool                     { return false }
func (result *FakeRDbRowsResult) Scan(dest ...interface{}) error { return nil }

type FakeRDbExecResult struct{}

func (result *FakeRDbExecResult) RowsAffected() int64 { return 0 }
func (result *FakeRDbExecResult) IsInsert() bool      { return false }
func (result *FakeRDbExecResult) IsUpdate() bool      { return false }
func (result *FakeRDbExecResult) IsDelete() bool      { return false }
func (result *FakeRDbExecResult) IsSelect() bool      { return false }
func (result *FakeRDbExecResult) String() string      { return "" }

type FakeRDbRowResult struct{}

func (result *FakeRDbRowResult) Scan(dest ...interface{}) error { return nil }

type FakeRDbTypeConv struct{}

func (conv *FakeRDbTypeConv) Bton(bigInt *big.Int) interface{} { return nil }
func (conv *FakeRDbTypeConv) Iton(i int) interface{}           { return nil }
func (conv *FakeRDbTypeConv) NtobReader() rdb.NtobReader {
	return &FakeRDbNtobReader{}
}

type FakeRDbNtobReader struct{}

func (reader *FakeRDbNtobReader) ScannableArg() interface{} { return nil }
func (reader *FakeRDbNtobReader) Parse() (*big.Int, error) {
	return nil, nil
}
