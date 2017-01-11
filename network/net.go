package network

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexrs95/concerto/io"
)

//PerformRequest performs a raquest to a given url and returns the content of the page as a string
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
