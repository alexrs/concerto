package setlist

import (
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/alexrs95/concerto/network"
	"github.com/antzucaro/matchr"
)

// ByCount orders the songs by count
type ByCount []SongStats

func (s ByCount) Len() int {
	return len(s)
}
func (s ByCount) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByCount) Less(i, j int) bool {
	return s[i].Count > s[j].Count
}

// GetSongList get a list of songs and returns a map. The key is the song name
// and the value is the number of times the song has been played in a concert
func GetSongList(s string) ([]SongStats, error) {
	data, err := network.PerformRequest(SearchURL + url.QueryEscape(s))
	if err != nil {
		return nil, err
	}
	resp := unmarshalResponse(data)
	artist, err := getMostSimilarArtist(s, resp)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	page := strings.Replace(artist.URL, "setlists", "stats", -1)
	m, err := findSongsInPage(page)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return m, nil
}

func unmarshalResponse(data string) []Artist {
	// The API returns a list of artist, but when there is only one artist
	// instead of returning a list of length 1, it returns a JSON Object. (WTF)
	if strings.Contains(data, `"artist":{"`) {
		var respObj ResponseObject
		json.Unmarshal([]byte(data), &respObj)
		return []Artist{respObj.ArtistObject.Artists}
	}
	var respList ResponseList //most common case
	json.Unmarshal([]byte(data), &respList)
	return respList.ArtistList.Artists
}

// returns a map whose key is the song title and the value is how many times
// it has been played in the last concerts
func findSongsInPage(page string) ([]SongStats, error) {
	l := []SongStats{}
	doc, err := goquery.NewDocument(page)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	doc.Find(".songRow").Each(func(i int, s *goquery.Selection) {
		song := s.Find(".songName a").First().Text()
		count, _ := strconv.Atoi(s.Find(".songCount span span").Text())
		l = append(l, SongStats{count, song})
	})
	if len(l) == 0 {
		return nil, errors.New("empty list - no songs")
	}
	sort.Sort(ByCount(l))
	return l, nil
}

// returns the most similar artist or an error
func getMostSimilarArtist(name string, artists []Artist) (Artist, error) {
	// first, we get the list of distances between the given name, and the list
	// of artists
	var dist []int
	for _, e := range artists {
		dist = append(dist, matchr.Levenshtein(e.Name, name))
	}
	// gen the index of the minimun value (faster than sorting and getting the
	// first element, this operation is O(n) and sorting is O(nlgn))
	min, err := MinIndex(dist)
	if err != nil {
		log.Println(err)
		return Artist{}, err
	}
	// of the distance is greater than some value, no similar artist found
	if dist[min] > len(name)/5 {
		return Artist{}, errors.New("No artist with similar name")
	}
	return artists[min], nil
}

// MinIndex returns the index of the smallest element of the list
func MinIndex(args []int) (int, error) {
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
