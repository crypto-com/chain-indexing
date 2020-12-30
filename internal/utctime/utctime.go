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

func (t *UTCTime) Before(target UTCTime) bool {
	return t.unixNano < target.unixNano
}

func (t *UTCTime) After(target UTCTime) bool {
	return t.unixNano >= target.unixNano
}
