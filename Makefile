postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -d postgres:12-alpine

createdb:
	docker exec -it postgres createdb --username=admin --owner=admin simple_bank
dropdb:
	docker exec -it postgres dropdb -U admin simple_bank
migrateup:
	migrate -path db/migration -database "postgres://admin:admin@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgres://admin:admin@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v ./...
server:
	go run main.go
	
.PHONY: postgres, createdb, dropdb, migrateup, migratedown, sqlc, test, server