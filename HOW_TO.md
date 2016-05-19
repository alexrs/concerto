#SetList
How to get the top 20 songs of an artist

## Statistics
http://www.setlist.fm/stats/<name-id>.html
ex:
http://www.setlist.fm/stats/metallica-3bd680c8.html

##Get <name-id>
API: `http://api.setlist.fm/rest/0.1/search/artists.json?artistName=Metallica`

It returns a list of artists that matches the search criteria.
The field url contains the url `http://www.setlist.fm/setlists/metallica-3bd680c8.html`.
Change `setlist` with `stats`.

##Select the correct group
As the API returns the results of all entries that contains the artistName, we should use some
distance measurement such as [Edit distance](https://en.wikipedia.org/wiki/Edit_distance).
 - [Wagner–Fischer algorithm](https://en.wikipedia.org/wiki/Wagner%E2%80%93Fischer_algorithm)
 - [Levenshtein distance](https://en.wikipedia.org/wiki/Levenshtein_distance)

##Get the top songs
- In the first version, the first 20 songs of the last year will be returned
- In the future, think about an algorithm to select the top songs of the last years

#Parse stats page
table class= statsTable > tbody > tr class=songRow
        td class=songName > span > a
        td class=songCount > span > span



#Spotify
##Search
Use the sarch api https://developer.spotify.com/web-api/search-item/ to get the ID of the song

##Playlist
- Create a new playlist - https://developer.spotify.com/web-api/create-playlist/ 
(First, it's necessary to have a logged user. Let's see how later)
- Add tracks to a playlist - https://developer.spotify.com/web-api/add-tracks-to-playlist/

