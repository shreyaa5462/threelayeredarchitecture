// package task
//
// import (
//
//	"ThreeLayeredArchitecture/models"
//	"database/sql"
//
// )
//
//	type TaskStore struct {
//		DB *sql.DB
//	}
//
//	func NewTaskStore(db *sql.DB) *TaskStore {
//		return &TaskStore{DB: db}
//	}
//
//	func (s *TaskStore) CreateTask(description string) (models.MYTask, error) {
//		query := "INSERT INTO mytask (description, completed) VALUES (?, ?)"
//		result, err := s.DB.Exec(query, description, false)
//		if err != nil {
//			return models.MYTask{}, err
//		}
//		id, _ := result.LastInsertId()
//		return models.MYTask{ID: int(id), Description: description, Completed: false}, nil
//	}
//
//	func (s *TaskStore) GetPendingTasks() ([]models.MYTask, error) {
//		query := "SELECT id, description, completed FROM mytask WHERE completed = FALSE ORDER BY id"
//		rows, err := s.DB.Query(query)
//		if err != nil {
//			return nil, err
//		}
//		defer rows.Close()
//
//		var tasks []models.MYTask
//		for rows.Next() {
//			var t models.MYTask
//			rows.Scan(&t.ID, &t.Description, &t.Completed)
//			tasks = append(tasks, t)
//		}
//		return tasks, nil
//	}
//
//	func (s *TaskStore) GetTaskByID(id int) (models.MYTask, error) {
//		query := "SELECT id, description, completed FROM mytask WHERE id = ?"
//		row := s.DB.QueryRow(query, id)
//		var t models.MYTask
//		err := row.Scan(&t.ID, &t.Description, &t.Completed)
//		return t, err
//	}
//
//	func (s *TaskStore) MarkTaskCompleted(id int) error {
//		query := "UPDATE mytask SET completed = TRUE WHERE id = ?"
//		_, err := s.DB.Exec(query, id)
//		return err
//	}
//
//	func (s *TaskStore) DeleteTask(id int) error {
//		query := "DELETE FROM mytask WHERE id = ?"
//		_, err := s.DB.Exec(query, id)
//		return err
//	}
package task

import (
	"ThreeLayeredArchitecture/models"
	"database/sql"
	"log"
)

type TaskStore struct {
	DB *sql.DB
}

func NewTaskStore(db *sql.DB) *TaskStore {
	return &TaskStore{DB: db}
}

func (s *TaskStore) CreateTask(task models.MYTask) (models.MYTask, error) {
	query := "INSERT INTO mytask (description, completed) VALUES (?, ?)"
	result, err := s.DB.Exec(query, task.Description, false)

	if err != nil {
		return models.MYTask{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return models.MYTask{}, err
	}

	return models.MYTask{ID: int(id), Description: task.Description, Completed: false}, nil
}

func (s *TaskStore) GetPendingTasks() ([]models.MYTask, error) {
	query := "SELECT id, description, completed FROM mytask WHERE completed = FALSE ORDER BY id"
	rows, err := s.DB.Query(query)

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

func (s *TaskStore) GetTaskByID(id int) (models.MYTask, error) {
	query := "SELECT id, description, completed FROM mytask WHERE id = ?"
	row := s.DB.QueryRow(query, id)

	var t models.MYTask
	err := row.Scan(&t.ID, &t.Description, &t.Completed)

	return t, err
}

func (s *TaskStore) MarkTaskCompleted(id int) error {
	query := "UPDATE mytask SET completed = TRUE WHERE id = ?"
	_, err := s.DB.Exec(query, id)

	return err
}

func (s *TaskStore) DeleteTask(id int) error {
	query := "DELETE FROM mytask WHERE id = ?"
	_, err := s.DB.Exec(query, id)

	return err
}
