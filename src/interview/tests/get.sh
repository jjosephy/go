#!/bin/bash

if [ $# -eq 0 ]
then
    echo "No arguments supplied"
    exit
fi

curl --header "api-version: 1.0" http://localhost:8080/interview?id=$1
