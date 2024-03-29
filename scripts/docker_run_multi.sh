#!/bin/bash
#
# From root of project, run: `bash scripts/docker_run_multi.sh`

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

if [ "$SERVERPORT" = "" ]; then
  echo ""
  echo -e "${REDBOLD}Please set your SERVERPORT environment variable in the .env file:${RESET} E.g.: ${CYANBOLD}SERVERPORT=5000${RESET}"
  echo ""
  exit 0
fi

if [ "$DOCKERPORT" = "" ]; then
  echo ""
  echo -e "${REDBOLD}Please set your DOCKERPORT environment variable in the .env file:${RESET} E.g.: ${CYANBOLD}DOCKERPORT=5000${RESET}"
  echo ""
  exit 0
fi

echo -e "${CYANBOLD}Starting Docker container:${RESET} golangdocker01${RESET}"
ERROR=$(docker run -d -p $DOCKERPORT:$SERVERPORT --rm --name golangdocker01 --hostname golangdocker01 $DOCKERIMAGE 2>&1 1>/dev/null)

EXITCODE=$?
if test $EXITCODE -ne 0; then
  echo ""
  echo -e "${REDBOLD}ERROR: ${ERROR}${RESET}"
  echo ""
  exit 1
else
echo -e "${GREENBOLD}...Complete.${RESET}"
echo ""
fi

DOCKERPORT2=$((DOCKERPORT + 1))

echo -e "${CYANBOLD}Starting Docker container:${RESET} golangdocker02${RESET}"
ERROR=$(docker run -d -p $DOCKERPORT2:$SERVERPORT --rm --name golangdocker02 --hostname golangdocker02 $DOCKERIMAGE 2>&1 1>/dev/null)

EXITCODE=$?
if test $EXITCODE -ne 0; then
  echo ""
  echo -e "${REDBOLD}ERROR: ${ERROR}${RESET}"
  echo ""
  exit 1
else
echo -e "${GREENBOLD}...Complete.${RESET}"
echo ""
fi

docker ps
echo ""