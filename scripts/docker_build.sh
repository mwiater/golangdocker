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
ERROR=$(swag init 2>&1 1>/dev/null)
status=$?
if test $status -ne 0
then
	echo -e "${REDBOLD}...Error: 'swag init' command failed:${RESET}"
  echo -e "${REDBOLD}ERROR: ${ERROR}${RESET}"
  echo ""
  exit 1
fi
echo -e "${GREENBOLD}...Complete.${RESET}"
echo ""

echo -e "${CYANBOLD}Formatting *.go files...${RESET}"
for i in *.go **/*.go ; do
  ERROR=$(gofmt -w "$i" 2>&1 1>/dev/null)
  status=$?
  if test $status -ne 0
  then
    echo -e "${REDBOLD}...Error: 'gofmt' command failed!${RESET}"
    echo -e "${REDBOLD}ERROR: ${ERROR}${RESET}"
    echo ""
    exit 1
  fi
  echo "  Formatted: $i"
done;
echo -e "${GREENBOLD}...Complete.${RESET}"
echo ""

echo -e "${CYANBOLD}Building Docker container:${RESET} ${DOCKERIMAGE}${RESET}"
ERROR=$(docker build -t $DOCKERIMAGE . 2>&1 1>/dev/null)
status=$?
if test $status -ne 0
then
	echo -e "${REDBOLD}...Error: 'docker build' command failed!${RESET}"
  echo -e "${REDBOLD}ERROR: ${ERROR}${RESET}"
  echo ""
  exit 1
fi
echo -e "${GREENBOLD}...Complete.${RESET}"
echo ""