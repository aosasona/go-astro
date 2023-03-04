.PHONY: build
TARGET_DIR ?= build/app
build:
	go generate ./...
	go build -o ${TARGET_DIR}

run-api-dev:
	air -c .air.toml

run-ui-dev:
	cd ui/ && yarn dev
