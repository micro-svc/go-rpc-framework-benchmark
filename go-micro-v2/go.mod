module github.com/micro-svc/go-rpc-framework-benchmark/go-micro-v2

go 1.14

require (
	github.com/gxxgle/go-utils v1.2.4
	github.com/micro-svc/go-rpc-framework-benchmark/model v0.0.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/montanaflynn/stats v0.6.3
	github.com/phuslu/log v1.0.44
)

replace github.com/micro-svc/go-rpc-framework-benchmark/model => ../model

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
