package rdb

import (
	"math/big"

	"github.com/crypto-com/chainindex/internal/utctime"
)

type TypeConv interface {
	// convert big.Int to native number pointer. Return nil if big.Int is nil
	Bton(*big.Int) interface{}
	Iton(int) interface{}
	// native number reader and parser to big.Int
	NtobReader() NtobReader

	// convert time.Time to native time format pointer. Return nil if time
	// is nil
	Tton(*utctime.UTCTime) interface{}
	// native time reader and parser to time.Time
	NtotReader() NtotReader
}

type NtobReader interface {
	// returns reference to a scannable argument of native numeric type
	ScannableArg() interface{}
	// parse the scannable argument reference to big.Int
	Parse() (*big.Int, error)
}

type NtotReader interface {
	// returns reference to a scannable argument of native time type
	ScannableArg() interface{}
	// parse the scannable argument reference to time.Time
	Parse() (*utctime.UTCTime, error)
}
