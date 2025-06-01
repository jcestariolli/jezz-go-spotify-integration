package model

type Href string

type Url string

type Uri string

type Genres []string

type ExternalUrls struct {
	Spotify string `json:"spotify"`
}

type Followers struct {
	Href  Href `json:"href,omitempty"`
	Total int  `json:"total"`
}

type Image struct {
	Url    Url `json:"url"`
	Height int `json:"height,omitempty"`
	Width  int `json:"width,omitempty"`
}

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
