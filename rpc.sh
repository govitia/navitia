#!/bin/bash
# This script is used to generate .pb.go files from .proto files

# go install google.golang.org/protobuf/cmd/protoc-gen-go
# https://developers.google.com/protocol-buffers/docs/reference/go-generated#invocation
protoc  --proto_path=./proto --go_out=./rpc --go_opt=paths=source_relative ./proto/*.proto