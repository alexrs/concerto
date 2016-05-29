package main

import (
	"fmt"
	"github.com/alexrs95/concerto/io"
	"github.com/alexrs95/concerto/setlist"
	"github.com/alexrs95/concerto/spotify"
	sp "github.com/zmb3/spotify"
	"log"
	"os"
)

func main() {
	filePath := os.Args[1]
	fmt.Println(filePath)
	s, err := io.ReadLines(filePath)
	if err != nil {
		log.Fatal(err)
	}

	client := spotify.DoAuth()
	tracks := []sp.SimpleTrack{}
	for _, e := range s {
		list, err := setlist.GetSongList(e)
		// if no error
		if err == nil {
			max := 10
			if len(list) < max {
				max = len(list) - 1
			}
			tracks = append(tracks, spotify.SearchSong(e, list[:max])...)
		}
	}
	user, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}
	playlist, err := client.CreatePlaylistForUser(user.ID, "Prueba NOS", false)
	if err != nil {
		log.Fatal(err)
	}

	spotify.AddTracksToPlaylist(client, user.ID, playlist.SimplePlaylist.ID, convertTracksToID(tracks))
}

func convertTracksToID(tracks []sp.SimpleTrack) []sp.ID {
	ids := make([]sp.ID, 0, len(tracks))
	for _, e := range tracks {
		ids = append(ids, e.ID)
	}
	fmt.Println(ids)
	return ids
}
