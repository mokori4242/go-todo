package models

import (
	"log"
	"context"
	"app/config"
	"fmt"
)

const tableNameS = "sessions"

func CreateSessionsTable(ctx context.Context) {
	cmd := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id SERIAL PRIMARY KEY,
		uuid UUID NOT NULL UNIQUE,
		email TEXT,
		user_id INT,
		created_at TIMESTAMP)`, tableNameS)

	if _, err := config.Db.ExecContext(ctx, cmd); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Successfully created sessions tables.")
}
