#!/usr/bin/env bash

clear

if [ $# == 1 ]; then
    if [ $1 = "true" ] || [ $1 = "false" ]; then
        printf "Argument is a boolean value\n"
        if [ -n $1 ]; then
            printf "Argument is not empty\n"
        else
            printf "Argument is empty\n"
        fi
    else
        printf "Argument is not a boolean value\n"
    fi
else
    printf "No arguments\n"
fi
