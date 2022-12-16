SHELL=/bin/bash

.DEFAULT_GOAL := list

# Color Console Output
RESET=\033[0m
REDBOLD=\033[1;31m
GREENBOLD=\033[1;32m
YELLOWBOLD=\033[1;33m
CYANBOLD=\033[1;36m

.PHONY: list
list:
	@echo ""
	@echo -e "${GREENBOLD}Targets in this Makefile:${RESET}"
	@echo ""
	@LC_ALL=C $(MAKE) -pRrq -f $(lastword $(MAKEFILE_LIST)) : 2>/dev/null | awk -v RS= -F: '/(^|\n)# Files(\n|$$)/,/(^|\n)# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'
	@echo ""

golang-run:
	scripts/go_run.sh

golang-build:
	scripts/go_build.sh

golang-build-arm64:
	scripts/go_build_arm64.sh

docker-build:
	scripts/docker_build.sh

docker-run:
	scripts/docker_run.sh