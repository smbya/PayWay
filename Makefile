include ./app/.env

start:
	docker-compose up -d

sqlc:
	docker-compose run --rm dev sqlc generate

migrate-up:
	docker-compose run --rm dev migrate -path ./${DB_MIGRATIONS_DIRECTORY} -database "$(DB_URL)" up

migrate-down:
	docker-compose run --rm dev migrate -path ./${DB_MIGRATIONS_DIRECTORY} -database "$(DB_URL)" down

migrate-reload:
	docker-compose run --rm dev migrate -path ./${DB_MIGRATIONS_DIRECTORY} -database "$(DB_URL)" drop
	docker-compose run --rm dev migrate -path ./${DB_MIGRATIONS_DIRECTORY} -database "$(DB_URL)" up

migrate-create:
	migrate create -ext sql -dir ./app/${DB_MIGRATIONS_DIRECTORY} -seq $(name)

tidy:
	cd app && go mod tidy
	docker-compose run --rm dev go mod tidy

web-rebuild:
	docker-compose up --build -d web
	docker-compose logs -f web