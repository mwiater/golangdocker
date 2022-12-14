#!/bin/bash
#
# From root of project, run: `bash go_build.sh`

clear

# Color Console Output
RESET='\033[0m'           # Text Reset
REDBOLD='\033[1;31m'      # Red (Bold)
GREENBOLD='\033[1;32m'    # Green (Bold)
YELLOWBOLD='\033[1;33m'   # Yellow (Bold)
CYANBOLD='\033[1;36m'     # Cyan (Bold)

echo -e "${CYANBOLD}Building Swagger docs...${RESET}"
swag init
echo -e "${GREENBOLD}...Complete.${RESET}"
echo ""

echo -e "${CYANBOLD}Formatting *.go files...${RESET}"
for i in *.go **/*.go ; do
  gofmt -w "$i"
  echo "Formatted: $i"
done;
echo -e "${GREENBOLD}...Complete${RESET}"
echo ""

echo -e "${CYANBOLD}Building Go app:${RESET} ${GREENBOLD}go build -o bin/golangdocker .${RESET}"
go build -o bin/golangdocker .
echo ""

echo -e "${GREENBOLD}Complete: Built native go binary.${RESET}"
echo ""