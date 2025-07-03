package task

import (
	task "ThreeLayeredArchitecture/models"
	"context"
	"errors"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/container"
	gofrHttp "gofr.dev/pkg/gofr/http"
	"gofr.dev/pkg/gofr/http/response"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestGetAllTasks(t *testing.T) {
	tests := []struct {
		name        string
		mockTasks   []task.MYTask
		mockError   error
		expectedRes any
		expectedErr error
	}{
		{
			name: "Success - returns tasks",
			mockTasks: []task.MYTask{
				{ID: 1, Description: "Task 1"},
				{ID: 2, Description: "Task 2"},
			},
			mockError:   nil,
			expectedRes: response.Raw{Data: []task.MYTask{{ID: 1, Description: "Task 1"}, {ID: 2, Description: "Task 2"}}},
			expectedErr: nil,
		},
		{
			name:        "Error - service returns error",
			mockTasks:   nil,
			mockError:   errors.New("db error"),
			expectedRes: nil,
			expectedErr: errors.New("db error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockContainer, _ := container.NewMockContainer(t)
			tctx := &gofr.Context{
				Context:   context.Background(),
				Container: mockContainer,
			}

			ctrl := gomock.NewController(t)
			mockSvc := NewMockTaskServiceInterface(ctrl)
			h := NewTaskHandler(mockSvc)

			mockSvc.EXPECT().GetPendingTasks(tctx).Return(tc.mockTasks, tc.mockError)

			req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
			tctx.Request = gofrHttp.NewRequest(req)

			resp, err := h.GetPendingTasks(tctx)

			assert.Equal(t, tc.expectedRes, resp)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestCreateTask(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		expectedTask   task.MYTask
		mockErr        error
		expectBindFail bool
	}{
		{
			name:         "Valid task creation",
			body:         `{"id": 1, "Description": "New Task"}`,
			expectedTask: task.MYTask{ID: 1, Description: "New Task"},
		},
		{
			name:           "Invalid JSON body",
			body:           `{invalid-json}`,
			expectBindFail: true,
		},
		{
			name:         "Service returns error",
			body:         `{"id": 2, "Description": "Fail Task"}`,
			expectedTask: task.MYTask{ID: 2, Description: "Fail Task"},
			mockErr:      errors.New("db failure"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockContainer, _ := container.NewMockContainer(t)
			tctx := &gofr.Context{
				Context:   context.Background(),
				Container: mockContainer,
			}

			ctrl := gomock.NewController(t)
			mockSvc := NewMockTaskServiceInterface(ctrl)
			h := NewTaskHandler(mockSvc)

			req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			tctx.Request = gofrHttp.NewRequest(req)

			if !tt.expectBindFail {
				mockSvc.EXPECT().CreateTask(tctx, tt.expectedTask).Return(tt.expectedTask, tt.mockErr)
			}

			resp, err := h.CreateTask(tctx)

			if tt.expectBindFail {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else if tt.mockErr != nil {
				assert.EqualError(t, err, tt.mockErr.Error())
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, response.Raw{Data: tt.expectedTask}, resp)
			}
		})
	}
}

func TestDeleteTask(t *testing.T) {
	tests := []struct {
		name        string
		param       string
		expectedMsg string
		mockErr     error
		expectFail  bool
	}{
		{
			name:        "Valid Delete",
			param:       "2",
			expectedMsg: "deleted",
		},
		{
			name:       "Invalid ID",
			param:      "abc",
			expectFail: true,
		},
		{
			name:    "Service returns error",
			param:   "5",
			mockErr: errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockContainer, _ := container.NewMockContainer(t)
			tctx := &gofr.Context{
				Context:   context.Background(),
				Container: mockContainer,
			}

			ctrl := gomock.NewController(t)
			mockSvc := NewMockTaskServiceInterface(ctrl)
			h := NewTaskHandler(mockSvc)

			req := httptest.NewRequest(http.MethodDelete, "/tasks", nil)
			req = mux.SetURLVars(req, map[string]string{"id": tt.param})
			tctx.Request = gofrHttp.NewRequest(req)

			if !tt.expectFail {
				id, _ := strconv.Atoi(tt.param)
				mockSvc.EXPECT().DeleteTask(tctx, id).Return(tt.mockErr)
			}
			id, _ := strconv.Atoi(tt.param)

			resp, err := h.DeleteTask(tctx, id)

			if tt.expectFail {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else if tt.mockErr != nil {
				assert.EqualError(t, err, tt.mockErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, response.Raw{Data: tt.expectedMsg}, resp)
			}
		})
	}
}

func TestMarkTaskComplete(t *testing.T) {
	tests := []struct {
		name        string
		param       string
		expectedMsg string
		mockErr     error
		expectFail  bool
	}{
		{
			name:        "Mark complete success",
			param:       "1",
			expectedMsg: "marked complete",
		},
		{
			name:       "Invalid ID param",
			param:      "bad",
			expectFail: true,
		},
		{
			name:    "Service error",
			param:   "9",
			mockErr: errors.New("failed update"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockContainer, _ := container.NewMockContainer(t)
			tctx := &gofr.Context{
				Context:   context.Background(),
				Container: mockContainer,
			}

			ctrl := gomock.NewController(t)
			mockSvc := NewMockTaskServiceInterface(ctrl)
			h := NewTaskHandler(mockSvc)

			req := httptest.NewRequest(http.MethodPatch, "/tasks?id="+tt.param, nil)
			tctx.Request = gofrHttp.NewRequest(req)

			if !tt.expectFail {
				id, _ := strconv.Atoi(tt.param)
				mockSvc.EXPECT().CompleteTask(tctx, id).Return(tt.expectedMsg, tt.mockErr)
			}

			resp, err := h.CompleteTask(tctx)

			if tt.expectFail {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else if tt.mockErr != nil {
				assert.EqualError(t, err, tt.mockErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, response.Raw{Data: tt.expectedMsg}, resp)
			}
		})
	}
}

func TestHandleTaskByID(t *testing.T) {
	tests := []struct {
		name       string
		pathParam  string
		expected   task.MYTask
		mockErr    error
		expectFail bool
	}{
		{
			name:      "Valid Task Fetch",
			pathParam: "2",
			expected:  task.MYTask{ID: 2, Description: "Test"},
		},
		{
			name:       "Invalid ID",
			pathParam:  "x",
			expectFail: true,
		},
		{
			name:      "Service Error",
			pathParam: "5",
			mockErr:   errors.New("not found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockContainer, _ := container.NewMockContainer(t)
			tctx := &gofr.Context{
				Context:   context.Background(),
				Container: mockContainer,
			}

			ctrl := gomock.NewController(t)
			mockSvc := NewMockTaskServiceInterface(ctrl)
			h := NewTaskHandler(mockSvc)

			req := httptest.NewRequest(http.MethodGet, "/tasks/"+tt.pathParam, nil)
			req = mux.SetURLVars(req, map[string]string{"id": tt.pathParam})
			tctx.Request = gofrHttp.NewRequest(req)

			if !tt.expectFail {
				id, _ := strconv.Atoi(tt.pathParam)
				mockSvc.EXPECT().GetTask(tctx, id).Return(tt.expected, tt.mockErr)
			}

			resp, err := h.GetTaskByID(tctx)

			if tt.expectFail {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else if tt.mockErr != nil {
				assert.EqualError(t, err, tt.mockErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, response.Raw{Data: tt.expected}, resp)
			}
		})
	}
}
