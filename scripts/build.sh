#!/usr/bin/env sh

set -e

for name in gin go-micro-v2 grpc rpcx-v5 twirp-v7
do
  cd $name
  echo ">>> $name"
  go build -tags "quic" -ldflags "-w -s" -o ../bin/$name-srv server/*.go
  go build -tags "quic" -ldflags "-w -s" -o ../bin/$name-cli client/*.go
  cd -
done
