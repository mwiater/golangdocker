#!/bin/bash
#
# From root of project, run: `bash scripts/golang_test.sh`

clear

# Color Console Output
RESET='\033[0m'           # Text Reset
REDBOLD='\033[1;31m'      # Red (Bold)
GREENBOLD='\033[1;32m'    # Green (Bold)
YELLOWBOLD='\033[1;33m'   # Yellow (Bold)
CYANBOLD='\033[1;36m'     # Cyan (Bold)

echo -e "${CYANBOLD}Running: go generate${RESET}"
go generate ./...
echo -e "${GREENBOLD}...Complete.${RESET}"

echo -e "${CYANBOLD}Running tests...${RESET}"
go clean -testcache
go test -v $(go list ./... | grep -v /docs)
echo -e "${GREENBOLD}...Complete.${RESET}"
echo ""