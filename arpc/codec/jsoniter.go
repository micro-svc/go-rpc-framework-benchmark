package codec

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.Config{
	EscapeHTML:             true,
	SortMapKeys:            true,
	ValidateJsonRawMessage: true,
	UseNumber:              true,
}.Froze()

type jsoniterCodec struct{}

func (c jsoniterCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (c jsoniterCodec) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
