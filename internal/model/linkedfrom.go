package model

type LinkedFrom struct {
	ExternalUrls ExternalUrls `json:"external_urls"`
	Href         Href         `json:"href"`
	Id           string       `json:"id"`
	Type         string       `json:"type"`
	Uri          Uri          `json:"uri"`
}
