#!/bin/bash
#
# From root of project, run: `bash docker_build.sh`

clear

for i in *.go **/*.go ; do
  gofmt -w "$i"
  echo "Formatted: $i"
done;

docker build -t mattwiater/golangdocker .