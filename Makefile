up:
	./script/docker_up.sh $(APP_MODE)
bash:
	docker exec -it restaurant_menu_app bash
buy:
	docker exec restaurant_menu_app ./command buy
recipe:
	docker exec restaurant_menu_app ./command recipe -j sample/recipes.json
recycle:
	docker exec restaurant_menu_app ./command recycle -fe
server:
	go build -o $(BIN_DIR)/ ./cmd/server
command:
	go build -o $(BIN_DIR)/ ./cmd/command
build_all: server command
.PHONY: up bash buy recipe recycle server command build_all 
