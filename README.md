# Concerto
Concerto is a command line tool that creates a playlist in Spotify with the top songs
artists use to play in concerts.

This songs are obtained from [setlist.fm](http://setlist.fm)

## Install
```BASH
go get github.com/alexrs95/concerto/cmd/concerto
```

## Execute

```BASH
# Create an application on Spotify: https://developer.spotify.com/my-applications
export SPOTIFY_ID=[your Spotify id]
export SPOTIFY_SECRET=[your Spotify secret]
export GODEBUG=http2client=0 #this is neccesary due to a bug on nginx. See https://github.com/zmb3/spotify/issues/20 and https://github.com/spotify/web-api/issues/398
concerto fileWithArtists playlistName
```