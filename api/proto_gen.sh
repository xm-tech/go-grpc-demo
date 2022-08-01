#!/bin/bash

cmd=/usr/local/bin/protoc

protofiles=(
  helloworld/helloworld.proto
)

for proto in "${protofiles[@]}"
do
  echo "proto gen: $proto"
  $cmd --go_out=../api --go_opt=paths=source_relative --go-grpc_out=../api --go-grpc_opt=paths=source_relative "$proto"
done

