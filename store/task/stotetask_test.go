package taskstore

import (
	"ThreeLayeredArchitecture/models/task"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/container"
	"testing"
)

func TestAdd(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)
	s := NewTaskStore()

	ctx := &gofr.Context{
		Context:   t.Context(),
		Request:   nil,
		Container: mockContainer,
	}

	tests := []struct {
		name     string
		input    models.Task
		mockFunc func()
		wantErr  bool
	}{
		{
			name: "Successful Add",
			input: models.Task{
				Description: "Test task",
				Completed:   false,
			},
			mockFunc: func() {
				mock.SQL.ExpectExec("INSERT INTO task (description, completed) VALUES (?, ?)").
					WithArgs("Test task", false).
					WillReturnResult(mock.SQL.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "Failed Add",
			input: models.Task{
				Description: "Test task",
				Completed:   false,
			},
			mockFunc: func() {
				mock.SQL.ExpectExec(`INSERT INTO task (description, completed) VALUES (?, ?)`).
					WithArgs("Test task", false).
					WillReturnError(sql.ErrConnDone)
			},
			wantErr: true,
		}, {
			name: "LastInsertId failure",
			input: models.Task{
				Description: "Test task",
				Completed:   false,
			},
			mockFunc: func() {
				_ = mock.SQL.NewResult(0, 1)
				// Wrap it to return an error when LastInsertId is called
				mock.SQL.ExpectExec("INSERT INTO task (description, completed) VALUES (?, ?)").
					WithArgs("Test task", false).
					WillReturnResult(sqlmock.NewErrorResult(errors.New("LastInsertId error")))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()

			got, err := s.Add(ctx, tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				assert.Equal(t, "Test task", got.Description)
			}
		})
	}
}

func TestGetPending(t *testing.T) {
	s := NewTaskStore()

	tests := []struct {
		name     string
		mockFunc func(mockSQL sqlmock.Sqlmock)
		wantErr  bool
		wantLen  int
	}{
		{
			name: "Successful retrieval of pending tasks",
			mockFunc: func(mockSQL sqlmock.Sqlmock) {
				rows := mockSQL.NewRows([]string{"id", "description", "completed"}).
					AddRow(1, "Test task 1", false).
					AddRow(2, "Test task 2", false)
				mockSQL.ExpectQuery(`SELECT id, description, completed FROM task WHERE completed = FALSE ORDER BY id`).
					WillReturnRows(rows)
			},
			wantErr: false,
			wantLen: 2,
		},
		{
			name: "No pending tasks found",
			mockFunc: func(mockSQL sqlmock.Sqlmock) {
				rows := mockSQL.NewRows([]string{"id", "description", "completed"})
				mockSQL.ExpectQuery(`SELECT id, description, completed FROM task WHERE completed = FALSE ORDER BY id`).
					WillReturnRows(rows)
			},
			wantErr: false,
			wantLen: 0,
		},
		{
			name: "Database query error",
			mockFunc: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.ExpectQuery(`SELECT id, description, completed FROM task WHERE completed = FALSE ORDER BY id`).
					WillReturnError(sql.ErrConnDone)
			},
			wantErr: true,
			wantLen: 0,
		},
		{
			name: "Scan error during iteration",
			mockFunc: func(mockSQL sqlmock.Sqlmock) {
				rows := mockSQL.NewRows([]string{"id", "description", "completed"}).
					AddRow(1, "Test task", "invalid_bool")
				mockSQL.ExpectQuery(`SELECT id, description, completed FROM task WHERE completed = FALSE ORDER BY id`).
					WillReturnRows(rows)
			},
			wantErr: true,
			wantLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockContainer, mock := container.NewMockContainer(t)
			ctx := &gofr.Context{
				Context:   t.Context(),
				Request:   nil,
				Container: mockContainer,
			}

			tt.mockFunc(mock.SQL)

			tasks, err := s.GetPending(ctx)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, tasks, tt.wantLen)
			}

			err = mock.SQL.ExpectationsWereMet()
			assert.NoError(t, err, "There were unfulfilled expectations: %s", err)
		})
	}
}

func TestGetByID(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)
	s := NewTaskStore()
	ctx := &gofr.Context{
		Context:   t.Context(),
		Request:   nil,
		Container: mockContainer,
	}

	row := mock.SQL.NewRows([]string{"id", "description", "completed"}).
		AddRow(1, "Sample task", false)

	mock.SQL.ExpectQuery(`SELECT id, description, completed FROM task WHERE id = ?`).
		WithArgs(1).
		WillReturnRows(row)

	got, err := s.GetByID(ctx, 1)

	assert.NoError(t, err)
	assert.Equal(t, 1, got.ID)
	assert.Equal(t, "Sample task", got.Description)
}

func TestMarkComplete(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)
	s := NewTaskStore()
	ctx := &gofr.Context{
		Context:   t.Context(),
		Request:   nil,
		Container: mockContainer,
	}

	mock.SQL.ExpectExec(`UPDATE task SET completed = TRUE WHERE id = ?`).
		WithArgs(1).
		WillReturnResult(mock.SQL.NewResult(1, 1))

	err := s.MarkComplete(ctx, 1)

	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   t.Context(),
		Request:   nil,
		Container: mockContainer,
	}

	mock.SQL.ExpectExec(`DELETE FROM task WHERE id = ?`).
		WithArgs(1).
		WillReturnResult(mock.SQL.NewResult(1, 1))

	s := NewTaskStore()
	err := s.Delete(ctx, 1)

	assert.NoError(t, err)
}
