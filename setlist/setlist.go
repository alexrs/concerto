package setlist

import (
	"encoding/json"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/alexrs95/concerto/network"
	"github.com/alexrs95/concerto/strop"
	"log"
	"net/url"
	"strconv"
	"strings"
)

// get a list of songs and returns a map. The key is the song name and the value is the number of times
// the song has been played in a concert
func GetSongList(s string) (map[string]int, error) {
	data, err := network.PerformRequest(SearchURL + url.QueryEscape(s))
	if err != nil {
		return nil, err
	}
	var resp Response
	json.Unmarshal([]byte(data), &resp)
	artist, err := getMostSimilarArtist(s, resp.ArtistList.Artist)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	page := strings.Replace(artist.Url, "setlists", "stats", -1)
	m, err := findSongsInPage(page)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return m, nil
}

// returns a map whose key is the song title and the value is how many times
// it has been played in the last concerts
func findSongsInPage(page string) (map[string]int, error) {
	m := make(map[string]int)
	doc, err := goquery.NewDocument(page)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	doc.Find(".songRow").Each(func(i int, s *goquery.Selection) {
		song := s.Find(".songName a").First().Text()
		count, _ := strconv.Atoi(s.Find(".songCount span span").Text())
		m[song] = count
	})
	if len(m) == 0 {
		return nil, errors.New("empty map - no songs")
	}
	return m, nil
}

// returns the most similar artist or an error
func getMostSimilarArtist(name string, artists []Artist) (Artist, error) {
	var dist []int
	for _, e := range artists {
		dist = append(dist, strop.EditDistance(e.Name, name))
	}
	min, err := strop.MinIndex(dist)
	if err != nil {
		log.Println(err)
		return Artist{}, err
	}
	if dist[min] > 10 {
		return Artist{}, errors.New("No artist with similar name")
	}
	return artists[min], nil
}
