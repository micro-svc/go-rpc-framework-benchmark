#!/usr/bin/env sh

set -e

for name in go-micro-v2 grpc rpcx-v5 twirp-v7
do
  cd $name
  echo ">>> $name"
  golangci-lint run
  cd -
done
