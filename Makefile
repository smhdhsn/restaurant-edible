up:
	docker run -p 3306:3306 -d --name food_db -e MYSQL_USER=food -e MYSQL_PASSWORD=food -e MYSQL_DATABASE=food -e MYSQL_RANDOM_ROOT_PASSWORD=true mysql
down:
	docker rm -f food_db
log:
	docker logs -f food_db
.PHONY: up down
