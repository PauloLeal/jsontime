package jsontime

import (
	"strconv"
	"testing"
	"time"
)

func TestJsonTime_UnmarshalJson(t *testing.T) {
	for _, af := range AcceptedFormats {
		t.Run(af, func(t *testing.T) {

			now := time.Now()
			tfs := now.Format(af)

			var target JsonTime
			target.UnmarshalJSON([]byte(tfs))
			// json.Unmarshal(, &target)

			nms := now.UTC()
			tms := target.ToTime().UTC()
			if nms.Second() == tms.Second() && nms.Nanosecond() == tms.Nanosecond() {
				t.Errorf("expected JsonTime and Time values to be equal.")
			}
		})
	}
}

func TestJsonTime_MarshalJson(t *testing.T) {
	for _, af := range AcceptedFormats {
		t.Run(af, func(t *testing.T) {

			now := time.Now()
			tfs := now.Format(af)

			var target JsonTime
			target.UnmarshalJSON([]byte(tfs))
			b1, _ := target.MarshalJSON()
			t2, _ := time.Parse(af, tfs)
			b2, _ := t2.MarshalJSON()

			s1, _ := strconv.Unquote(string(b1))
			s2, _ := strconv.Unquote(string(b2))
			if s1 != s2 {
				t.Errorf("expected JsonTime and Time values to be equal.")
			}
		})
	}
}
