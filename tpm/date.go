package tpm

import (
	"strings"
	"time"
)

const (
	shortTimeFormat = "2006-01-02"
	longTimeFormat  = "2006-01-02 15:04:05"
)

func marshalCustomDateToJSON(t time.Time, format string) ([]byte, error) {
	if t.IsZero() {
		return []byte(`""`), nil
	}

	return []byte(`"` + t.Format(format) + `"`), nil
}

func unmarshalCustomDateToJSON(b []byte, format string) (time.Time, error) {
	value := strings.Trim(string(b), `"`)
	if value == "" || value == "null" {
		return time.Time{}, nil
	}

	return time.Parse(format, value)
}

type ShortDate time.Time

func (d ShortDate) MarshalJSON() ([]byte, error) {
	return marshalCustomDateToJSON(time.Time(d), shortTimeFormat)
}

func (d *ShortDate) UnmarshalJSON(b []byte) (err error) {
	t, err := unmarshalCustomDateToJSON(b, shortTimeFormat)
	if err != nil {
		return err
	}

	*d = ShortDate(t)

	return nil
}

type LongDate time.Time

func (d *LongDate) UnmarshalJSON(b []byte) (err error) {
	t, err := unmarshalCustomDateToJSON(b, longTimeFormat)
	if err != nil {
		return err
	}

	*d = LongDate(t)

	return nil
}

func (d LongDate) MarshalJSON() ([]byte, error) {
	return marshalCustomDateToJSON(time.Time(d), longTimeFormat)
}
