#!/bin/bash

cd scripts
PATH=$PATH:$GOPATH/bin

proto_dir=../proto
pb_dir=../

protoc --go_out=${pb_dir} --go-grpc_out=${pb_dir} -I ${proto_dir} ${proto_dir}/*.proto
