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

func (c jsoniterCodec) Encode(i interface{}) ([]byte, error) {
	return json.Marshal(i)
}

func (c jsoniterCodec) Decode(data []byte, i interface{}) error {
	return json.Unmarshal(data, i)
}
