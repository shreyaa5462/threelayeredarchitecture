package task

import (
	"ThreeLayeredArchitecture/models"

	"gofr.dev/pkg/gofr"
	"log"
)

type TaskStore struct {
}

func NewTaskStore() *TaskStore {
	return &TaskStore{}
}

func (s *TaskStore) CreateTask(ctx *gofr.Context, task models.MYTask) (models.MYTask, error) {
	query := "INSERT INTO mytask (description, completed) VALUES (?, ?)"
	result, err := ctx.SQL.Exec(query, task.Description, false)

	if err != nil {
		return models.MYTask{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return models.MYTask{}, err
	}

	return models.MYTask{ID: int(id), Description: task.Description, Completed: false}, nil
}

func (s *TaskStore) GetPendingTasks(ctx *gofr.Context) ([]models.MYTask, error) {
	query := "SELECT id, description, completed FROM mytask WHERE completed = FALSE ORDER BY id"
	rows, err := ctx.SQL.Query(query)

	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rows.Close(); err != nil {
			log.Println("Error closing rows:", err)
		}
	}()

	var tasks []models.MYTask

	for rows.Next() {
		var t models.MYTask
		if err := rows.Scan(&t.ID, &t.Description, &t.Completed); err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (s *TaskStore) GetTaskByID(ctx *gofr.Context, id int) (models.MYTask, error) {
	query := "SELECT id, description, completed FROM mytask WHERE id = ?"
	row := ctx.SQL.QueryRow(query, id)

	var t models.MYTask
	err := row.Scan(&t.ID, &t.Description, &t.Completed)

	return t, err
}

func (s *TaskStore) MarkTaskCompleted(ctx *gofr.Context, id int) error {
	query := "UPDATE mytask SET completed = TRUE WHERE id = ?"
	_, err := ctx.SQL.Exec(query, id)

	return err
}

func (s *TaskStore) DeleteTask(ctx *gofr.Context, id int) error {
	query := "DELETE FROM mytask WHERE id = ?"
	_, err := ctx.SQL.Exec(query, id)

	return err
}
