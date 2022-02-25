postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -d postgres:12-alpine
restartpq:
	docker restart postgres
createdb:
	docker exec -it postgres createdb --username=admin --owner=admin simple_bank
dropdb:
	docker exec -it postgres dropdb -U admin simple_bank
migrateup:
	migrate -path db/migration -database "postgres://admin:admin@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgres://admin:admin@localhost:5432/simple_bank?sslmode=disable" -verbose down
migrateup1:
	migrate -path db/migration -database "postgres://admin:admin@localhost:5432/simple_bank?sslmode=disable" -verbose up 1
migratedown1:
	migrate -path db/migration -database "postgres://admin:admin@localhost:5432/simple_bank?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go  github.com/vukieuhaihoa/simplebank/db/sqlc Store
.PHONY: postgres, restartpq, createdb, dropdb, migrateup, migratedown, sqlc, test, server, mock, migrateup1, migratedown1