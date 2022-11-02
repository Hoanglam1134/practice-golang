migrate-up:
	migrate -database "postgres://postgres:postgres@localhost:5432/orders?sslmode=disable" -path sql/migration up

migrate-down:
	migrate -database "postgres://postgres:postgres@localhost:5432/orders?sslmode=disable" -path sql/migration down

sqlc:
	sqlc generate 

test:
	go test -v -cover ./...

server:
	go run cmd/main.go

protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    api/*.proto