.PHONY: build
TARGET_DIR ?= build/app
build:
	go generate ./...
	go build -o ${TARGET_DIR}

start-ui-dev:
	cd ui/ && yarn dev

start-dev:
	docker compose -f ./deployments/docker-compose.dev.yml up -d

start-dev-clean:
	docker compose -f ./deployments/docker-compose.dev.yml up --force-recreate --build api

start:
	docker compose up -d
