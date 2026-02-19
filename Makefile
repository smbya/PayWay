start:
	docker-compose up -d

sqlc: 
	docker-compose exec dev sqlc generate

web-rebuild:
	docker-compose up --build -d web
	docker-compose logs -f web