package model

type Image struct {
	Url    Url `json:"url"`
	Height int `json:"height,omitempty"`
	Width  int `json:"width,omitempty"`
}
