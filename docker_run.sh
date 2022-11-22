#!/bin/bash
#
# From root of project, run: `bash docker_run.sh`

clear

for i in *.go **/*.go ; do
  gofmt -w "$i"
  echo "Formatted: $i"
done;

docker run -it -p 5000:5000 --rm mattwiater/golangdocker