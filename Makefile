build:
	@go build -o bin/ecom cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/ecom

migration-create:
	@migrate create -ext sql -dir cmd/migrate/migrations $(name)

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down

migrate-status:
	@go run cmd/migrate/main.go version
