include .env

up:
	docker-compose up -d

go:
	docker-compose exec -it md2s /bin/sh

db:
	docker exec -it ${DB_HOST} psql -U ${DB_USER} -d ${DB_NAME}