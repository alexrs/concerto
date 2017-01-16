package concerto

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"

	"github.com/zmb3/spotify"
)

const redirectURI = "http://localhost:8080/callback"

var (
	auth = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadPrivate,
		spotify.ScopePlaylistReadPrivate, spotify.ScopePlaylistModifyPublic,
		spotify.ScopePlaylistModifyPrivate, spotify.ScopePlaylistReadCollaborative)
	ch    = make(chan *spotify.Client)
	state = fmt.Sprintf("%x", md5.New().Sum(nil))
)

func doAuth() *spotify.Client {
	startServer()
	url := auth.AuthURL(state)
	fmt.Println("Please, log in into spotify by visiting: ", url)
	client := <-ch
	user, err := client.CurrentUser()
	if err != nil {
		log.Fatalf("Error obtaining current user: %v", err)
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

	// error handling
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	// get a new client
	client := auth.NewClient(tok)
	// print confirmation message in the browser
	fmt.Fprintf(w, "Login completed!")
	// send the client to the channel
	ch <- &client
}
