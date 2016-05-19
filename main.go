package main

import (
	"bufio"
	"fmt"
	"github.com/alexrs95/concerto/lib"
	"log"
	"os"
)

func main() {
	filePath := os.Args[1]
	fmt.Println(filePath)
	s, err := readLines(filePath)
	if err != nil {
		log.Println(err)
	}
	for _, e := range s {
		list, err := lib.GetSongList(e)
		// if no error
		if err == nil {
			fmt.Println("----" + e)
			fmt.Println(list)
		}
	}
}

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
	return lines, scanner.Err()
}
