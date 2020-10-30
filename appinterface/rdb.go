package appinterface

// Relational database interface

type RDbConn interface {
	Begin() (RDbTx, error)
	Exec(sql string, args ...interface{}) (RDbExecResult, error)
	Query(sql string, args ...interface{}) (RDbRowsResult, error)
	QueryRow(sql string, args ...interface{}) RDbRowResult
}

type RDbTx interface {
	Exec(sql string, args ...interface{}) (RDbExecResult, error)
	Query(sql string, args ...interface{}) (RDbRowsResult, error)
	QueryRow(sql string, args ...interface{}) RDbRowResult
	Commit() error
	Rollback() error
}

// Implementing RDbConn and RDbTx interface automatically fulfills RDbRunner
type RDbRunner interface {
	Exec(sql string, args ...interface{}) (RDbExecResult, error)
	Query(sql string, args ...interface{}) (RDbRowsResult, error)
	QueryRow(sql string, args ...interface{}) RDbRowResult
}

type RDbExecResult interface {
	RowsAffected() int64
	IsInsert() bool
	IsUpdate() bool
	IsDelete() bool
	IsSelect() bool
	String() string
}

type RDbRowsResult interface {
	Close()
	Err() error
	ExecResult() RDbExecResult
	Next() bool
	Scan(dest ...interface{}) error
}

type RDbRowResult interface {
	Scan(dest ...interface{}) error
}
