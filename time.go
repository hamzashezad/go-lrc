package main

import (
	"regexp"
	"strconv"
	"time"
)

type Timestamp struct {
	time.Duration
}

func parseInt64(s string) time.Duration {
	i, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		return 0
	}

	return time.Duration(i)
}

func parseTimestamp(str string) Timestamp {
	rg := regexp.MustCompile("(?P<Minutes>\\d+):(?P<Seconds>\\d+).(?P<Milliseconds>\\d+)")
	match := rg.FindStringSubmatch(str)

	t := Timestamp{}

	t.Duration = time.Duration(time.Minute*parseInt64(match[1]) + time.Second*parseInt64(match[2]) + time.Millisecond*parseInt64(match[3]))

	return t
}
