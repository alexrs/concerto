package spotify

import (
	"fmt"
	"github.com/alexrs95/concerto/setlist"
	"github.com/zmb3/spotify"
)

func DoAuth() {
	doAuth()
}

func SearchSong(artist string, titles []setlist.SongStats) []spotify.SimpleTrack {
	l := []spotify.SimpleTrack{}
	for _, t := range titles {
		song, err := searchSong(t.Name)
		if err == nil {
			for _, s := range song {
				if containsArtist(artist, s.SimpleTrack.Artists) && isSong(t.Name, s.SimpleTrack.Name) {
					l = append(l, s.SimpleTrack)
					fmt.Println(s.SimpleTrack.ID)
				}
			}
		}
	}
	return l
}
