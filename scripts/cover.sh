#!/bin/bash
go test -coverprofile=cover.out -covermode=atomic -v ./...
go tool cover -html=cover.out -o cover.html
if [ $? -eq 0 ]
then
  exit 0
else
  exit 1
fi
