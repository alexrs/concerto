package io

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

//ReadLines returns a slice with the lines of a given file
func ReadLines(path string) ([]string, error) {
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

// ReadContent returns the content of a ReaderCloser
func ReadContent(r io.ReadCloser) ([]byte, error) {
	content, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return content, nil
}
