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

test:
	go test -v ./...

certificate:
	go run $GOROOT/usr/local/go/src/crypto/tls/generate_cert.go --host=localhost    


main:
	go run main.go

pfx:
	openssl pkcs12 -export -out domain.name.pfx -inkey domain.name.key -in domain.name.crt

.PHONY: postgres migrate sqlc createdb dropdb migrateup migratedown test main