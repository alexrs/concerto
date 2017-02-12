package main

import (
	"bufio"
	"log"
	"os"
	"sync"

	"fmt"

	"github.com/alexrs95/concerto/pkg/concerto"
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
	client := concerto.DoAuth()
	// Now, the user is obtained to create the playlist.
	user, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}

	// iterate over the list of groups to get the songs
	tracks := []spotify.SimpleTrack{}
	var wg sync.WaitGroup
	for _, e := range groups {
		// Increment the WaitGroup counter.
		wg.Add(1)
		go addSongs(e, &tracks, &wg)
	}
	wg.Wait()

	playlist, err := client.CreatePlaylistForUser(user.ID, playlistName, false)
	if err != nil {
		log.Fatal(err)
	}
	// Finally, the songs are added to the playlist
	concerto.AddTracksToPlaylist(client, user.ID, playlist.SimplePlaylist.ID, concerto.ConvertTracksToID(tracks))
}

func addSongs(s string, tracks *[]spotify.SimpleTrack, wg *sync.WaitGroup) {
	defer wg.Done()
	list, err := concerto.GetSongList(s)
	// if no error
	if err == nil {
		// max number of songs per artist. This will be a parameter in the future
		max := 10
		// if the number of songs returned is lower than the max, its value
		// is updated.
		if len(list) < max {
			max = len(list) - 1
		}
		*tracks = append(*tracks, concerto.SearchSong(s, list[:max])...)
	}
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
