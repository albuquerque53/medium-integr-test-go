up:
	docker-compose --file build/docker-compose.yml up -d || docker compose --file build/docker-compose.yml up -d
down:
	docker-compose --file build/docker-compose.yml down || docker compose --file build/docker-compose.yml down
api:
	docker exec -it users_api bash
migrate_up:
	migrate -path internal/infra/db/migrations -database "mysql://users_dev:users_password@tcp(users_db:3306)/users_db" -verbose up
migrate_down:
	migrate -path internal/infra/db/migrations -database "mysql://users_dev:users_password@tcp(users_db:3306)/users_db" -verbose down
run:
	go run ./cmd/main.go