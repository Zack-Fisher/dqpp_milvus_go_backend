#!/bin/bash

go mod init example.com/m/v2
go mod tidy

go build -o bin/host main.go