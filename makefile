postgresClient:
	sudo docker run --name=root -p 8001:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d postgres:11-alpine

postgrestransl:
	sudo docker run --name=root12 -p 8002:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d postgres:12-alpine

migrate:
	migrate create -ext sql -dir db/migration -seq init_schema  

sqlc:
	sqlc generate	

createdb:                
	sudo docker exec -it root createdb --username=root --owner=root client

dropdb:
	sudo docker exec -it root dropdb client	

migrateup:
	migrate -path db/migration -database "postgresql://root:postgres@localhost:5432/client?sslmode=disable" -verbose up 	


migratedown:
	migrate -path db/migration -database "postgresql://root:postgres@localhost:5432/client?sslmode=disable" -verbose down

test:
	go test -v ./...

certificate:
	go run $GOROOT/usr/local/go/src/crypto/tls/generate_cert.go --host=localhost    


main:
	go run main.go

dockerRun:
	sudo docker run --name client --network client-network -e GIN_MODE=release -p 8000:8000 -e DB_SOURCE_CLIENT="postgresql://root:postgres@root:5432/client?sslmode=disable" client:latest	

pfx:
	openssl pkcs12 -export -out domain.name.pfx -inkey domain.name.key -in domain.name.crt

key:
	openssl rand -hex 64	

keys:
	openssl rand -hex 64 | head -c 32

.PHONY: postgres migrate sqlc createdb dropdb migrateup migratedown test main postgrestransl key keys