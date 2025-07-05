APP_CONTAINER := app
IMPORT_ENTRYPOINT := cmd/import-csv/main.go

.PHONY: dev/run/import dev/run/server

# ──────────
# CSV → DB 取り込み
# (MySQL起動 → マイグレーション → CSVインポート)
# ──────────
dev/run/import:
	@echo "==> Booting MySQL and running migrations..."
	docker compose up -d mysql migrate

	@echo "==> Importing CSV to DB..."
	docker compose run --rm $(APP_CONTAINER) \
		go run $(IMPORT_ENTRYPOINT)

	@echo "==> Done!"

# ──────────
# HTTP サーバー起動
# (MySQL起動 → マイグレーション → app起動)
# ──────────
dev/run/server:
	@echo "==> Starting full stack..."
	docker compose up --build
