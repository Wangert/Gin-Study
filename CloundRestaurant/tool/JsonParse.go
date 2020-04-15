package tool

import (
	"io"
	"encoding/json"
)

type JsonParse struct {

}

func (jsonParse *JsonParse) Decode(io io.ReadCloser, v interface{}) error {
	return json.NewDecoder(io).Decode(v)
}

