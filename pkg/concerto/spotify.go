package concerto

import (
	"log"

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

// ConvertTracksToID returns a list of ids given a list of tracks
func ConvertTracksToID(tracks []spotify.SimpleTrack) []spotify.ID {
	// make a slice of len 0 and capacity len(tracks)
	ids := make([]spotify.ID, 0, len(tracks))
	for _, e := range tracks {
		ids = append(ids, e.ID)
	}
	return ids
}

// SearchSong returns a list of songs from spotify
func SearchSong(artist string, titles []SongStats) []spotify.SimpleTrack {
	songs := []spotify.SimpleTrack{}
	for _, t := range titles {
		// search the song on Spotify
		song, err := searchSong(t.Name)
		if err == nil {
			// iterate over the list of songs
			for _, s := range song {
				// check if the song is already included in the list of songs,
				// the name of the artist coincides and the name of the song coincides
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

// addTracks add a list of songs to a paylist
func addTracks(client *spotify.Client, userID string, playlistID spotify.ID,
	tracks []spotify.ID) spotify.ID {
	snapshotID, err := client.AddTracksToPlaylist(userID, playlistID, tracks...)
	if err != nil {
		log.Println("Error adding tracks")
	}
	return spotify.ID(snapshotID)
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
