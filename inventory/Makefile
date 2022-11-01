migrate-up:
	migrate -database "postgres://postgres:postgres@localhost:5432/inventory?sslmode=disable" -path sql/migration up

migrate-down:
	migrate -database "postgres://postgres:postgres@localhost:5432/inventory?sslmode=disable" -path sql/migration down

sqlc:
	sqlc generate 

test:
	go test -v -cover ./...

server:
	go run cmd/main.go