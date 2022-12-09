#!/bin/bash
#
# From root of project, run: `bash docker_run.sh`
# You must have a DOCKERIMAGE envionment variable set, e.g.: add `export DOCKERIMAGE={your-docker-hub-account-username}/{your-docker-hub-image-name}` to your ~/.bashrc file.

clear

# Color Console Output
RESET='\033[0m'           # Text Reset
REDBOLD='\033[1;31m'      # Red (Bold)
GREENBOLD='\033[1;32m'    # Green (Bold)
YELLOWBOLD='\033[1;33m'   # Yellow (Bold)
CYANBOLD='\033[1;36m'     # Cyan (Bold)

if [ "$DOCKERIMAGE" = "" ]; then
  echo ""
  echo -e "${REDBOLD}Please set your DOCKERIMAGE environment variable:${RESET} ${CYANBOLD}{your-docker-hub-account-username}/{your-docker-hub-image-name}${RESET}"
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

echo -e "${CYANBOLD}Running Docker container:${RESET} ${DOCKERIMAGE}${RESET}"
docker run -it -p 5000:5000 --rm $DOCKERIMAGE