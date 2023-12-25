package main

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestParseLyric_EmptyString(t *testing.T) {
	expected := []Lyric{}
	reader := strings.NewReader("")

	parsed, err := ParseLRC(reader)
	if err != nil {
		t.Fatalf("error parsing lyrics: %v", err)
	}

	if !reflect.DeepEqual(expected, parsed) {
		t.Fatalf("expected %v, got %v", expected, parsed)
	}
}

func TestParseLyric_PartInvalidString(t *testing.T) {
	lyrics := `
some
invalid value
[01:01.00]Test
`
	expected := []Lyric{
		{
			Timestamp{time.Minute*1 + 1*time.Second + 0*time.Second},
			"Test",
		},
	}

	reader := strings.NewReader(lyrics)

	parsed, err := ParseLRC(reader)
	if err != nil {
		t.Fatalf("error parsing lyrics: %v", err)
	}

	if !reflect.DeepEqual(expected, parsed) {
		t.Fatalf("expected %v, got %v", expected, parsed)
	}
}

func TestParseLyric_InvalidString(t *testing.T) {
	lyrics := `
some
invalid value
`
	expected := []Lyric{}

	reader := strings.NewReader(lyrics)

	parsed, err := ParseLRC(reader)
	if err != nil {
		t.Fatalf("error parsing lyrics: %v", err)
	}

	if !reflect.DeepEqual(expected, parsed) {
		t.Fatalf("expected %v, got %v", expected, parsed)
	}
}

func TestParseLyric(t *testing.T) {
	lyrics := `
[00:01.37]Bury all your secrets in my skin
[00:02.47]Come away with innocence, and leave me with my sins
`

	reader := strings.NewReader(lyrics)

	parsed, err := ParseLRC(reader)
	if err != nil {
		t.Fatalf("error parsing lyrics: %v", err)
	}

	expected := []Lyric{
		{
			Time: Timestamp{1*time.Second + 37*time.Millisecond},
			Text: "Bury all your secrets in my skin",
		},
		{
			Time: Timestamp{2*time.Second + 47*time.Millisecond},
			Text: "Come away with innocence, and leave me with my sins",
		},
	}

	if parsed == nil {
		return
	}

	if !reflect.DeepEqual(parsed, expected) {
		t.Fatalf("expected %v, got %v", expected, parsed)
	}
}
