#!/bin/bash

proto_dir="api"
go_dir="api"

for proto_file in `find $proto_dir -name "*.proto"`; do
  protoc -I=$proto_dir --go_out=plugins=grpc:$go_dir $proto_file;
done
