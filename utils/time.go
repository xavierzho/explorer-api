package utils

import (
	"strconv"
	"time"
)

// Time is a wrapper over big.Int to implement only unmarshalText
// for json decoding.
type Time time.Time

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (t *Time) UnmarshalText(text []byte) (err error) {
	input, err := strconv.ParseInt(string(text), 10, 64)
	if err != nil {
		return
	}

	var timestamp = time.Unix(input, 0)
	*t = Time(timestamp)

	return nil
}

// Time returns t's time.Time form
func (t *Time) Time() time.Time {
	return time.Time(*t)
}

// MarshalText implements the encoding.TextMarshaler
func (t *Time) MarshalText() (text []byte, err error) {
	return []byte(strconv.FormatInt(t.Time().Unix(), 10)), nil
}
func (t *Time) FormatDate() string {
	return t.Time().Format("2006-01-02")
}
