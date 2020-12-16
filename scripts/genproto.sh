#!/bin/bash

PATH=$PATH:$GOPATH/bin

proto_dir=../../pb
pb_dir=../../pb

protoc --go_out=plugins=grpc:genproto -I ${pb_dir} ${proto_dir}/*.proto
