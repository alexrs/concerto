package spotify

import (
	"log"

	"github.com/alexrs95/concerto/pkg/setlist"
	"github.com/zmb3/spotify"
)

// DoAuth returns a pointer to spotify.Client
func DoAuth() *spotify.Client {
	return doAuth()
}

// AddTracksToPlaylist add a list of tracks to a given playlist
func AddTracksToPlaylist(client *spotify.Client, userID string, playlistID spotify.ID, tracks []spotify.ID) {
	len := len(tracks)
	// the maximum number of songs that can be added at once to a playlist is 100
	max := 100
	// compute how many iterations are needed to add the songs
	iter := (len / max)

	start := 0
	end := max
	if len < max {
		end = len
	}
	for i := 0; i <= iter; i++ {
		addTracks(client, userID, playlistID, tracks[start:end])
		// increase the strart index of the first song to add
		start += max
		// decreases the len of the list
		len = len - max
		// if the new len is smaller than the max number of allowed songs to add, update the value
		// of end to avoid overflows
		if len < max {
			end += len
		} else {
			end += max
		}
	}
}

// addTracks add a list of songs to a paylist
func addTracks(client *spotify.Client, userID string, playlistID spotify.ID,
	tracks []spotify.ID) spotify.ID {
	snapshotID, err := client.AddTracksToPlaylist(userID, playlistID, tracks...)
	if err != nil {
		log.Println("Error adding tracks")
	}
	return spotify.ID(snapshotID)
}

// SearchSong returns a list of songs from spotify
func SearchSong(artist string, titles []setlist.SongStats) []spotify.SimpleTrack {
	songs := []spotify.SimpleTrack{}
	for _, t := range titles {
		song, err := searchSong(t.Name)
		if err == nil {
			for _, s := range song {
				if !containsTrack(s.SimpleTrack, songs) &&
					containsArtist(artist, s.SimpleTrack.Artists) &&
					isSong(t.Name, s.SimpleTrack.Name) {
					songs = append(songs, s.SimpleTrack)
				}
			}
		}
	}
	return songs
}

// containsTrack returns true if the list of songs contains a given song
func containsTrack(track spotify.SimpleTrack, list []spotify.SimpleTrack) bool {
	for _, v := range list {
		if v.Name == track.Name {
			return true
		}
	}
	return false
}
