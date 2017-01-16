package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	sp "github.com/zmb3/spotify"
)

// Input: file
// Output: array of artists
func TestReadLines(t *testing.T) {
	expected := []string{"Green Day", "The Beatles", "The Killers", "The Kooks"}
	lines, err := readLines("testdata/test.txt")
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if !assert.ObjectsAreEqual(expected, lines) {
		t.Errorf("Expected %s, got %s", expected, lines)
	}
}

// Input: file that does not exist
// Output: error
func TestReadLinesFileNotExists(t *testing.T) {
	_, err := readLines("testdata/nonexisting.txt")
	if err == nil {
		t.Errorf("Expected err, got %v", err)
	}
}

// Input: empty file
// Output: error
func TestReadLinesEmptyFile(t *testing.T) {
	_, err := readLines("testdata/empty.txt")
	if err == nil {
		t.Errorf("Expected error, got %v", err)
	}
}

// Input: Empty slice
// Output: Empty slice
func TestConvertTracksToIDEmpty(t *testing.T) {
	tracks := []sp.SimpleTrack{}
	ids := convertTracksToID(tracks)

	if len(ids) != 0 {
		t.Errorf("Expected empty slice, got %v", tracks)
	}
}
