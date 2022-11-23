package tpm

import "time"

type ShortDate struct {
	time.Time
}

func (t *ShortDate) UnmarshalJSON(b []byte) (err error) {
	value := string(b)

	if value == "null" || value == "" {
		return
	}

	date, err := time.Parse(`"2006-01-02"`, value)
	if err != nil {
		return err
	}
	t.Time = date
	return
}

type LongDate struct {
	time.Time
}

func (t *LongDate) UnmarshalJSON(b []byte) (err error) {
	value := string(b)

	if value == "null" || value == "" {
		return
	}

	date, err := time.Parse(`"2006-01-02 15:04:05"`, value)
	if err != nil {
		return err
	}
	t.Time = date
	return
}
