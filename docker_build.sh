#!/bin/bash
#
# From root of project, run: `bash docker_build.sh '{your-docker-hub-account-username}/{your-docker-hub-image-name}{:optional-version}'`
# E.g.: `bash docker_build.sh 'mattwiater/golangdocker'` or `bash docker_build.sh 'mattwiater/golangdocker:v1'`

clear

# Color Console Output
RESET='\033[0m'           # Text Reset
REDBOLD='\033[1;31m'      # Red (Bold)
GREENBOLD='\033[1;32m'    # Green (Bold)
YELLOWBOLD='\033[1;33m'   # Yellow (Bold)
CYANBOLD='\033[1;36m'     # Cyan (Bold)

if [ -z "$*" ]
then
  echo ""
  echo -e "${REDBOLD}You must supply an image tag argument to run the docker image {your-docker-hub-account-username}/{your-docker-hub-image-name}{:optional-version}.${RESET} E.g.: bash docker_run.sh 'mattwiater/golangdocker:v1'"
  echo ""
  exit 0
fi

echo -e "${CYANBOLD}Building Swagger docs...${RESET}"
swag init
echo -e "${GREENBOLD}...Complete.${RESET}"
echo ""

echo -e "${CYANBOLD}Formatting *.go files...${RESET}"
for i in *.go **/*.go ; do
  gofmt -w "$i"
  echo "Formatted: $i"
done;
echo -e "${GREENBOLD}...Complete.${RESET}"
echo ""

echo -e "${CYANBOLD}Building Docker container:${RESET} ${GREENBOLD}$1${RESET}"
docker build -t $1 .
echo -e "${GREENBOLD}...Complete.${RESET}"
echo ""