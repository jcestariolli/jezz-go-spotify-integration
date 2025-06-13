package model

type LinkedFrom struct {
	ExternalURLs ExternalURLs `json:"external_ur_ls"`
	Href         Href         `json:"href"`
	ID           string       `json:"id"`
	Type         string       `json:"type"`
	URI          URI          `json:"uri"`
}
