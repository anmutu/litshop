#!/bin/bash
dir=$GOPATH
echo $dir

go build -o genrepository ./cmd/main.go

mv genrepository $dir/bin

echo "install success"