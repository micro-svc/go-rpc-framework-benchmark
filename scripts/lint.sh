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
  golangci-lint run
  cd -
done < "$file"
