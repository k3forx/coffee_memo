DB_USER := root
DB_PASS := password
DB_PORT := 3306
DB_NAME := coffee_app

.PHONY: migrate-create
migrate-create: ## Create new migration file (usage: make migrate-create NAME=<string>)
	goose -dir ./migration create $(NAME) sql && goose -dir ./migration fix

.PHONY: migrate-status
migrate-status: ## Show current status of migration
	GOOSE_DRIVER=mysql GOOSE_DBSTRING="$(DB_USER):$(DB_PASS)@/$(DB_NAME)" goose -dir ./migration status

.PHONY: migrate-up
migrate-up: ## Migrate the DB to the most recent version available
	GOOSE_DRIVER=mysql GOOSE_DBSTRING="$(DB_USER):$(DB_PASS)@/$(DB_NAME)" goose -dir ./migration up

.PHONY: migrate-down
migrate-down: ## Roll back the version by 1
	GOOSE_DRIVER=mysql GOOSE_DBSTRING="$(DB_USER):$(DB_PASS)@/$(DB_NAME)" goose -dir ./migration down

.PHONY: mysql
mysql: ## Conntect to mysql container
	docker-compose exec mysql mysql -u$(DB_USER) -p$(DB_PASS) $(DB_NAME)

.PHONY: generate-ent-schema
generate-ent-schema: ## Generate ent schema
	cd ./go && go run ariga.io/entimport/cmd/entimport -dsn "mysql://$(DB_USER):$(DB_PASS)@tcp(localhost:$(DB_PORT))/$(DB_NAME)" -schema-path ./pkg/ent/schema && go run -mod=mod entgo.io/ent/cmd/ent generate ./pkg/ent/schema
