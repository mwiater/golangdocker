#!/bin/bash
#
# From root of project, run: `bash scripts/docker_build.sh`

clear

if [ ! -f ../.env ]
then
  export $(cat .env | xargs)
fi

# Color Console Output
RESET='\033[0m'           # Text Reset
REDBOLD='\033[1;31m'      # Red (Bold)
GREENBOLD='\033[1;32m'    # Green (Bold)
YELLOWBOLD='\033[1;33m'   # Yellow (Bold)
CYANBOLD='\033[1;36m'     # Cyan (Bold)

if [ "$DOCKERIMAGE" = "" ]; then
  echo ""
  echo -e "${REDBOLD}Please set your DOCKERIMAGE environment variable in the .env file: ${RESET} E.g.: ${CYANBOLD}DOCKERIMAGE=mattwiater/golangdocker${RESET}"
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

echo -e "${CYANBOLD}Building Docker container:${RESET} ${DOCKERIMAGE}${RESET}"
docker build -t $DOCKERIMAGE .
echo -e "${GREENBOLD}...Complete.${RESET}"
echo ""