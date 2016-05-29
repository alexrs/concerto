package spotify

import (
	"github.com/alexrs95/concerto/setlist"
	"github.com/zmb3/spotify"
	"log"
)

func DoAuth() *spotify.Client {
	return doAuth()
}

func AddTracksToPlaylist(client *spotify.Client, userID string, playlistID spotify.ID, tracks []spotify.ID) {
	len := len(tracks)
	if len > 100 {
		len = 100
	}
	_, err := client.AddTracksToPlaylist(userID, playlistID, tracks[:len]...)
	if err != nil {
		log.Println(err)
	}
}

func SearchSong(artist string, titles []setlist.SongStats) []spotify.SimpleTrack {
	l := []spotify.SimpleTrack{}
	for _, t := range titles {
		song, err := searchSong(t.Name)
		if err == nil {
			for _, s := range song {
				if !containsTrack(s.SimpleTrack, l) &&
					containsArtist(artist, s.SimpleTrack.Artists) &&
					isSong(t.Name, s.SimpleTrack.Name) {
					l = append(l, s.SimpleTrack)
				}
			}
		}
	}
	return l
}

func containsTrack(track spotify.SimpleTrack, list []spotify.SimpleTrack) bool {
	for _, v := range list {
		if v.Name == track.Name {
			return true
		}
	}
	return false
}
