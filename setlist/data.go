package setlist

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
