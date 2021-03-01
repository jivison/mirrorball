package indexing

import (
	"github.com/jivison/gowon-indexer/lib/customerrors"
	"github.com/jivison/gowon-indexer/lib/db"
	"github.com/jivison/gowon-indexer/lib/graph/model"
)

// GetArtist gets and optionally creates an indexed artist
func (i Indexing) GetArtist(artistInput model.ArtistInput, create bool) (*db.Artist, error) {
	artist := new(db.Artist)

	query := db.Db.Model(artist)

	err := ParseArtistInput(query, artistInput).Limit(1).Select()

	if err != nil && create == true && artistInput.Name != nil {
		artist = &db.Artist{
			Name: *artistInput.Name,
		}

		db.Db.Model(artist).Insert()
	} else if err != nil {
		return nil, customerrors.EntityDoesntExistError("artist")
	}

	return artist, nil
}

// GetAlbum returns (and optionally creates) an album from the database
func (i Indexing) GetAlbum(albumInput model.AlbumInput, create bool) (*db.Album, error) {
	album := new(db.Album)

	query := db.Db.Model(album).
		Relation("Artist")

	err := ParseAlbumInput(query, albumInput).Limit(1).Select()

	if err != nil && create == true && albumInput.Name != nil && albumInput.SafeGetArtistName() != nil {
		artist, _ := i.GetArtist(*albumInput.Artist, true)

		album = &db.Album{
			Name: *albumInput.Name,

			ArtistID: artist.ID,
			Artist:   artist,
		}

		db.Db.Model(album).Insert()
	} else if err != nil {
		return nil, customerrors.EntityDoesntExistError("album")
	}

	return album, nil
}

// GetTrack returns (and optionally creates) a track from the database
func (i Indexing) GetTrack(trackInput model.TrackInput, create bool) (*db.Track, error) {
	track := new(db.Track)

	query := db.Db.Model(track).
		Relation("Artist").
		Relation("Album")

	err := ParseTrackInput(query, trackInput).Limit(1).Select()

	if err != nil && create == true && trackInput.Name != nil && trackInput.SafeGetArtistName() != nil {
		track = i.SaveTrack(*trackInput.Name, *trackInput.SafeGetArtistName(), trackInput.SafeGetAlbumName())
	}

	return track, nil
}

// GetArtistCount gets and optionally creates an artist count
func (i Indexing) GetArtistCount(artist *db.Artist, user *db.User, create bool) (*db.ArtistCount, error) {
	artistCount := new(db.ArtistCount)

	err := db.Db.Model(artistCount).Where("user_id=?", user.ID).Where("artist_id=?", artist.ID).Limit(1).Select()

	if err != nil && create == true {
		artistCount = &db.ArtistCount{
			UserID: user.ID,
			User:   user,

			ArtistID: artist.ID,
			Artist:   artist,
		}

		db.Db.Model(artistCount).Insert()
	} else if err != nil {
		return nil, customerrors.EntityDoesntExistError("artist count")
	}

	return artistCount, nil
}

// GetAlbumCount gets and optionally creates an album count
func (i Indexing) GetAlbumCount(album *db.Album, user *db.User, create bool) (*db.AlbumCount, error) {
	albumCount := new(db.AlbumCount)

	err := db.Db.Model(albumCount).Where("user_id=?", user.ID).Where("album_id=?", album.ID).Limit(1).Select()

	if err != nil && create == true {
		albumCount = &db.AlbumCount{
			UserID: user.ID,
			User:   user,

			AlbumID: album.ID,
			Album:   album,
		}

		db.Db.Model(albumCount).Insert()
	} else if err != nil {
		return nil, customerrors.EntityDoesntExistError("album count")
	}

	return albumCount, nil
}

// GetTrackCount gets and optionally creates an track count
func (i Indexing) GetTrackCount(track *db.Track, user *db.User, create bool) (*db.TrackCount, error) {
	trackCount := new(db.TrackCount)

	err := db.Db.Model(trackCount).Where("user_id=?", user.ID).Where("track_id=?", track.ID).Limit(1).Select()

	if err != nil && create == true {
		trackCount = &db.TrackCount{
			UserID: user.ID,
			User:   user,

			TrackID: track.ID,
			Track:   track,
		}

		db.Db.Model(trackCount).Insert()
	} else if err != nil {
		return nil, customerrors.EntityDoesntExistError("track count")
	}

	return trackCount, nil
}

// SaveTrack saves a track in the database
func (i Indexing) SaveTrack(trackName, artistName string, albumName *string) *db.Track {

	artist, _ := i.GetArtist(model.ArtistInput{Name: &artistName}, true)
	var album *db.Album = nil

	if albumName != nil {
		album, _ = i.GetAlbum(model.AlbumInput{
			Name:   albumName,
			Artist: &model.ArtistInput{Name: &artistName},
		}, true)
	}

	track := &db.Track{
		Name: trackName,
	}

	if album != nil {
		track.Album = album
		track.AlbumID = &album.ID
	}

	if artist != nil {
		track.Artist = artist
		track.ArtistID = artist.ID
	}

	db.Db.Model(track).Insert()

	return track
}