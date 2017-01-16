package main

import (
	"bufio"
	"log"
	"os"

	"fmt"

	"github.com/alexrs95/concerto/setlist"
	sp "github.com/alexrs95/concerto/spotify"
	"github.com/zmb3/spotify"
)

func main() {
	// check if the program is called with two args
	if len(os.Args) < 3 {
		log.Fatal(`Error. 
Run: concerto artistFile playlistName
		`)
	}
	filePath := os.Args[1]
	playlistName := os.Args[2]

	// get the group/artist names
	groups, err := readLines(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Authenticate on spotify
	client := sp.DoAuth()

	// iterate over the list of groups to get the songs
	tracks := []spotify.SimpleTrack{}
	for _, e := range groups {
		//TODO - Paralelize this for
		list, err := setlist.GetSongList(e)
		// if no error
		if err == nil {
			// max number of songs per artist. This will be a parameter in the future
			max := 10
			// if the number of songs returned is lower than the max, its value
			// is updated.
			if len(list) < max {
				max = len(list) - 1
			}
			tracks = append(tracks, sp.SearchSong(e, list[:max])...)
		}
	}

	// Now, the user is obtained to create the playlist.
	user, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}
	playlist, err := client.CreatePlaylistForUser(user.ID, playlistName, false)
	if err != nil {
		log.Fatal(err)
	}

	// Finally, the songs are added to the playlist
	sp.AddTracksToPlaylist(client, user.ID, playlist.SimplePlaylist.ID, convertTracksToID(tracks))
}

//readLines returns a slice with the lines of a given file
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) == 0 {
		return nil, fmt.Errorf("Empty file")
	}
	return lines, scanner.Err()
}

// returns a list of ids to add the songs to the playlist
func convertTracksToID(tracks []spotify.SimpleTrack) []spotify.ID {
	// Make a slice of len 0 and capacity len(tracks)
	ids := make([]spotify.ID, 0, len(tracks))
	for _, e := range tracks {
		ids = append(ids, e.ID)
	}
	return ids
}
