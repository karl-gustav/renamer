#! /bin/bash

set -e

go build -o jenkins-cli
GOOS=windows go build -o jenkins-cli.exe
