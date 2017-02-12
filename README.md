# Concerto
Concerto is a command line tool that creates a playlist in Spotify with the top songs
artists usually plays in concerts.

This songs are obtained from [setlist.fm](http://setlist.fm)

## Install
```BASH
go get -u github.com/alexrs95/concerto/cmd/concerto
```

## Execute

```BASH
# Create an application on Spotify: https://developer.spotify.com/my-applications
export SPOTIFY_ID=[your Spotify id]
export SPOTIFY_SECRET=[your Spotify secret]
concerto fileWithArtists playlistName

#Example
 SPOTIFY_ID=[your Spotify id] SPOTIFY_SECRET=[yout Spotify secret] concerto testdata/test.txt test
```