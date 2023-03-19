include .env

.PHONY: migrate-create
migrate-create:
	@docker run --rm -v "$(shell pwd)/migrations:/migrations" \
		migrate/migrate:v4.15.2 create \
		-ext sql -dir /migrations $(name)

.PHONY: migrate-up
migrate-up:
	@docker run --rm -v "$(shell pwd)/migrations:/migrations" \
		migrate/migrate:v4.15.2 up \
		-database $(DB_DRIVER)://$(DB_USER):$(MYSQL_ROOT_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME) \
		-path /migrations

.PHONY: migrate-down
migrate-down:
	@docker run --rm -v "$(shell pwd)/migrations:/migrations" \
		migrate/migrate:v4.15.2 down \
		-database $(DB_DRIVER)://$(DB_USER):$(MYSQL_ROOT_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME) \
		-path /migrations

.PHONY: proto
proto:
	@protoc --plugin=grpc \
		--go_out=internal/common/genproto \
		--go-grpc_out=internal/common/genproto \
		--proto_path=api/protobuf auth.proto