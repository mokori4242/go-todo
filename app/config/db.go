package config

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitDB() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		Config.DBHost, Config.DBPort, Config.DBUser, Config.DBPassword, Config.DBName)
	var err error
	Db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	if err = Db.Ping(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Successfully connected to the database.")
}