package model

import (
	"strings"

	"github.com/samber/lo"
)

type TracksIDs []ID

func (a TracksIDs) String() string {
	return strings.Join(lo.Map(a, func(trackID ID, _ int) string {
		return trackID.String()
	}), ",")
}

type SimplifiedTrack struct {
	Artists          []SimplifiedArtist `json:"artists"`
	AvailableMarkets []AvailableMarket  `json:"available_markets"`
	DiscNumber       int                `json:"disc_number"`
	DurationMs       int                `json:"duration_ms"`
	Explicit         bool               `json:"explicit"`
	ExternalURLs     ExternalURLs       `json:"external_ur_ls"`
	Href             Href               `json:"href"`
	ID               ID                 `json:"id"`
	IsPlayable       bool               `json:"is_playable"`
	LinkedFrom       LinkedFrom         `json:"linked_from"`
	Restrictions     Restrictions       `json:"restrictions"`
	Name             Name               `json:"name"`
	TrackNumber      int                `json:"track_number"`
	Type             Type               `json:"type"`
	URI              URI                `json:"uri"`
	IsLocal          bool               `json:"is_local"`
}

type Track struct {
	SimplifiedTrack
	Album       SimplifiedAlbum `json:"album"`
	ExternalIDs ExternalIDs     `json:"external_i_ds"`
	Popularity  int             `json:"popularity"`
}

type MultipleTracks struct {
	Tracks []Track `json:"tracks"`
}

type SimplifiedTracksPaginated struct {
	Pagination
	Items []SimplifiedTrack `json:"items"`
}
