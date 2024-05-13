.PHONY: test up down execute

test:
	go test ./...

up:
	docker-compose up --build -d

down:
	docker-compose down

execute:
	docker-compose exec app bash
