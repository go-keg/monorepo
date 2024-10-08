GIT_VERSION=$(shell git describe --tags --always | sed "s|refs/tags/||g")
GIT_BRANCH=$(shell git branch | sed -n '/\* /s///p' | sed "s/-master//g")

# Includes
include .env
include .env.k8s
include scripts/build.mk
include scripts/compose.mk
include scripts/init.mk
include scripts/k8s.mk

version:
	@echo $(GIT_VERSION) $(GIT_BRANCH)

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\.0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
