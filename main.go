package main

import (
	"fmt"
	"github.com/alexrs95/concerto/io"
	"github.com/alexrs95/concerto/setlist"
	"github.com/alexrs95/concerto/spotify"
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

	spotify.DoAuth()

	for _, e := range s {
		list, err := setlist.GetSongList(e)
		// if no error
		if err == nil {
			max := 10
			if len(list) < max {
				max = len(list) - 1
			}
			spotify.SearchSong(e, list[:max])
		}
	}
}
