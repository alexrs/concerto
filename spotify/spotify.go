package spotify

import (
	"log"

	"github.com/alexrs95/concerto/setlist"
	"github.com/zmb3/spotify"
)

func DoAuth() *spotify.Client {
	return doAuth()
}

func AddTracksToPlaylist(client *spotify.Client, userID string, playlistID spotify.ID, tracks []spotify.ID) {
	len := len(tracks)
	max := 100
	iter := (len / max)
	if iter == 0 {
		addTracks(client, userID, playlistID, tracks, 0, len)
	} else {
		start := 0
		for i := 0; i <= iter; i++ {
			addTracks(client, userID, playlistID, tracks, start, max)
			start += 100
			len = len - 100
			if len < 100 {
				max += len
			} else {
				max += 100
			}
		}
	}

}

func addTracks(client *spotify.Client, userID string, playlistID spotify.ID,
	tracks []spotify.ID, start int, max int) spotify.ID {
	snapshotID, err := client.AddTracksToPlaylist(userID, playlistID, tracks[start:max]...)
	if err != nil {
		log.Println("Ups:", err, "start:", start, "end:", max, tracks[start:max])
	}
	return spotify.ID(snapshotID)
}

// SearchSong returns
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
