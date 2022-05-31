.PHONY: up purge shell build_server build_all proto_inventory proto_recipe proto_menu proto_all

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
build_server:
	@go build -o $(BIN_DIR)/ ./cmd/server

# builds all the entry points of the application.
build_all: build_server

# compiles proto files related to edible inventory.
proto_inventory:
	@protoc --go_out=internal/protos/edible/inventory/ --go-grpc_out=require_unimplemented_servers=false:internal/protos/edible/inventory/ protos/edible/inventory/*.proto

# compiles proto files related to edible recipe.
proto_recipe:
	@protoc --go_out=internal/protos/edible/recipe/ --go-grpc_out=require_unimplemented_servers=false:internal/protos/edible/recipe/ protos/edible/recipe/*.proto

# compiles proto files related to edible menu.
proto_menu:
	@protoc --go_out=internal/protos/edible/menu/ --go-grpc_out=require_unimplemented_servers=false:internal/protos/edible/menu/ protos/edible/menu/*.proto

# compiles all proto files.
proto_all: proto_menu proto_recipe proto_inventory
