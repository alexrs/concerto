package spotify

import (
	"errors"
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
