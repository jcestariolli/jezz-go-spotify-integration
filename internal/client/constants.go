package client

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
)

const (
	ContentTypeJSON = "application/json"
)

var (
	httpNewRequest = http.NewRequest
	ioReadAll      = io.ReadAll
	jsonUnmarshal  = json.Unmarshal
	reflectValueOf = reflect.ValueOf
)
