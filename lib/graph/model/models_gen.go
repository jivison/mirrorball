// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/jivison/gowon-indexer/lib/meta"
)

type Album struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Artist *Artist  `json:"artist"`
	Tracks []*Track `json:"tracks"`
}

type AlbumCount struct {
	Album     *Album `json:"album"`
	Playcount int    `json:"playcount"`
}

type AlbumInput struct {
	Artist *ArtistInput `json:"artist"`
	Name   *string      `json:"name"`
}

type AlbumPlaysSettings struct {
	PageInput *PageInput  `json:"pageInput"`
	Album     *AlbumInput `json:"album"`
	Sort      *string     `json:"sort"`
}

type AlbumTopTracksResponse struct {
	Album     *Album                 `json:"album"`
	TopTracks []*AmbiguousTrackCount `json:"topTracks"`
}

type AmbiguousTrack struct {
	Name   string   `json:"name"`
	Artist string   `json:"artist"`
	Albums []*Album `json:"albums"`
}

type AmbiguousTrackCount struct {
	Name      string `json:"name"`
	Playcount int    `json:"playcount"`
}

type Artist struct {
	ID   int      `json:"id"`
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

type ArtistCount struct {
	Artist    *Artist `json:"artist"`
	Playcount int     `json:"playcount"`
	User      *User   `json:"user"`
}

type ArtistInput struct {
	Name *string `json:"name"`
}

type ArtistPlaysSettings struct {
	PageInput *PageInput   `json:"pageInput"`
	Artist    *ArtistInput `json:"artist"`
	Sort      *string      `json:"sort"`
}

type ArtistRankResponse struct {
	Artist    *Artist      `json:"artist"`
	Rank      int          `json:"rank"`
	Playcount int          `json:"playcount"`
	Listeners int          `json:"listeners"`
	Above     *ArtistCount `json:"above"`
	Below     *ArtistCount `json:"below"`
}

type ArtistSearchCriteria struct {
	Keywords *string `json:"keywords"`
}

type ArtistSearchResult struct {
	ArtistID        int    `json:"artistID"`
	ArtistName      string `json:"artistName"`
	ListenerCount   int    `json:"listenerCount"`
	GlobalPlaycount int    `json:"globalPlaycount"`
}

type ArtistSearchResults struct {
	Artists []*ArtistSearchResult `json:"artists"`
}

type ArtistTopAlbumsResponse struct {
	Artist    *Artist       `json:"artist"`
	TopAlbums []*AlbumCount `json:"topAlbums"`
}

type ArtistTopTracksResponse struct {
	Artist    *Artist                `json:"artist"`
	TopTracks []*AmbiguousTrackCount `json:"topTracks"`
}

type GuildMember struct {
	UserID  int    `json:"userID"`
	GuildID string `json:"guildID"`
	User    *User  `json:"user"`
}

type PageInfo struct {
	RecordCount int `json:"recordCount"`
}

type PageInput struct {
	Limit  *int `json:"limit"`
	Offset *int `json:"offset"`
}

type Play struct {
	ID          int    `json:"id"`
	ScrobbledAt int    `json:"scrobbledAt"`
	User        *User  `json:"user"`
	Track       *Track `json:"track"`
}

type PlaysInput struct {
	User      *UserInput  `json:"user"`
	Track     *TrackInput `json:"track"`
	Sort      *string     `json:"sort"`
	Timerange *Timerange  `json:"timerange"`
}

type PlaysResponse struct {
	Plays    []*Play   `json:"plays"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type RateYourMusicAlbum struct {
	RateYourMusicID  string  `json:"rateYourMusicID"`
	Title            string  `json:"title"`
	ReleaseYear      *int    `json:"releaseYear"`
	ArtistName       string  `json:"artistName"`
	ArtistNativeName *string `json:"artistNativeName"`
}

type RateYourMusicArtist struct {
	ArtistName       string  `json:"artistName"`
	ArtistNativeName *string `json:"artistNativeName"`
}

type Rating struct {
	RateYourMusicAlbum *RateYourMusicAlbum `json:"rateYourMusicAlbum"`
	Rating             int                 `json:"rating"`
}

type RatingsResponse struct {
	Ratings  []*Rating `json:"ratings"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type RatingsSettings struct {
	User      *UserInput  `json:"user"`
	Album     *AlbumInput `json:"album"`
	PageInput *PageInput  `json:"pageInput"`
	Rating    *int        `json:"rating"`
}

type SearchSettings struct {
	Exact *bool      `json:"exact"`
	User  *UserInput `json:"user"`
}

type Tag struct {
	Name        string `json:"name"`
	Occurrences int    `json:"occurrences"`
}

type TagInput struct {
	Name *string `json:"name"`
}

type TagsResponse struct {
	Tags     []*Tag    `json:"tags"`
	PageInfo *PageInfo `json:"pageInfo"`
}

type TagsSettings struct {
	Artists   []*ArtistInput `json:"artists"`
	Keyword   *string        `json:"keyword"`
	PageInput *PageInput     `json:"pageInput"`
}

type TaskStartResponse struct {
	TaskName string `json:"taskName"`
	Success  bool   `json:"success"`
	Token    string `json:"token"`
}

type Timerange struct {
	From *meta.Date `json:"from"`
	To   *meta.Date `json:"to"`
}

type Track struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Artist *Artist `json:"artist"`
	Album  *Album  `json:"album"`
}

type TrackCount struct {
	Track     *Track `json:"track"`
	Playcount int    `json:"playcount"`
}

type TrackInput struct {
	Artist *ArtistInput `json:"artist"`
	Album  *AlbumInput  `json:"album"`
	Name   *string      `json:"name"`
}

type TrackPlaysSettings struct {
	PageInput *PageInput  `json:"pageInput"`
	Track     *TrackInput `json:"track"`
	Sort      *string     `json:"sort"`
}

type TrackTopAlbumsResponse struct {
	Track     *AmbiguousTrack `json:"track"`
	TopAlbums []*TrackCount   `json:"topAlbums"`
}

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	DiscordID string    `json:"discordID"`
	UserType  *UserType `json:"userType"`
}

type UserInput struct {
	DiscordID      *string `json:"discordID"`
	LastFMUsername *string `json:"lastFMUsername"`
	WavyUsername   *string `json:"wavyUsername"`
}

type WhoFirstArtistResponse struct {
	Rows   []*WhoFirstRow `json:"rows"`
	Artist *Artist        `json:"artist"`
}

type WhoFirstRow struct {
	User        *User `json:"user"`
	ScrobbledAt int   `json:"scrobbledAt"`
}

type WhoKnowsAlbumResponse struct {
	Rows  []*WhoKnowsRow `json:"rows"`
	Album *Album         `json:"album"`
}

type WhoKnowsArtistResponse struct {
	Rows   []*WhoKnowsRow `json:"rows"`
	Artist *Artist        `json:"artist"`
}

type WhoKnowsRow struct {
	User      *User `json:"user"`
	Playcount int   `json:"playcount"`
}

type WhoKnowsSettings struct {
	GuildID *string `json:"guildID"`
	Limit   *int    `json:"limit"`
}

type WhoKnowsTrackResponse struct {
	Rows  []*WhoKnowsRow  `json:"rows"`
	Track *AmbiguousTrack `json:"track"`
}

type UserType string

const (
	UserTypeWavy   UserType = "Wavy"
	UserTypeLastfm UserType = "Lastfm"
)

var AllUserType = []UserType{
	UserTypeWavy,
	UserTypeLastfm,
}

func (e UserType) IsValid() bool {
	switch e {
	case UserTypeWavy, UserTypeLastfm:
		return true
	}
	return false
}

func (e UserType) String() string {
	return string(e)
}

func (e *UserType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserType", str)
	}
	return nil
}

func (e UserType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
