#!/usr/bin/env sh

set -e

file="./scripts/build.list"

if [ ! -f "$file" ]; then
  echo "$file not found"
  exit 1
fi

_gomod_tidy() {
  while IFS="|" read -r name; do
    echo "go mod tidy $name"
    cd $name
    rm -f go.sum && go mod tidy
    cd ..
  done <"$file"
}

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
    go build -tags "quic utp" -ldflags "-w -s" -o ../bin/$name-srv server/*.go
    go build -tags "quic utp" -ldflags "-w -s" -o ../bin/$name-cli client/*.go
    cd ..
  done <"$file"
}

case "$1" in
gomod-tidy)
  _gomod_tidy
  ;;
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
