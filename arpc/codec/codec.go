package codec

import (
	"errors"

	"github.com/lesismal/arpc/codec"
)

type Codec = codec.Codec

var (
	Json     Codec = &codec.JSONCodec{}
	Jsoniter Codec = jsoniterCodec{}
	Protobuf Codec = protoCodec{}

	ErrInvalidMessage = errors.New("invalid message")
)
