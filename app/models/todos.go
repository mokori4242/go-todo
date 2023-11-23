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

func GetTodo(ctx context.Context, id int) (todo Todo, err error) {
	cmd := `SELECT * FROM todos WHERE id = $1`
	row := config.Db.QueryRowContext(ctx, cmd, id)
	if err = row.Scan(
		&todo.ID,
		&todo.UserID,
		&todo.Content,
		&todo.CreatedAt,
	); err != nil {
		log.Fatalln(err)
	}
	return todo, err
}

func GetTodos(ctx context.Context) (todos []Todo, err error) {
	cmd := `SELECT * FROM todos`
	rows, err := config.Db.QueryContext(ctx, cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(
			&todo.ID,
			&todo.UserID,
			&todo.Content,
			&todo.CreatedAt,
		); err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}

func (u *User) GetTodosByUser(ctx context.Context) (todos []Todo, err error) {
	cmd := `SELECT * FROM todos WHERE user_id = $1`
	rows, err := config.Db.QueryContext(ctx, cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(
			&todo.ID,
			&todo.UserID,
			&todo.Content,
			&todo.CreatedAt,
		); err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}

func (t *Todo) UpdateTodoContent(ctx context.Context, content string) (err error) {
	cmd := `UPDATE todos SET content = $1 WHERE id = $2`
	_, err = config.Db.ExecContext(ctx, cmd, content, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (t *Todo)DeleteTodo(ctx context.Context) (err error) {
	cmd := `DELETE FROM todos WHERE id = $1`
	_, err = config.Db.ExecContext(ctx, cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully deleted todo.")
	return err
}