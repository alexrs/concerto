package lib

import (
	"encoding/json"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/antzucaro/matchr"
	"log"
	"net/url"
	"strconv"
	"strings"
)

type Artist struct {
	Mbid      string `json:"@mbid"`
	Name      string `json:"@name"`
	ShortName string `json:"@shortName"`
	Url       string `json:"url"`
}

type Response struct {
	ArtistList ArtistList `json:"artists"`
}

type ArtistList struct {
	ItemsPerPage string   `json:"@itemsPerPage"`
	Page         string   `json:"@page"`
	Total        string   `json:"@total"`
	Artist       []Artist `json:"artist"`
}

const SearchURL string = "http://api.setlist.fm/rest/0.1/search/artists.json?artistName="

// get a list of songs and returns a map. The key is the song name and the value is the number of times
// the song has been played in a concert
func GetSongList(s string) (map[string]int, error) {
	var resp Response
	data, err := performRequest(SearchURL + url.QueryEscape(s))
	if err != nil {
		return nil, err
	}
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
		dist = append(dist, editDistance(e.Name, name))
	}
	min, err := min(dist)
	if err != nil {
		log.Println(err)
		return Artist{}, err
	}
	return artists[min], nil
}

func editDistance(s1, s2 string) int {
	return matchr.Levenshtein(s1, s2)
}

func min(args []int) (int, error) {
	if len(args) == 0 {
		return 0, errors.New("empty list")
	}
	min := args[0]
	minIndex := 0
	for i, e := range args {
		if e < min {
			min = e
			minIndex = i
		}
	}
	return minIndex, nil
}
