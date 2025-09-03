build:
	@go build -o bin/ecom src/main.go

test:
	@go test -v ./...

run: build
	@./bin/ecom

migration:
	@migrate create -ext sql -dir src/db/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run src/db/migrate/main.go up

migrate-down:
	@go run src/db/migrate/main.go down

