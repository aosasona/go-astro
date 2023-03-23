.PHONY: build
TARGET_DIR ?= build/app
build:
	go generate ./...
	CGO_ENABLED=0 go build -o ${TARGET_DIR}

start-ui-dev:
	cd ui/ && yarn dev

start-dev:
	docker compose -f ./deployments/docker-compose.dev.yml up -d

start-dev-clean:
	docker compose -f ./deployments/docker-compose.dev.yml up --force-recreate --build api

start:
	docker compose up -d

start-clean:
	docker compose up --force-recreate --build app
