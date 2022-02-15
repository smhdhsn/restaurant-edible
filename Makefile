up:
	docker run -p 3306:3306 -d --name food_db -e MYSQL_USER=food -e MYSQL_PASSWORD=food -e MYSQL_DATABASE=food -e MYSQL_RANDOM_ROOT_PASSWORD=true mysql
down:
	docker rm -f food_db
log:
	docker logs -f food_db
recipe:
	go run cmd/command/*.go recipe -j ./sample/recipes.json
buy:
	go run cmd/command/*.go buy
test:
	go test ./...
.PHONY: up down log recipe buy test
