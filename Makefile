.PHONY: up down force-down

up:
	docker compose up -d

down:
	docker compose down

force-down:
	docker compose down -v