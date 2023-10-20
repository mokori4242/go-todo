login-db: #DBコンテナにログイン
	docker compose exec db bash
login-posgre: #posgreにログイン
	docker compose exec db psql -U posgre
login-app: #appコンテナにログイン
	docker compose exec app bash
show-table: #posgreのテーブルを表示
	docker-compose exec db psql -U posgre -c "\dt"
up: #コンテナを立ち上げる
	docker compose up --build
down: #コンテナを終了する
	docker compose down --remove-orphans
go:
	docker compose exec app go run main.go
test:
	docker compose exec app go test -v
parallel-test:
	docker compose exec app go test -v -short -parallel 3
bench-test:
	docker compose exec app go test -v -cover -short -bench .
mod-init: #依存関係を管理するgo.modファイルの新規作成
	docker compose exec app go mod init
mod-tidy: #依存関係を管理するgo.modファイルの更新
	docker compose exec app go mod tidy