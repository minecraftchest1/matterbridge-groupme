#!/bin/bash
env GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
env GOOS=linux GOARM=5 GOARCH=arm go build -o bin/main-linux-arm5 main.go
env GOOS=linux GOARM=6 GOARCH=arm go build -o bin/main-linux-arm6 main.go
env GOOS=linux GOARM=7 GOARCH=arm go build -o bin/main-linux-arm7 main.go

