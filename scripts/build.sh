#!/usr/bin/env sh

set -e

file="./scripts/build.list"

if [ ! -f "$file" ]; then
  echo "$file not found"
  exit 1
fi

_lint() {
  while IFS="|" read -r name; do
    echo "lint $name"
    cd $name
    golangci-lint run
    cd ..
  done <"$file"
}

_build() {
  while IFS="|" read -r name; do
    echo "build $name"
    cd $name
    go build -tags "quic" -ldflags "-w -s" -o ../bin/$name-srv server/*.go
    go build -tags "quic" -ldflags "-w -s" -o ../bin/$name-cli client/*.go
    cd ..
  done <"$file"
}

case "$1" in
lint)
  _lint
  ;;
build)
  _build
  ;;
*)
  echo "illegal command: $1"
  exit 1
  ;;
esac
