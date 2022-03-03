BIN_DIR ?= bin

up:
	./scripts/docker_up.sh $(ENV)
.PHONY: up

shell:
	docker exec -it food_app sh
.PHONY: shell

api:
	go build -o $(BIN_DIR)/ ./cmd/api
.PHONY: api

command:
	go build -o $(BIN_DIR)/ ./cmd/command
.PHONY: command

build_all: api command
.PHONY: build_all
