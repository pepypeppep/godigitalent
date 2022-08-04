#!make
include .env
export

sqlc:
	sqlc generate

migration:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrate:
	go run db/migrate.go