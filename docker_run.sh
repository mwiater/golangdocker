#!/bin/bash
#
# From root of project, run: `bash docker_run.sh '{your-docker-hub-account-username}/{your-docker-hub-image-name}{:optional-version}'`
# E.g.: `bash docker_run.sh 'mattwiater/golangdocker'` or `bash docker_run.sh 'mattwiater/golangdocker:v1'`

clear

if [ -z "$*" ]
then
  echo ""
  echo "You must supply an image tag argument to run the docker image {your-docker-hub-account-username}/{your-docker-hub-image-name}{:optional-version}. E.g.: bash docker_run.sh 'mattwiater/golangdocker:v1'";
  echo ""
  exit 0
fi

for i in *.go **/*.go ; do
  gofmt -w "$i"
  echo "Formatted: $i"
done;

docker run -it -p 5000:5000 --rm $1