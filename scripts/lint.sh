#!/bin/bash
$(go env GOPATH)/bin/golangci-lint run ./... --issues-exit-code 1
if [ $? -eq 0 ]
then
  exit 0
else
  exit 1
fi
