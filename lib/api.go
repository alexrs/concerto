package lib

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func performRequest(url string) (string, error) {
	fmt.Println(url)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return "", err
	}
	content, _ := readContent(resp.Body)
	return string(content), nil
}

func readContent(r io.ReadCloser) ([]byte, error) {
	content, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return content, nil
}
