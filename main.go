package main

import (
	"fmt"
	"github.com/alexrs95/concerto/io"
	"github.com/alexrs95/concerto/setlist"
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
	for _, e := range s {
		list, err := setlist.GetSongList(e)
		// if no error
		if err == nil {
			fmt.Println("----" + e)
			fmt.Println(list)
		}
	}
}
