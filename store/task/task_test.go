package task

import (
	"ThreeLayeredArchitecture/models"
	"context"
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
		Context:   context.Background(),
		Request:   nil,
		Container: mockContainer,
	}

	tests := []struct {
		name     string
		input    models.MYTask
		mockFunc func()
		wantErr  bool
	}{
		{
			name: "Successful Add",
			input: models.MYTask{
				Description: "Test task",
				Completed:   false,
			},
			mockFunc: func() {
				mock.SQL.ExpectExec("INSERT INTO mytask (description, completed) VALUES (?, ?)").
					WithArgs("Test task", false).
					WillReturnResult(mock.SQL.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "Failed Add",
			input: models.MYTask{
				Description: "Test task",
				Completed:   false,
			},
			mockFunc: func() {
				mock.SQL.ExpectExec(`INSERT INTO mytask (description, completed) VALUES (?, ?)`).
					WithArgs("Test task", false).
					WillReturnError(sql.ErrConnDone)
			},
			wantErr: true,
		}, {
			name: "LastInsertId failure",
			input: models.MYTask{
				Description: "Test task",
				Completed:   false,
			},
			mockFunc: func() {
				_ = mock.SQL.NewResult(0, 1)
				// Wrap it to return an error when LastInsertId is called
				mock.SQL.ExpectExec("INSERT INTO mytask (description, completed) VALUES (?, ?)").
					WithArgs("Test task", false).
					WillReturnResult(sqlmock.NewErrorResult(errors.New("LastInsertId error")))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()

			got, err := s.CreateTask(ctx, tt.input)

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
	mockContainer, mock := container.NewMockContainer(t)
	s := NewTaskStore()
	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   nil,
		Container: mockContainer,
	}

	rows := mock.SQL.NewRows([]string{"id", "description", "completed"}).
		AddRow(1, "Test task", false)

	mock.SQL.ExpectQuery(`SELECT id, description, completed FROM mytask WHERE completed = FALSE ORDER BY id`).
		WillReturnRows(rows)

	tasks, err := s.GetPendingTasks(ctx)

	assert.NoError(t, err)
	assert.Len(t, tasks, 1)
	assert.Equal(t, "Test task", tasks[0].Description)
}

func TestGetByID(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)
	s := NewTaskStore()
	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   nil,
		Container: mockContainer,
	}

	row := mock.SQL.NewRows([]string{"id", "description", "completed"}).
		AddRow(1, "Sample task", false)

	mock.SQL.ExpectQuery(`SELECT id, description, completed FROM mytask WHERE id = ?`).
		WithArgs(1).
		WillReturnRows(row)

	got, err := s.GetTaskByID(ctx, 1)

	assert.NoError(t, err)
	assert.Equal(t, 1, got.ID)
	assert.Equal(t, "Sample task", got.Description)
}

func TestMarkComplete(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)
	s := NewTaskStore()
	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   nil,
		Container: mockContainer,
	}

	mock.SQL.ExpectExec(`UPDATE mytask SET completed = TRUE WHERE id = ?`).
		WithArgs(1).
		WillReturnResult(mock.SQL.NewResult(1, 1))

	err := s.MarkTaskCompleted(ctx, 1)

	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   context.Background(),
		Request:   nil,
		Container: mockContainer,
	}

	mock.SQL.ExpectExec(`DELETE FROM mytask WHERE id = ?`).
		WithArgs(1).
		WillReturnResult(mock.SQL.NewResult(1, 1))

	s := NewTaskStore()
	err := s.DeleteTask(ctx, 1)

	assert.NoError(t, err)
}
