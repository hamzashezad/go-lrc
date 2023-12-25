package main

import (
	"bufio"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"
)

func main() {
	path := flag.String("f", "", "path to .lrc file")
	if *path == "" {
		slog.Error("path to .lrc is required")
		os.Exit(1)
	}
	flag.Parse()

	file, err := os.Open(*path)
	if err != nil {
		slog.Error("error opening file: %s", err)
	}

	reader := bufio.NewReader(file)

	slog.Info("parsing LRC file")

	donech := make(chan bool)
	ticker := time.NewTicker(time.Millisecond * 20)
	lyrics, err := ParseLRC(reader)
	if err != nil {
		slog.Error("error parsing LRC file: %s", err)
	}

	go playLRC(lyrics, ticker, donech)

	<-donech
	ticker.Stop()
}

func playLRC(lines []Lyric, ticker *time.Ticker, done chan bool) {
	now := time.Now()

	idx := 0
	for {
		select {
		case t := <-ticker.C:
			if t.After(now.Add(lines[idx].Time.Duration)) {
				fmt.Println(lines[idx].Text)

				idx += 1
			}

			if idx >= len(lines) {
				done <- true
				return
			}
		}
	}
}
