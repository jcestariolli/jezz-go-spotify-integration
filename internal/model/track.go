package model

type SimplifiedTrack struct {
	Artists          []SimplifiedArtist `json:"artists"`
	AvailableMarkets []AvailableMarket  `json:"available_markets"`
	DiscNumber       int                `json:"disc_number"`
	DurationMs       int                `json:"duration_ms"`
	Explicit         bool               `json:"explicit"`
	ExternalUrls     ExternalUrls       `json:"external_urls"`
	Href             Href               `json:"href"`
	Id               Id                 `json:"id"`
	IsPlayable       bool               `json:"is_playable"`
	LinkedFrom       LinkedFrom         `json:"linked_from"`
	Restrictions     Restrictions       `json:"restrictions"`
	Name             Name               `json:"name"`
	TrackNumber      int                `json:"track_number"`
	Type             Type               `json:"type"`
	Uri              Uri                `json:"uri"`
	IsLocal          bool               `json:"is_local"`
}

type Tracks struct {
	Href     Href              `json:"href"`
	Limit    Limit             `json:"limit"`
	Next     Next              `json:"next"`
	Offset   Offset            `json:"offset"`
	Previous Previous          `json:"previous"`
	Total    Total             `json:"total"`
	Items    []SimplifiedTrack `json:"items"`
}
