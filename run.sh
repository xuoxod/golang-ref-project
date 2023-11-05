#!/usr/bin/env bash

go build -o lab ./cmd/*.go
go mod tidy

clear

if [ $# == 1 ]; then
    if [ $1 = "true" ] || [ $1 = "false" ] && [ -n $1 ]; then
        ./lab -env=$1
    else
        ./lab
    fi
else
    ./lab
fi
