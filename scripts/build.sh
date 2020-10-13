#!/usr/bin/env sh

set -e

file="./scripts/build.packages"

if [ ! -f "$file" ]
then
  echo "$file not found"
  exit 1
fi

while IFS="|" read -r name
do
  cd $name
  echo ">>> $name"
  go build -tags "quic" -ldflags "-w -s" -o ../bin/$name-srv server/*.go
  go build -tags "quic" -ldflags "-w -s" -o ../bin/$name-cli client/*.go
  cd -
done < "$file"
