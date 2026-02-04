include .env
export

run:
	go run main.go

build:
	go build main.go

test:
	go test ./...

sql:
	sqlc generate

DB_DSN=$(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)?$(DB_PARAMS)
goose-up:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_DSN)" up

goose-down:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_DSN)" down

goose-status:
	goose -dir $(MIGRATIONS_DIR) $(DB_DRIVER) "$(DB_DSN)" status
