APP_CONTAINER = app

IMPORT_ENTRYPOINT = cmd/import-csv/main.go

dev/run/import:
	docker compose run --rm $(APP_CONTAINER) go run $(IMPORT_ENTRYPOINT)

dev/run/server:
	docker compose up --build
