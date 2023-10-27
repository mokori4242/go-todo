package models

import (
	"log"
	"database/sql"
	"app/config"
	"fmt"
	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Config.DBHost, config.Config.DBPort, config.Config.DBUser, config.Config.DBPassword, config.Config.DBName)
	Db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id SERIAL PRIMARY KEY,
		uuid UUID NOT NULL UNIQUE,
		name TEXT,
		email TEXT,
		password TEXT,
		created_at TIMESTAMP)`, tableNameUser)

	Db.Exec(cmdU)
}