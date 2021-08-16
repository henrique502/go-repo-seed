#!/bin/bash
rm -rf ./release
mkdir ./release
go build -o ./release/alerts ./entrypoint/lambda-alerts/main.go
go build -o ./release/integrations ./entrypoint/lambda-integrations/main.go
go build -o ./release/teams ./entrypoint/lambda-teams/main.go
exit 0
