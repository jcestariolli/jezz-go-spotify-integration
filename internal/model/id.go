package model

type Id string

func (i Id) String() string {
	return string(i)
}

type Name string

type Type string
