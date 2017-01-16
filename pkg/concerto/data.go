package concerto

// Artist contains the data of an artist
type Artist struct {
	Mbid      string `json:"@mbid"`
	Name      string `json:"@name"`
	ShortName string `json:"@shortName"`
	URL       string `json:"url"`
}

// ResponseList contains a list of artists
type ResponseList struct {
	ArtistList ArtistList `json:"artists"`
}

// ResponseObject contains an ArtistObject
type ResponseObject struct {
	ArtistObject ArtistObject `json:"artists"`
}

//ArtistObject is the data contained in the response with one artist
type ArtistObject struct {
	ItemsPerPage string `json:"@itemsPerPage"`
	Page         string `json:"@page"`
	Total        string `json:"@total"`
	Artists      Artist `json:"artist"`
}

// ArtistList is the data contained in a response that returns a list of artists
type ArtistList struct {
	ItemsPerPage string   `json:"@itemsPerPage"`
	Page         string   `json:"@page"`
	Total        string   `json:"@total"`
	Artists      []Artist `json:"artist"`
}

// SongStats contains the statistics of a songs
type SongStats struct {
	Count int
	Name  string
}

// SearchURL is the URL for searching artists by name
const SearchURL string = "http://api.setlist.fm/rest/0.1/search/artists.json?artistName="
