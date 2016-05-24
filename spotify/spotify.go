package spotify

import (
	"github.com/zmb3/spotify"
)

func DoAuth() {
	doAuth()
}

func SearchSong(title string) ([]spotify.FullTrack, error) {
	return searchSong(title)
}
