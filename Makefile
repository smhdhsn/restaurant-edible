up:
	./scripts/docker_up.sh $(ENV)
shell:
	docker exec -it food_app sh
api:
	go build -o $(BIN_DIR)/ ./cmd/api
command:
	go build -o $(BIN_DIR)/ ./cmd/command
build_all: api command
.PHONY: up shell api command build_all
