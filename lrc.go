package main

import (
	"bufio"
	"io"
	"regexp"
)

type Lyric struct {
	Time Timestamp
	Text string
}

func parseLyric(line string) *Lyric {
	timeRegex := regexp.MustCompile("^\\[(?P<Time>\\d{2}:\\d{2}.\\d{2})](?P<Line>.*$)")
	match := timeRegex.FindStringSubmatch(line)

	if len(match) == 0 {
		return nil
	}

	return &Lyric{
		Time: parseTimestamp(match[1]),
		Text: match[2],
	}
}

func ParseLRC(reader io.Reader) ([]Lyric, error) {
	scanner := bufio.NewScanner(reader)

	lyrics := []Lyric{}

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		line := scanner.Text()
		lyric := parseLyric(line)
		if lyric != nil {
			lyrics = append(lyrics, *lyric)
		}
	}

	return lyrics, nil
}
