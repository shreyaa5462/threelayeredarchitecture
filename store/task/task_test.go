package task_test

import (
	"ThreeLayeredArchitecture/models"
	"ThreeLayeredArchitecture/store/task"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

//func TestCreateTask(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("unexpected error: %s", err)
//	}
//
//	store := task.NewTaskStore(db)
//	description := "Clean room"
//
//	mock.ExpectExec("INSERT  INTO mytask").
//		WithArgs(description, false).
//		WillReturnResult(sqlmock.NewResult(1, 1))
//
//	result, err := store.CreateTask(description)
//	if err != nil {
//		t.Errorf("expected no error, got %v", err)
//	}
//	if result.ID != 1 || result.Description != description || result.Completed != false {
//		t.Errorf("unexpected task result: %+v", result)
//	}
//}

func TestCreateTask(t *testing.T) {
	tests := []struct {
		description string
		expected    models.MYTask
		expectErr   bool
	}{
		{
			description: "add",
			expected: models.MYTask{
				ID:          1,
				Description: "add",
				Completed:   false,
			},
			expectErr: false,
		},

		{
			description: "subtract",
			expected: models.MYTask{
				ID:          2,
				Description: "subtract",
				Completed:   false,
			},
			expectErr: false,
		},
	}

	for i, tt := range tests {
		i = i + 1
		db, mock, err := sqlmock.New()
		if err != nil {
			log.Fatalf("failed to create mock db: %v", err)
		}
		store := task.NewTaskStore(db)
		mock.ExpectExec("INSERT INTO mytask").
			WithArgs(tt.description, false).
			WillReturnResult(sqlmock.NewResult(int64(i), 1))
		result, err := store.CreateTask(tt.description)
		if (err != nil) != tt.expectErr {
			t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
		}

		if result.ID != tt.expected.ID || result.Description != tt.expected.Description || result.Completed != tt.expected.Completed {
			t.Errorf("unexpected result: got %+v, want %+v", result, tt.expected)
		}
	}
}

//
//func TestGetPendingTasks(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("unexpected error: %s", err)
//	}
//
//	store := task.NewTaskStore(db)
//
//	rows := sqlmock.NewRows([]string{"id", "description", "completed"}).
//		AddRow(1, "Buy milk", false).
//		AddRow(2, "Read book", false)
//
//	mock.ExpectQuery("SELECT id, description, completed FROM mytask WHERE completed = FALSE ORDER BY id").
//		WillReturnRows(rows)
//
//	tasks, err := store.GetPendingTasks()
//	if err != nil {
//		t.Errorf("unexpected error: %v", err)
//	}
//	if len(tasks) != 2 {
//		t.Errorf("expected 2 tasks, got %d", len(tasks))
//	}
//}

func TestGetPendingTasks(t *testing.T) {
	tests := []struct {
		name        string
		mockRows    *sqlmock.Rows
		expected    []models.MYTask
		expectErr   bool
		expectedLen int
	}{
		{
			name: "Two pending tasks",
			mockRows: sqlmock.NewRows([]string{"id", "description", "completed"}).
				AddRow(1, "Buy milk", false).
				AddRow(2, "Read book", false),
			expected: []models.MYTask{
				{ID: 1, Description: "Buy milk", Completed: false},
				{ID: 2, Description: "Read book", Completed: false},
			},
			expectErr:   false,
			expectedLen: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("failed to create mock db: %v", err)
			}

			store := task.NewTaskStore(db)

			mock.ExpectQuery("SELECT id, description, completed FROM mytask WHERE completed = FALSE ORDER BY id").
				WillReturnRows(tt.mockRows)

			tasks, err := store.GetPendingTasks()

			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}

			if len(tasks) != tt.expectedLen {
				t.Errorf("expected %d tasks, got %d", tt.expectedLen, len(tasks))
			}

			for i, expectedTask := range tt.expected {
				if tasks[i] != expectedTask {
					t.Errorf("expected task %v, got %v", expectedTask, tasks[i])
				}
			}
		})
	}
}

//func TestGetTaskByID(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("unexpected error: %s", err)
//	}
//
//	store := task.NewTaskStore(db)
//
//	rows := sqlmock.NewRows([]string{"id", "description", "completed"}).
//		AddRow(1, "Finish homework", false)
//
//	mock.ExpectQuery("SELECT id, description, completed FROM mytask WHERE id = ?").
//		WithArgs(1).
//		WillReturnRows(rows)
//
//	result, err := store.GetTaskByID(1)
//	if err != nil {
//		t.Errorf("unexpected error: %v", err)
//	}
//	if result.ID != 1 || result.Description != "Finish homework" || result.Completed != false {
//		t.Errorf("unexpected task: %+v", result)
//	}
//}

// writing testing by making structs
func TestGetTaskByID(t *testing.T) {
	tests := []struct {
		name      string
		inputID   int
		mockRows  *sqlmock.Rows
		expected  models.MYTask
		expectErr bool
	}{
		{
			name:    "Valid Task ID - returns task",
			inputID: 1,
			mockRows: sqlmock.NewRows([]string{"id", "description", "completed"}).
				AddRow(1, "Finish homework", false),
			expected: models.MYTask{
				ID:          1,
				Description: "Finish homework",
				Completed:   false,
			},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("failed to create mock DB: %v", err)
			}

			store := task.NewTaskStore(db)

			mock.ExpectQuery("SELECT id, description, completed FROM mytask WHERE id = ?").
				WithArgs(tt.inputID).
				WillReturnRows(tt.mockRows)

			result, err := store.GetTaskByID(tt.inputID)

			if (err != nil) != tt.expectErr {
				t.Errorf("expected error = %v, got %v", tt.expectErr, err)
			}

			if result != tt.expected {
				t.Errorf("expected task = %+v, got %+v", tt.expected, result)
			}
		})
	}
}

//func TestMarkTaskCompleted(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("unexpected error: %s", err)
//	}
//
//	store := task.NewTaskStore(db)
//
//	mock.ExpectExec("UPDATE mytask SET completed = TRUE WHERE id = ?").
//		WithArgs(1).
//		WillReturnResult(sqlmock.NewResult(0, 1))
//
//	err = store.MarkTaskCompleted(1)
//	if err != nil {
//		t.Errorf("unexpected error: %v", err)
//	}
//}

func TestMarkTaskCompleted(t *testing.T) {
	tests := []struct {
		name      string
		inputID   int
		expectErr bool
		affected  int64
	}{
		{
			name:      "Mark task with ID 1 as completed",
			inputID:   1,
			expectErr: false,
			affected:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("failed to create mock DB: %v", err)
			}

			store := task.NewTaskStore(db)

			mock.ExpectExec("UPDATE mytask SET completed = TRUE WHERE id = ?").
				WithArgs(tt.inputID).
				WillReturnResult(sqlmock.NewResult(0, tt.affected))

			err = store.MarkTaskCompleted(tt.inputID)

			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}
		})
	}
}

//func TestDeleteTask(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("unexpected error: %s", err)
//	}
//
//	store := task.NewTaskStore(db)
//
//	mock.ExpectExec("DELETE FROM mytask WHERE id = ?").
//		WithArgs(1).
//		WillReturnResult(sqlmock.NewResult(0, 1))
//
//	err = store.DeleteTask(1)
//	if err != nil {
//		t.Errorf("unexpected error: %v", err)
//	}
//}

func TestDeleteTask(t *testing.T) {
	tests := []struct {
		name      string
		inputID   int
		expectErr bool
		affected  int64
	}{
		{
			name:      "Delete task with ID 1",
			inputID:   1,
			expectErr: false,
			affected:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("failed to create mock DB: %v", err)
			}

			store := task.NewTaskStore(db)

			mock.ExpectExec("DELETE FROM mytask WHERE id = ?").
				WithArgs(tt.inputID).
				WillReturnResult(sqlmock.NewResult(0, tt.affected))

			err = store.DeleteTask(tt.inputID)

			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}
		})
	}
}
