GIT_VERSION=$(shell git describe --tags --always | sed "s|refs/tags/||g")
GIT_BRANCH=$(shell git branch | sed -n '/\* /s///p' | sed "s/-main//g")

# Includes
include .env
include scripts/base.mk
include scripts/compose.mk
include scripts/init.mk

%.lint:
	$(eval SERVICE:= $*)
	@clear
	@echo "lint: $(SERVICE)"
	golangci-lint run -c .golangci.yml --fix ./cmd/$(SERVICE)/... ./internal/app/$(SERVICE)/...

lint:
	@clear
	golangci-lint run -c .golangci.yml --fix ./...

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
