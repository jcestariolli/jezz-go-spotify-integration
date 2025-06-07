package model

type Artist struct {
	ExternalUrls ExternalUrls `json:"external_urls"`
	Followers    Followers    `json:"followers"`
	Genres       Genres       `json:"genres"`
	Href         Href         `json:"href"`
	Id           string       `json:"id"`
	Images       []Image      `json:"images"`
	Name         string       `json:"name"`
	Popularity   int          `json:"popularity"`
	Type         string       `json:"type"`
	Uri          Uri          `json:"uri"`
}

type Artists []Artist

type MultipleArtists struct {
	Artists Artists `json:"artists"`
}
