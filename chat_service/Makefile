CURRENT_DIR=$(shell pwd)

-include .env

DB_URL="postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"


run:
	go run cmd/main.go

gen-proto:
	./script/gen-proto.sh


migrate_file:
	migrate create -ext sql -dir migrations -seq user_chats_table

migrate_up:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrate_down:
	migrate -path migrations -database "$(DB_URL)" -verbose down 

migrate_forse:
	migrate -path migrations -database "$(DB_URL)" -verbose forse 5
