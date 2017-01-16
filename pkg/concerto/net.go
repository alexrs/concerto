package concerto

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//PerformRequest performs a raquest to a given url and returns the content of the page as a string
func PerformRequest(url string) (string, error) {
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(content), nil
}
