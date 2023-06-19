package jsontime

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

// AcceptedFormats Change this to accept (or not) more formats
var AcceptedFormats = []string{
	"2006-01-02T15:04:050700",
	"2006-01-02T15:04:05-0700",
	"2006-01-02T15:04:05Z07:00",
	"2006-01-02T15:04:05Z-07:00",
	"2006-01-02T15:04:05.99999Z07:00",
	"2006-01-02T15:04:05.99999Z-07:00",
	"2006-01-02 15:04:05.9999", // extra - mariadb format
	"2006-01-02 15:04:05",      // extra - mariadb format
}

type JsonTime time.Time

// ToTime Helper to avoid lots of casts in code
func (jt *JsonTime) ToTime() time.Time {
	return time.Time(*jt)
}

func (jt *JsonTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(jt.ToTime())
}

func (jt *JsonTime) UnmarshalJSON(b []byte) error {
	s := string(b)

	if ss, err := strconv.Unquote(s); err == nil {
		s = ss
	}

	if s == "" {
		return nil
	}

	for _, af := range AcceptedFormats {
		t, err := time.Parse(af, s)
		if err == nil {
			*jt = JsonTime(t)
			return nil
		}
	}

	return errors.New("unparsable date")
}
