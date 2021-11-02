package utctime

import (
	"fmt"
	"time"
	gotime "time"

	jsoniter "github.com/json-iterator/go"
)

// UTCTime is a wrapper of Golang time.Time in UTC
type UTCTime struct {
	unixNano int64
}

func FromUnixNano(unixNano int64) UTCTime {
	return UTCTime{
		unixNano,
	}
}

func FromTime(t gotime.Time) UTCTime {
	return UTCTime{
		t.UTC().UnixNano(),
	}
}

func MustParse(layout string, value string) UTCTime {
	parsed, err := Parse(layout, value)
	if err != nil {
		panic(err)
	}

	return parsed
}

func Parse(layout string, value string) (UTCTime, error) {
	t, err := gotime.Parse(layout, value)
	if err != nil {
		return UTCTime{}, err
	}

	return FromTime(t), nil
}

func Now() UTCTime {
	return UTCTime{
		gotime.Now().UTC().UnixNano(),
	}
}

func (t UTCTime) Add(d gotime.Duration) UTCTime {
	newTime := gotime.Unix(0, t.unixNano).UTC().Add(d)
	return FromUnixNano(newTime.UnixNano())
}

func (t UTCTime) Year() int {
	return gotime.Unix(0, t.unixNano).UTC().Year()
}

func (t UTCTime) Month() Month {
	switch gotime.Unix(0, t.unixNano).UTC().Month() {
	case gotime.January:
		return January
	case gotime.February:
		return February
	case gotime.March:
		return March
	case gotime.April:
		return April
	case gotime.May:
		return May
	case gotime.June:
		return June
	case gotime.July:
		return July
	case gotime.August:
		return August
	case gotime.September:
		return September
	case gotime.October:
		return October
	case gotime.November:
		return November
	case gotime.December:
		return December
	}
	return 0
}

func (t UTCTime) Day() int {
	return gotime.Unix(0, t.unixNano).UTC().Day()
}

func (t UTCTime) Hour() int {
	return gotime.Unix(0, t.unixNano).UTC().Hour()
}

func (t UTCTime) Minute() int {
	return gotime.Unix(0, t.unixNano).UTC().Minute()
}

func (t UTCTime) Second() int {
	return gotime.Unix(0, t.unixNano).UTC().Second()
}

func (t UTCTime) UnixNano() int64 {
	return t.unixNano
}

func (t UTCTime) String() string {
	return fmt.Sprint(gotime.Unix(0, t.unixNano).UTC())
}

func (t UTCTime) MarshalJSON() ([]byte, error) {
	result, err := jsoniter.Marshal(time.Unix(0, t.unixNano).UTC())
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t *UTCTime) UnmarshalJSON(data []byte) error {
	var timeVal gotime.Time
	if err := jsoniter.Unmarshal(data, &timeVal); err != nil {
		return err
	}
	t.unixNano = timeVal.UTC().UnixNano()

	return nil
}

// A Month specifies a month of the year (January = 1, ...).
type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)
