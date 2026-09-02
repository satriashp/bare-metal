package repository

import (
	"todo-core/internal/db"
	"todo-core/internal/models"
)

func GetTodos() ([]models.Todo, error) {
	rows, err := db.DB.Query(`
		SELECT id, text, done
		FROM todos
		ORDER BY id DESC
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var t models.Todo
		err := rows.Scan(&t.ID, &t.Text, &t.Done)
		if err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}

	return todos, nil
}

func CreateTodo(text string) error {
	_, err := db.DB.Exec(`
		INSERT INTO todos (text, done)
		VALUES ($1, false)
	`, text)

	return err
}
