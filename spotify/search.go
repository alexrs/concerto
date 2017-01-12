package spotify

import (
	"errors"
	"log"

	"github.com/antzucaro/matchr"
	"github.com/zmb3/spotify"
)

func searchSong(title string) ([]spotify.FullTrack, error) {
	res, err := spotify.Search(title, spotify.SearchTypeTrack)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if res.Tracks == nil {
		return nil, errors.New("empty track list")
	}
	return res.Tracks.Tracks, nil
}

// containsArtist returns true if the edit distance between the artist names is smaller than
// some threshold
func containsArtist(name string, artists []spotify.SimpleArtist) bool {
	for _, v := range artists {
		threshold := min(len(name), len(v.Name)) / 5
		if matchr.Levenshtein(name, v.Name) < threshold {
			return true
		}
	}
	return false
}

// isSong returns true if the edit distance between two titles is smaller than
// some threshold
func isSong(s1, s2 string) bool {
	threshold := min(len(s1), len(s2)) / 5
	if matchr.Levenshtein(s1, s2) < threshold {
		return true
	}
	return false
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
