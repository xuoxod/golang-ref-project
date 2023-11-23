#!/usr/bin/env bash

clear

case $# in
0)
    printf "\tNo arguments\n"
    ;;

1)
    if [ $1 = "-env" ] || [ $1 = "testdbc" ] || [ $1 = "querydb" ] || [ $1 = "genhash" ] || [ $1 = "comhash" ] || [ $1 = "datesta" ] || [ $1 = "timesta" ] || [ $1 = "dtstamp" ]; then
        printf "Argument is valid\n"
    else
        printf "Argument is Not valid\n"
    fi
    ;;
esac
