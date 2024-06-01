#!/bin/bash
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/gocure gocure.go
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/gocureAPI api/api.go
docker build -t rodrigoodhin/gocure:latest .