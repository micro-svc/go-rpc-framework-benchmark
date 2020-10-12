module github.com/micro-svc/go-rpc-framework-benchmark/grpc

go 1.14

require (
	github.com/gxxgle/go-utils v1.2.4
	github.com/micro-svc/go-rpc-framework-benchmark/model v0.0.0
	github.com/montanaflynn/stats v0.6.3
	github.com/phuslu/log v1.0.44
	google.golang.org/grpc v1.33.0
)

replace github.com/micro-svc/go-rpc-framework-benchmark/model => ../model
