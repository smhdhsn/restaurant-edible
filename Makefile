up:
	./script/docker_up.sh $(APP_MODE)
bash:
	docker exec -it restaurant_edible_app bash
build_server:
	go build -o $(BIN_DIR)/ ./cmd/server
build_all: build_server
proto_menu:
	protoc protos/edible/menu/*.proto --go_out=plugins=grpc:internal/
proto_recipe:
	protoc protos/edible/recipe/*.proto --go_out=plugins=grpc:internal/
proto_inventory:
	protoc protos/edible/inventory/*.proto --go_out=plugins=grpc:internal/
proto_all: proto_menu proto_recipe proto_inventory
.PHONY: up bash build_server build_all proto_menu proto_recipe proto_inventory proto_all
