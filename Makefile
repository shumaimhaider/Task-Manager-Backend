MIGRATE_CMD := go run migrations/*.go


# DB migrations

migrate-init:
	$(MIGRATE_CMD) init

migrate-up:
	$(MIGRATE_CMD) up

migrate-down:
	$(MIGRATE_CMD) down


# Run Locally

dev-local:
	go run routes/main.go
