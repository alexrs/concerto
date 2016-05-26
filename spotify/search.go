package spotify

import (
	"errors"
	"github.com/alexrs95/concerto/strop"
	"github.com/zmb3/spotify"
	"log"
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

func containsArtist(name string, artists []spotify.SimpleArtist) bool {
	for _, v := range artists {
		if strop.EditDistance(name, v.Name) < 5 {
			return true
		}
	}
	return false
}

func isSong(s1, s2 string) bool {
	if strop.EditDistance(s1, s2) < 5 {
		return true
	}
	return false
}
