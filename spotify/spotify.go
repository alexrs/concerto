package spotify

import (
	"github.com/zmb3/spotify"
)

func init() {
	keys := getSpotifyKeys()
	auth := spotify.NewAuthenticator("", spotify.ScopeUserReadPrivate)
	auth.SetAuthInfo(keys.ClientID, keys.Secret)
}

func searchSong() (string, error) {
	return "", nil
}
