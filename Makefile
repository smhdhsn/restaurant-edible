up:
	./scripts/docker_up.sh $(ENV)
bash:
	docker exec -it food_app bash
api:
	go build -o $(BIN_DIR)/ ./cmd/api
command:
	go build -o $(BIN_DIR)/ ./cmd/command
build_all: api command
.PHONY: up bash api command build_all
