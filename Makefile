
POSTGRES_DSN = "postgres://root:secret@postgres:5432"
DB_NAME = "agency_db"

install-tools:
	@echo installing tools && \
	@go install go.uber.org/mock/mockgen@latest \
	@go install github.com/swaggo/swag/cmd/swag@latest \
	@echo done

generate:
	@echo running code generation
	@go generate ./...
	@echo done

.PHONY: up
up: ## run the application on docker
	@docker compose up --build -d

.PHONY: down
down: ## stop the application on docker
	@docker compose down

.PHONY: db
db: ## create the database
	@docker exec -it postgres createdb --username=root --owner=root ${DB_NAME}

.PHONY: db-test
db-test: ## create the test database
	@docker exec -it postgres createdb --username=root --owner=root ${DB_NAME}_test

.PHONY: drop-db ## drop the database
	@docker exec -it postgres dropdb ${DB_NAME}

.PHONY: migration
migration: ## create new migration file
	@migrate create -ext sql -dir db/migrations -seq $(name)

.PHONY: migrate
migrate: ## apply all up migrations
	@echo "start migration in production database"
	@migrate -source file://db/migrations -database $(POSTGRES_DSN)/${DB_NAME}?sslmode=disable up $(version)
	@echo "done"
	@echo "start migration in test database"
	@migrate -source file://db/migrations -database $(POSTGRES_DSN)/${DB_NAME}_test?sslmode=disable up $(version)
	@echo "done"

.PHONY: migrate-down
migrate-down: ## apply all down migrations
	@migrate -source file://db/migrations -database $(POSTGRES_DSN)/${DB_NAME}?sslmode=disable down $(version)
	@migrate -source file://db/migrations -database $(POSTGRES_DSN)/${DB_NAME}_test?sslmode=disable down $(version)

.PHONY: test
test: ## unit testing with coverage
	@go test -v -cover ./...

.PHONY: doc
doc: ## generate docs
	swag init -g ./application/swagger.go

