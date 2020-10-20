package appinterface

import (
	"math/big"
	"time"
)

type RDbTypeConv interface {
	// convert big.Int to native number pointer. Return nil if big.Int is nil
	Bton(*big.Int) interface{}
	Iton(int) interface{}
	// native number reader and parser to big.Int
	NtobReader() RDbNtobReader

	// convert time.Time to native time format pointer. Return nil if time
	// is nil
	Tton(*time.Time) interface{}
	// native time reader and parser to time.Time
	NtotReader() RDbNtotReader
}

type RDbNtobReader interface {
	// returns reference to a scannable argument of native numeric type
	ScannableArg() interface{}
	// parse the scannable argument reference to big.Int
	Parse() (*big.Int, error)
}

type RDbNtotReader interface {
	// returns reference to a scannable argument of native time type
	ScannableArg() interface{}
	// parse the scannable argument reference to time.Time
	Parse() (*time.Time, error)
}
