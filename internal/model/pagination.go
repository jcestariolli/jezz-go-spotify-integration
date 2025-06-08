package model

type Limit int

type Offset int

type Total int

type Next string

type Previous string

type Pagination struct {
	Href     Href      `json:"href"`
	Limit    Limit     `json:"limit"`
	Next     *Next     `json:"next,omitempty"`
	Offset   Offset    `json:"offset"`
	Previous *Previous `json:"previous,omitempty"`
	Total    Total     `json:"total"`
}
