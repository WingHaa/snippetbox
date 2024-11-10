dev:
	go run ./cmd/web

build:
	docker compose build
	
up:
	docker compose up

db-connection:
	docker compose exec db mysql -u docker -pdocker snippetbox

down:
	docker compose down
