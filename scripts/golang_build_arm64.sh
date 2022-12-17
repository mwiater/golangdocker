#!/bin/bash
#
# From root of project, run: `bash golang_build_arm64.sh`
#
# Find the architec ture of the taget machine. On Linus, for example: `uname -a` #=>
# Linux piarmy-01 5.4.0-1074-raspi #85-Ubuntu SMP PREEMPT Fri Nov 4 13:34:24 UTC 2022 aarch64 aarch64 aarch64 GNU/Linux
# In the case above aarch64 is the same as arm64 which my current version of go supports (See: `go tool dist list`)

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

echo -e "${CYANBOLD}Building Go app:${RESET} ${GREENBOLD}env GOOS=linux GOARCH=arm64 go build -o bin/golangdocker-arm64 .${RESET}"
env GOOS=linux GOARCH=arm64 go build -o bin/golangdocker-arm64 .
echo ""

echo -e "${GREENBOLD}Complete: Built arm64/aarch64 go binary.${RESET}"
echo ""