#!/usr/bin/env sh

set -e

# standard
protoc --go_out=paths=source_relative:model/standard model/model.proto

# grpc
protoc --go_out=paths=source_relative:model/grpc --go-grpc_out=paths=source_relative:model/grpc model/model.proto

# go-micro/v2
protoc --go_out=paths=source_relative:model/micro-v2 --micro_out=paths=source_relative:model/micro-v2 model/model.proto

# twirp/v7
protoc --go_out=paths=source_relative:model/twirp-v7 --twirp_out=paths=source_relative:model/twirp-v7 model/model.proto
