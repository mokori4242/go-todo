package models

import (
	"log"
	"context"
	"app/config"
	"fmt"
)

const tableNameT = "todos"

func CreateTodosTable(ctx context.Context) {
	cmd := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id SERIAL PRIMARY KEY,
		user_id INT,
		content TEXT,
		created_at TIMESTAMP)`, tableNameT)

	if _, err := config.Db.ExecContext(ctx, cmd); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Successfully created tables.")
}