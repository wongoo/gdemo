#!/bin/sh

for file in `ls *.go` ; do
    echo "================================="
    echo "-------> go run $file"
    go run $file
done