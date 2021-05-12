package rdb

import (
	"math/big"

	"github.com/crypto-com/chain-indexing/internal/utctime"
)

type TypeConv interface {
	// convert big.Int to native number pointer. Return nil if big.Int is nil
	Bton(*big.Int) interface{}
	// convert big.Float to native number pointer. Return nil if big.Float is nil
	BFton(*big.Float) interface{}
	Iton(int) interface{}
	// native number reader and parser to big.Int
	NtobReader() NtobReader
	// native number reader and parser to big.Float
	NtobfReader() NtobfReader

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

type NtobfReader interface {
	// returns reference to a scannable argument of native numeric type
	ScannableArg() interface{}
	// parse the scannable argument reference to big.Float
	Parse() (*big.Float, error)
}

type NtotReader interface {
	// returns reference to a scannable argument of native time type
	ScannableArg() interface{}
	// parse the scannable argument reference to time.Time
	Parse() (*utctime.UTCTime, error)
}
