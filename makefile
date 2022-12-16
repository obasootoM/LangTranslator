postgres:
	sudo docker run --name=root -p 8001:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d postgres:11-alpine

migrate:
	migrate create -ext sql -dir db/migration -seq init_schema  

sqlc:
	sqlc generate	

createdb:
	sudo docker exec -it root createdb --username=root --owner=root client

dropdb:
	sudo docker exec -it root dropdb client	

migrateup:
	migrate -path db/migration -database "postgresql://root:postgres@localhost:8001/client?sslmode=disable" -verbose up 	


migratedown:
	migrate -path db/migration -database "postgresql://root:postgres@localhost:8001/client?sslmode=disable" -verbose down


.PHONY: postgres migrate sqlc createdb dropdb migrateup migratedown