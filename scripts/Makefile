.PHONY: build
TARGET_DIR ?= build/app
build:
	go generate ./...
	go build -o ${TARGET_DIR}

start-dev:
	docker compose -f ./deployments/docker-compose.dev.yml up -d

start:
	docker compose up -d
