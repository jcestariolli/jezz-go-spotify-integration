package model

type Image struct {
	URL    URL `json:"url"`
	Height int `json:"height,omitempty"`
	Width  int `json:"width,omitempty"`
}
