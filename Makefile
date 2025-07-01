include .env
export

migrate-create:  ### create new migration
	migrate create -ext sql -dir migrations '$(word 2,$(MAKECMDGOALS))'
.PHONY: migrate-create

migrate-up: ### migration up
	migrate -path migrations -database '$(POSTGRES_URL)?sslmode=disable' up
.PHONY: migrate-up
 
