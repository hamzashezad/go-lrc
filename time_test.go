package main

import (
	"reflect"
	"testing"
	"time"
)

func TestParseTimestamp(t *testing.T) {
	tbl := map[Timestamp]Timestamp{
		{time.Duration(0)}:                                          parseTimestamp("[00:00.00]"),
		{time.Duration(time.Minute * 10)}:                           parseTimestamp("[10:00.00][10:00.00]"),
		{time.Duration(time.Minute*100000 + time.Millisecond*2000)}: parseTimestamp("[-100000:00.2000]"),
	}

	for expected, got := range tbl {
		if !reflect.DeepEqual(got, expected) {
			t.Logf("expected %v got %v", expected, got)
			t.Fail()
		}
	}
}
