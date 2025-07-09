package user

import (
	user "ThreeLayeredArchitecture/models/user"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/container"
	gofrHttp "gofr.dev/pkg/gofr/http"
	"gofr.dev/pkg/gofr/http/response"
)

func TestCreateUser(t *testing.T) {
	mockContainer, _ := container.NewMockContainer(t)
	tctx := &gofr.Context{
		Context:   t.Context(),
		Request:   nil,
		Container: mockContainer,
	}

	tests := []struct {
		name           string
		inputBody      string
		existingUser   user.User
		mockErr        error
		expectedResult any
		expectedErr    error
		useMock        bool
	}{
		{
			name:           "Successful CreateUser",
			inputBody:      `{"id":10,"name":"Alice"}`,
			existingUser:   user.User{ID: 10, Name: "Alice"},
			mockErr:        nil,
			expectedResult: response.Raw{Data: user.User{ID: 10, Name: "Alice"}},
			expectedErr:    nil,
			useMock:        true,
		},
		{
			name:        "Failed Binding",
			inputBody:   `{Alice}`,
			expectedErr: &json.SyntaxError{},
			useMock:     false,
		},
		{
			name:         "Service Error",
			inputBody:    `{"id":20,"name":"Bob"}`,
			existingUser: user.User{ID: 20, Name: "Bob"},
			mockErr:      errors.New("service failure"),
			expectedErr:  errors.New("service failure"),
			useMock:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockSvc := NewMockUserService(ctrl)
			h := NewUserHandler(mockSvc)

			req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(tt.inputBody))
			req.Header.Set("Content-Type", "application/json")
			reqCtx := gofrHttp.NewRequest(req)
			tctx.Request = reqCtx

			if tt.useMock {
				mockSvc.EXPECT().CreateUser(tctx, tt.existingUser).Return(tt.existingUser, tt.mockErr)
			}

			val, err := h.CreateUser(tctx)

			if tt.expectedErr != nil {
				switch tt.expectedErr.(type) {
				case *json.SyntaxError:
					assert.IsType(t, &json.SyntaxError{}, err)
				default:
					assert.EqualError(t, tt.expectedErr, err.Error())
				}

				assert.Nil(t, val)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, val)
			}
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	mockContainer, _ := container.NewMockContainer(t)
	tctx := &gofr.Context{
		Context:   t.Context(),
		Request:   nil,
		Container: mockContainer,
	}

	tests := []struct {
		name           string
		mockResult     []user.User
		mockErr        error
		expectedResult any
		expectedErr    error
	}{
		{
			name:           "Successful GetAllUsers",
			mockResult:     []user.User{{ID: 1, Name: "A"}, {ID: 2, Name: "B"}},
			mockErr:        nil,
			expectedResult: response.Raw{Data: []user.User{{ID: 1, Name: "A"}, {ID: 2, Name: "B"}}},
			expectedErr:    nil,
		},
		{
			name:        "Service Error",
			mockErr:     errors.New("db error"),
			expectedErr: errors.New("db error"),
			// no result expected on error
			expectedResult: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockSvc := NewMockUserService(ctrl)
			h := NewUserHandler(mockSvc)

			req := httptest.NewRequest(http.MethodGet, "/users", http.NoBody)
			req.Header.Set("Content-Type", "application/json")
			tctx.Request = gofrHttp.NewRequest(req)

			mockSvc.EXPECT().GetAllUsers(tctx).Return(tt.mockResult, tt.mockErr)

			val, err := h.GetAllUsers(tctx)

			if tt.expectedErr != nil {
				assert.EqualError(t, tt.expectedErr, err.Error())
				assert.Nil(t, val)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, val)
			}
		})
	}
}

func TestGetUserByID(t *testing.T) {
	mockContainer, _ := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   t.Context(),
		Request:   nil,
		Container: mockContainer,
	}

	tests := []struct {
		name           string
		pathParam      string
		mockUser       user.User
		mockErr        error
		expectedResult any
		expectedErr    error
		useMock        bool
	}{
		{
			name:           "Successful GetUserByID",
			pathParam:      "5",
			mockUser:       user.User{ID: 5, Name: "Eve"},
			expectedResult: response.Raw{Data: user.User{ID: 5, Name: "Eve"}},
			expectedErr:    nil,
			useMock:        true,
		},
		{
			name:        "Invalid ID",
			pathParam:   "x",
			expectedErr: &strconv.NumError{},
			useMock:     false,
		},
		{
			name:        "Service Error",
			pathParam:   "7",
			mockErr:     errors.New("not found"),
			expectedErr: errors.New("not found"),
			useMock:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockSvc := NewMockUserService(ctrl)
			h := NewUserHandler(mockSvc)

			req := httptest.NewRequest(http.MethodGet, "/users/{id}", http.NoBody)
			req = mux.SetURLVars(req, map[string]string{"id": tt.pathParam})
			ctx.Request = gofrHttp.NewRequest(req)

			if tt.useMock {
				id, _ := strconv.Atoi(tt.pathParam)
				mockSvc.EXPECT().GetUserByID(ctx, id).Return(tt.mockUser, tt.mockErr)
			}

			val, err := h.GetUserbyID(ctx)

			if tt.expectedErr != nil {
				switch tt.expectedErr.(type) {
				case *strconv.NumError:
					assert.IsType(t, &strconv.NumError{}, err)
				default:
					assert.EqualError(t, tt.expectedErr, err.Error())
				}

				assert.Nil(t, val)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, val)
			}
		})
	}
}
