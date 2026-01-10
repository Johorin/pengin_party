# マイグレーションファイルを作成
create-migration:
	@if [ -z "$(name)" ]; then \
		echo "Error: grep parameter is required. Usage: make migration name=<migration_name>"; \
		exit 1; \
	fi
	docker compose exec backend migrate create -ext sql -dir db/migrations $(name)

# マイグレーションを適用
migrate-up:
	docker compose exec backend migrate -path db/migrations -database "mysql://user:hogehoge@tcp(db)/pengin_party" up

# マイグレーションをdown
migrate-down:
	docker compose exec backend migrate -path db/migrations -database "mysql://user:hogehoge@tcp(db)/pengin_party" down

# マイグレーション状態をクリア（schema_migrationsテーブルのレコードを削除）
migrate-clear:
	docker compose exec mysql-db mysql -u user -phogehoge pengin_party -e "DELETE FROM schema_migrations;"