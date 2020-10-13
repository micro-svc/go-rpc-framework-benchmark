package codec

import (
	"github.com/smallnest/rpcx/v5/protocol"
	"github.com/smallnest/rpcx/v5/share"
)

var (
	JSONiter protocol.SerializeType = 11
)

func Register() {
	share.RegisterCodec(JSONiter, &jsoniterCodec{})
}
