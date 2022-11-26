#!/bin/bash
#
# From root of project, run: `bash go_run.sh`

clear

for i in *.go **/*.go ; do
  gofmt -w "$i"
  echo "Formatted: $i"
done;

go run .