package model

import "strconv"

type Limit int

func (l Limit) Int() int {
	return int(l)
}

func (l Limit) String() string {
	return strconv.Itoa(int(l))
}

type Offset int

func (l Offset) Int() int {
	return int(l)
}

func (l Offset) String() string {
	return strconv.Itoa(int(l))
}

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
