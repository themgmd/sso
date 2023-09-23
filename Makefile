.PHONY: dev run build new_migration install_tooling up

install_tooling:
	go install github.com/rubenv/sql-migrate/...@latest

new_migration:
	@read -p "Enter Migration Name:" migration_name; \
	sql-migrate new --config=dbconfig.yml -env=universal_table $$migration_name

build: 
	go build -o .bin/main cmd/main/app.go

up:
	docker-compose up -d --build

dev:
	go run cmd/main/app.go

run: build
	.bin/main