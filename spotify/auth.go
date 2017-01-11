package spotify

import (
	"fmt"
	"log"
	"net/http"

	"os"

	"github.com/zmb3/spotify"
)

const redirectURI = "http://localhost:8080/callback"

var (
	auth = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadPrivate,
		spotify.ScopePlaylistReadPrivate, spotify.ScopePlaylistModifyPublic,
		spotify.ScopePlaylistModifyPrivate, spotify.ScopePlaylistReadCollaborative)
	ch    = make(chan *spotify.Client)
	state = "abc123"
)

func doAuth() *spotify.Client {
	startServer()
	auth.SetAuthInfo(os.Getenv("CONCERTO_CLIENT_ID"), os.Getenv("CONCERTO_SECRET"))
	url := auth.AuthURL(state)
	fmt.Println("Please, log in into spotify by visiting: ", url)
	client := <-ch
	user, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You are logged in as:", user.ID)
	return client
}

func startServer() {
	//configure http server
	http.HandleFunc("/callback", completeAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	// start the server in a new goroutine
	go http.ListenAndServe(":8080", nil)
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	client := auth.NewClient(tok)
	fmt.Fprintf(w, "Login completed!")
	ch <- &client
}
