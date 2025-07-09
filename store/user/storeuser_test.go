package userstore

import (
	"ThreeLayeredArchitecture/models/user"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/container"
)

func TestCreateUser(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   t.Context(),
		Request:   nil,
		Container: mockContainer,
	}

	tests := []struct {
		name     string
		user     models.User
		mockfunc func()
		expected models.User
		err      error
	}{
		{
			name: "Successful create user",
			user: models.User{Name: "John Doe"},
			mockfunc: func() {
				mock.SQL.ExpectExec("INSERT INTO user (name) VALUES (?)").
					WithArgs("John Doe").
					WillReturnResult(mock.SQL.NewResult(1, 1))
			},
			expected: models.User{ID: 1, Name: "John Doe"},
			err:      nil,
		},
		{
			name: "Failed create user - database error",
			user: models.User{Name: "Jane Doe"},
			mockfunc: func() {
				mock.SQL.ExpectExec("INSERT INTO user (name) VALUES (?)").
					WithArgs("Jane Doe").
					WillReturnError(sql.ErrConnDone)
			},
			expected: models.User{},
			err:      sql.ErrConnDone,
		},
		{
			name: "Create user with empty name",
			user: models.User{Name: ""},
			mockfunc: func() {
				mock.SQL.ExpectExec("INSERT INTO user (name) VALUES (?)").
					WithArgs("").
					WillReturnResult(mock.SQL.NewResult(2, 1))
			},
			expected: models.User{ID: 2, Name: ""},
			err:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc()

			userStore := NewUserStore()
			result, err := userStore.CreateUser(ctx, tt.user)

			if !assert.Equal(t, tt.err, err) {
				t.Errorf("%v: error = %v, wantErr %v", tt.name, err, tt.err)
			}

			if !assert.Equal(t, tt.expected, result) {
				t.Errorf("%v: \nExpected = %v\nGot = %v", tt.name, tt.expected, result)
			}
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   t.Context(),
		Request:   nil,
		Container: mockContainer,
	}

	successRows := mock.SQL.NewRows([]string{"id", "name"}).
		AddRow(1, "John Doe").
		AddRow(2, "Jane Smith").
		AddRow(3, "Bob Johnson")

	emptyRows := mock.SQL.NewRows([]string{"id", "name"})

	tests := []struct {
		name     string
		mockfunc func()
		expected []models.User
		err      error
	}{
		{
			name: "Successful get all users",
			mockfunc: func() {
				mock.SQL.ExpectQuery("SELECT id, name FROM user").
					WillReturnRows(successRows)
			},
			expected: []models.User{
				{ID: 1, Name: "John Doe"},
				{ID: 2, Name: "Jane Smith"},
				{ID: 3, Name: "Bob Johnson"},
			},
			err: nil,
		},
		{
			name: "Get all users - empty result",
			mockfunc: func() {
				mock.SQL.ExpectQuery("SELECT id, name FROM user").
					WillReturnRows(emptyRows)
			},
			expected: nil,
			err:      nil,
		},
		{
			name: "Database query error",
			mockfunc: func() {
				// Simulate an error when ctx.SQL.Query is called
				mock.SQL.ExpectQuery("SELECT id, name FROM user").
					WillReturnError(sql.ErrConnDone)
			},
			expected: nil, // The function returns nil, err on initial query error
			err:      sql.ErrConnDone,
		},
		{
			name: "Scan error during iteration",
			mockfunc: func() {
				// Simulate a scan error by providing a value that cannot be scanned into an int
				mock.SQL.ExpectQuery("SELECT id, name FROM user").
					WillReturnRows(mock.SQL.NewRows([]string{"id", "name"}).
						AddRow("not_an_integer", "John Doe")) // "not_an_integer" cannot be scanned into user.ID (int)
			},
			expected: nil, // The function returns nil, err on scan error
			err:      nil, // Set to nil here to avoid the compile error
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc()

			userStore := NewUserStore()
			result, err := userStore.GetAllUsers(ctx)

			if tt.name == "Scan error during iteration" {
				assert.NotNil(t, err, "Expected an error for scan failure")

				assert.Nil(t, result, "Expected nil result on scan error")

				if err != nil {
					assert.Contains(t, err.Error(), "Scan", "Expected error message to contain 'Scan'")
				}
			} else {
				if !assert.Equal(t, tt.err, err) {
					t.Errorf("%v: error = %v, wantErr %v", tt.name, err, tt.err)
				}

				if !assert.Equal(t, tt.expected, result) {
					t.Errorf("%v: \nExpected = %v\nGot = %v", tt.name, tt.expected, result)
				}
			}
		})
	}
}

func TestGetUserByID(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)

	ctx := &gofr.Context{
		Context:   t.Context(),
		Request:   nil,
		Container: mockContainer,
	}

	successRow := mock.SQL.NewRows([]string{"id", "name"}).
		AddRow(1, "John Doe")

	tests := []struct {
		name     string
		userID   int
		mockfunc func()
		expected models.User
		err      error
	}{
		{
			name:   "Successful get user by ID",
			userID: 1,
			mockfunc: func() {
				mock.SQL.ExpectQuery("SELECT id, name FROM user WHERE id = ?").
					WithArgs(1).
					WillReturnRows(successRow)
			},
			expected: models.User{ID: 1, Name: "John Doe"},
			err:      nil,
		},
		{
			name:   "User not found",
			userID: 999,
			mockfunc: func() {
				mock.SQL.ExpectQuery("SELECT id, name FROM user WHERE id = ?").
					WithArgs(999).
					WillReturnError(sql.ErrNoRows)
			},
			expected: models.User{},
			err:      sql.ErrNoRows,
		},
		{
			name:   "Database connection error",
			userID: 1,
			mockfunc: func() {
				mock.SQL.ExpectQuery("SELECT id, name FROM user WHERE id = ?").
					WithArgs(1).
					WillReturnError(sql.ErrConnDone)
			},
			expected: models.User{},
			err:      sql.ErrConnDone,
		},
		{
			name:   "Get user with ID 0",
			userID: 0,
			mockfunc: func() {
				mock.SQL.ExpectQuery("SELECT id, name FROM user WHERE id = ?").
					WithArgs(0).
					WillReturnError(sql.ErrNoRows)
			},
			expected: models.User{},
			err:      sql.ErrNoRows,
		},
		{
			name:   "Get user with negative ID",
			userID: -1,
			mockfunc: func() {
				mock.SQL.ExpectQuery("SELECT id, name FROM user WHERE id = ?").
					WithArgs(-1).
					WillReturnError(sql.ErrNoRows)
			},
			expected: models.User{},
			err:      sql.ErrNoRows,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockfunc()

			userStore := NewUserStore()
			result, err := userStore.GetUserByID(ctx, tt.userID)

			if !assert.Equal(t, tt.err, err) {
				t.Errorf("%v: error = %v, wantErr %v", tt.name, err, tt.err)
			}

			if !assert.Equal(t, tt.expected, result) {
				t.Errorf("%v: \nExpected = %v\nGot = %v", tt.name, tt.expected, result)
			}
		})
	}
}
