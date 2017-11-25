#! /bin/bash

set -e

go build
GOOS=windows go build
