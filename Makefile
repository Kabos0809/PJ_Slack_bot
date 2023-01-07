up:
	docker compose up -d --build
down:
	docker compose down
logs-go: 
	docker compose logs go
logs-db:
	docker compose logs db