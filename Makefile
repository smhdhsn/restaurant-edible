up:
	./scripts/docker_up.sh $(ENV)
bash:
	docker exec -it food_app bash
buy:
	docker exec food_app ./command buy
recipe:
	docker exec food_app ./command recipe -j sample/recipes.json
recycle:
	docker exec food_app ./command recycle -fe
api:
	go build -o $(BIN_DIR)/ ./cmd/api
command:
	go build -o $(BIN_DIR)/ ./cmd/command
build_all: api command
.PHONY: up bash buy recipe recycle api command build_all 
