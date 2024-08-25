.PHONY: test up down execute

test:
	go test ./...

up:
	docker compose up --build -d

down:
	docker compose down

execute_app:
	docker compose exec app bash

execute_rabbit:
	docker compose exec rabbit bash
