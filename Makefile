start:
	docker-compose up -d

sqlc: 
	docker-compose exec dev sqlc generate

tidy:
	cd app && go mod tidy

web-rebuild:
	docker-compose up --build -d web
	docker-compose logs -f web