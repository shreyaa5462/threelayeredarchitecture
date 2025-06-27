package taskstore

import (
	"ThreeLayeredArchitecture/models"
	"database/sql"
)

type TaskStore struct {
	DB *sql.DB
}

func NewTaskStore(db *sql.DB) *TaskStore {
	return &TaskStore{DB: db}
}

func (ts *TaskStore) Add(description string) (models.Task, error) {
	query := "INSERT INTO tasks (description, completed) VALUES (?, ?)"
	result, err := ts.DB.Exec(query, description, false)
	if err != nil {
		return models.Task{}, err
	}
	id, _ := result.LastInsertId()
	return models.Task{ID: int(id), Description: description, Completed: false}, nil
}

func (ts *TaskStore) GetPending() ([]models.Task, error) {
	query := "SELECT id, description, completed FROM tasks WHERE completed = FALSE ORDER BY id"
	rows, err := ts.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Description, &task.Completed)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (ts *TaskStore) GetByID(id int) (models.Task, error) {
	query := "SELECT id, description, completed FROM tasks WHERE id = ?"
	row := ts.DB.QueryRow(query, id)
	var task models.Task
	err := row.Scan(&task.ID, &task.Description, &task.Completed)
	return task, err
}

func (ts *TaskStore) MarkComplete(id int) error {
	query := "UPDATE tasks SET completed = TRUE WHERE id = ?"
	_, err := ts.DB.Exec(query, id)
	return err
}

func (ts *TaskStore) Delete(id int) error {
	query := "DELETE FROM tasks WHERE id = ?"
	_, err := ts.DB.Exec(query, id)
	return err
}
