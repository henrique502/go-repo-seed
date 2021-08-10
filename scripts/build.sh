#!/bin/bash
rm -rf ./release
mkdir ./release
go build -o ./release/app ./main.go
if [ $? -eq 0 ]
then
  exit 0
else
  exit 1
fi
