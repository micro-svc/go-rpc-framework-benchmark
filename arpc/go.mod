module github.com/micro-svc/go-rpc-framework-benchmark/arpc

go 1.14

require (
	github.com/golang/protobuf v1.4.2
	github.com/gxxgle/go-utils v1.2.4
	github.com/json-iterator/go v1.1.10
	github.com/lesismal/arpc v0.0.0-20201013181910-f4499e7405f3
	github.com/lesismal/arpcext v0.0.0-20201012120830-1c9688695667 // indirect
	github.com/micro-svc/go-rpc-framework-benchmark/model v0.0.0
	github.com/montanaflynn/stats v0.6.3
	github.com/phuslu/log v1.0.44
)

replace github.com/micro-svc/go-rpc-framework-benchmark/model => ../model

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
