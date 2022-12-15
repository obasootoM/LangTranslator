postgres:
	sudo docker run --name postgres11 -p 8001:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d postgres:11-alpine

migrate:
	migrate create -ext sql -dir db/migration -seq init_schema

sqlc:
	sqlc generate	

.PHONY: postgres migrate sqlc