#!/usr/bin/env sh

set -e

file="./scripts/benchmark.packages"

if [ ! -f "$file" ]
then
  echo "$file not found"
  exit 1
fi

touch go-rpc-framework-benchmark.result
echo "" > go-rpc-framework-benchmark.result

while IFS="|" read -r name package server client
do
  pm2 start "${server}" --name ${name}-benchmark
  sleep 3s
  echo "\n### ${name}\n\n\`\`\`sh\n$ ${client}" >> go-rpc-framework-benchmark.result
  ${client} >> go-rpc-framework-benchmark.result 2>&1
  echo "\`\`\`" >> go-rpc-framework-benchmark.result
  pm2 del ${name}-benchmark
  sleep 15s # wait resource release
done < "$file"
