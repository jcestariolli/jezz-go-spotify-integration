package model

type ID string

func (i ID) String() string {
	return string(i)
}

type Name string

type Type string
