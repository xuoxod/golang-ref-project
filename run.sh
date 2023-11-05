#!/usr/bin/env bash

go build -o lab ./cmd/*.go
go mod tidy

clear

./lab
