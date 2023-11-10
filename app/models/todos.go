package models

import (
	"time"
	"log"
	"context"
	"app/config"
	"fmt"
)

const tableNameT = "todos"

type Todo struct {
	ID int
	UserID int
	Content string
	CreatedAt time.Time
}

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

func (u *User) CreateTodo(ctx context.Context, content string) (err error) {
	cmd := `INSERT INTO todos (
		user_id,
		content,
		created_at) VALUES ($1, $2, $3)`

	_, err = config.Db.ExecContext(ctx, cmd, u.ID, content, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}