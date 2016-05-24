package network

import (
	"fmt"
	"github.com/alexrs95/concerto/io"
	"log"
	"net/http"
)

func PerformRequest(url string) (string, error) {
	fmt.Println(url)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return "", err
	}
	content, err := io.ReadContent(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(content), nil
}
