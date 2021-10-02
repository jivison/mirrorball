package indexing

import (
	"github.com/jivison/gowon-indexer/lib/db"
	apihelpers "github.com/jivison/gowon-indexer/lib/helpers/api"
	helpers "github.com/jivison/gowon-indexer/lib/helpers/generic"
	"github.com/jivison/gowon-indexer/lib/meta"
	"github.com/jivison/gowon-indexer/lib/services/lastfm"
)

func (i Indexing) GenerateCountsFromScrobbles(scrobbles []lastfm.RecentTrack, user db.User) ([]db.ArtistCount, []db.AlbumCount, []db.TrackCount, []db.Play, error) {
	artistsMap, albumsMap, tracksMap, err := i.ConvertScrobbles(scrobbles)

	if err != nil {
		return nil, nil, nil, nil, err
	}

	var dbScrobbles []db.Play

	artistCounts := meta.CreateArtistConversionCounter()
	albumCounts := meta.CreateAlbumConversionCounter()
	trackCounts := meta.CreateTrackConversionCounter()

	for _, scrobble := range scrobbles {
		artist, _, _ := artistsMap.Get(scrobble.Artist.Text)
		artistCounts.Increment(scrobble.Artist.Text)

		track, _, _ := tracksMap.Get(scrobble.Artist.Text, scrobble.Album.Text, scrobble.Name)
		trackCounts.Increment(artist.Name, scrobble.Album.Text, scrobble.Name)

		if scrobble.Album.Text != "" {

			album, _, _ := albumsMap.Get(scrobble.Artist.Text, scrobble.Album.Text)
			albumCounts.Increment(artist.Name, album.Name)
		}

		timestamp, _ := apihelpers.ParseUnix(scrobble.Timestamp.UTS)
		dbScrobbles = append(dbScrobbles, db.Play{
			UserID: user.ID,
			User:   &user,

			TrackID: track.ID,
			Track:   &track,

			ScrobbledAt: timestamp,
		})
	}

	var dbArtistCounts []db.ArtistCount
	var dbAlbumCounts []db.AlbumCount
	var dbTrackCounts []db.TrackCount

	for artist, count := range artistCounts.GetMap() {
		dbArtist, _, _ := artistsMap.Get(artist)
		dbArtistCounts = append(dbArtistCounts, db.ArtistCount{Artist: &dbArtist, ArtistID: dbArtist.ID, User: &user, UserID: user.ID, Playcount: count.Value.(int32)})
	}

	for artist, artistAlbums := range albumCounts.GetMap() {
		for album, count := range artistAlbums.Value.(meta.ConversionMap).GetMap() {
			dbAlbum, _, _ := albumsMap.Get(artist, album)
			dbAlbumCounts = append(dbAlbumCounts, db.AlbumCount{
				Album: &dbAlbum, AlbumID: dbAlbum.ID, User: &user, UserID: user.ID, Playcount: count.Value.(int32),
			})
		}
	}

	for artist, artistAlbums := range trackCounts.GetMap() {
		for album, albumTracks := range artistAlbums.Value.(meta.ConversionMap).GetMap() {
			for track, count := range albumTracks.Value.(meta.ConversionMap).GetMap() {
				dbTrack, _, _ := tracksMap.Get(artist, album, track)
				dbTrackCounts = append(dbTrackCounts, db.TrackCount{Track: &dbTrack, TrackID: dbTrack.ID, User: &user, UserID: user.ID, Playcount: count.Value.(int32)})
			}
		}
	}

	return dbArtistCounts, dbAlbumCounts, dbTrackCounts, dbScrobbles, nil
}

func (i Indexing) ConvertScrobbles(scrobbles []lastfm.RecentTrack) (meta.ArtistConversionMap, meta.AlbumConversionMap, meta.TrackConversionMap, error) {
	artistsList, albumsList, tracksList := i.GenerateUniqueLists(scrobbles)

	artistsMap, err := i.ConvertArtists(artistsList)

	if err != nil {
		return artistsMap, meta.AlbumConversionMap{}, meta.TrackConversionMap{}, err
	}

	albumsMap, err := i.ConvertAlbums(albumsList, &artistsMap)

	if err != nil {
		return artistsMap, albumsMap, meta.TrackConversionMap{}, err
	}

	tracksMap, err := i.ConvertTracks(tracksList, &artistsMap, &albumsMap)

	if err != nil {
		return artistsMap, albumsMap, tracksMap, err
	}

	return artistsMap, albumsMap, tracksMap, nil
}

func (I Indexing) GenerateUniqueLists(scrobbles []lastfm.RecentTrack) ([]ArtistToConvert, []AlbumToConvert, []TrackToConvert) {
	artists := meta.CreateArtistConversionCounter()
	albums := meta.CreateAlbumConversionCounter()
	tracks := meta.CreateTrackConversionCounter()

	for _, scrobble := range scrobbles {
		artists.Increment(scrobble.Artist.Text)
		albums.Increment(scrobble.Artist.Text, scrobble.Album.Text)
		tracks.Increment(scrobble.Artist.Text, scrobble.Album.Text, scrobble.Name)
	}

	var artistsList []ArtistToConvert
	var albumsList []AlbumToConvert
	var tracksList []TrackToConvert

	for artist := range artists.GetMap() {
		artistsList = append(artistsList, artist)
	}

	for artist, artistAlbums := range albums.GetMap() {
		for album := range artistAlbums.Value.(meta.ConversionMap).GetMap() {
			albumsList = append(albumsList, AlbumToConvert{
				AlbumName:  album,
				ArtistName: artist,
			})
		}
	}

	for artist, artistAlbums := range tracks.GetMap() {
		for album, albumTracks := range artistAlbums.Value.(meta.ConversionMap).GetMap() {
			for track := range albumTracks.Value.(meta.ConversionMap).GetMap() {
				var albumName *string

				if album != "" {
					copiedAlbumName := helpers.DeepCopy(album)
					albumName = &copiedAlbumName
				}

				tracksList = append(tracksList, TrackToConvert{
					ArtistName: artist,
					TrackName:  track,
					AlbumName:  albumName,
				})
			}
		}
	}

	return artistsList, albumsList, tracksList
}
