package codec

import (
	// nolint
	"github.com/golang/protobuf/proto"
)

type protoCodec struct{}

func (c protoCodec) Marshal(v interface{}) ([]byte, error) {
	p, ok := v.(proto.Message)
	if !ok {
		return nil, ErrInvalidMessage
	}
	return proto.Marshal(p)
}

func (c protoCodec) Unmarshal(data []byte, v interface{}) error {
	p, ok := v.(proto.Message)
	if !ok {
		return ErrInvalidMessage
	}
	return proto.Unmarshal(data, p)
}
