#! /usr/bin/bash

go build -o lab ./*.go

clear

if [ $# == 2 ];
then
    ./lab -argument=$1 -action=$2
else
    printf "\n\tProgram Ended\n\n"
fi