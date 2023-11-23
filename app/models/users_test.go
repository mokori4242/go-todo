package models

import (
	"context"
	"log"
	"testing"
	"app/config"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateUsersTable(t *testing.T) {
	// モックデータベースとsqlmockオブジェクトを作成
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// configパッケージのDbをモックのデータベース接続に置き換える
	config.Db = db

	mock.ExpectExec(`^CREATE TABLE IF NOT EXISTS users\(` +
		` id SERIAL PRIMARY KEY,` +
		` uuid UUID NOT NULL UNIQUE,` +
		` name TEXT,` +
		` email TEXT,` +
		` password TEXT,` +
		` created_at TIMESTAMP\)$`).
		WillReturnResult(sqlmock.NewResult(0, 0))

	// テーブル作成関数を実行
	CreateUsersTable(context.Background())

	// モックが期待した通りの操作を受けたか確認
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}