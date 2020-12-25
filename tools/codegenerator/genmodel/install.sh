#!/bin/bash
dir=$GOPATH
echo $dir

go build -o genmodel ./cmd/main.go

mv genmodel $dir/bin

echo "install success"