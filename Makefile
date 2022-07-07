APP_MODE ?= local

# runs the script which loads the containers of the application.
up:
	@./script/docker_up.sh $(APP_MODE)

# deletes application's containers.
purge:
	@docker rm -f restaurant_edible_app restaurant_edible_db
	@docker volume rm restaurant_edible

# accesses the shell of application's container.
shell:
	@docker exec -it restaurant_edible_app bash

# builds server's http entry point.
build-server:
	@go build -o $(BIN_DIR)/ ./cmd/server

# builds all the entry points of the application.
build-all: build-server

# compiles proto files related to edible inventory.
proto-inventory:
	@protoc --go_out=internal/protos/edible/inventory/ --go-grpc_out=require_unimplemented_servers=false:internal/protos/edible/inventory/ protos/edible/inventory/*.proto

# compiles proto files related to edible recipe.
proto-recipe:
	@protoc --go_out=internal/protos/edible/recipe/ --go-grpc_out=require_unimplemented_servers=false:internal/protos/edible/recipe/ protos/edible/recipe/*.proto

# compiles proto files related to edible menu.
proto-menu:
	@protoc --go_out=internal/protos/edible/menu/ --go-grpc_out=require_unimplemented_servers=false:internal/protos/edible/menu/ protos/edible/menu/*.proto

# compiles all proto files.
proto-all: proto-menu proto-recipe proto-inventory

.PHONY: up purge shell build-server build-all proto-inventory proto-recipe proto-menu proto-all