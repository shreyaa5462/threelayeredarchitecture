//package task_test
//
//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//
//	"ThreeLayeredArchitecture/handlers/task"
//	"ThreeLayeredArchitecture/models"
//)
//
//type MockTaskService struct{}
//
//func (m *MockTaskService) CreateTask(description string) (models.MYTask, error) {
//	return models.MYTask{ID: 1, Description: description, Completed: false}, nil
//}
//
//func (m *MockTaskService) GetPendingTasks() ([]models.MYTask, error) {
//	return []models.MYTask{
//		{ID: 1, Description: "Mock Task 1", Completed: false},
//		{ID: 2, Description: "Mock Task 2", Completed: false},
//	}, nil
//}
//
//func (m *MockTaskService) GetTask(id int) (models.MYTask, error) {
//	return models.MYTask{ID: id, Description: fmt.Sprintf("Mock Task %d", id), Completed: false}, nil
//}
//
//func (m *MockTaskService) CompleteTask(id int) error {
//	return nil
//}
//
//func (m *MockTaskService) DeleteTask(id int) error {
//	return nil
//}
//
//func setup() *task.TaskHandler {
//	return task.NewTaskHandler(&MockTaskService{})
//}
//
//func TestHandleTasks_POST(t *testing.T) {
//	handler := setup()
//	taskData := models.MYTask{Description: "Test Task"}
//	body, _ := json.Marshal(taskData)
//
//	req := httptest.NewRequest(http.MethodPost, "/task", bytes.NewReader(body))
//	w := httptest.NewRecorder()
//	handler.HandleTasks(w, req)
//
//	if w.Code != http.StatusOK {
//		t.Errorf("Expected status 200, got %d", w.Code)
//	}
//}
//
//func TestHandleTasks_GET(t *testing.T) {
//	handler := setup()
//	req := httptest.NewRequest(http.MethodGet, "/task", nil)
//	w := httptest.NewRecorder()
//	handler.HandleTasks(w, req)
//
//	if w.Code != http.StatusOK {
//		t.Errorf("Expected status 200, got %d", w.Code)
//	}
//}
//
//func TestHandleTasks_PUT(t *testing.T) {
//	handler := setup()
//	req := httptest.NewRequest(http.MethodPut, "/task?id=1", nil)
//	w := httptest.NewRecorder()
//	handler.HandleTasks(w, req)
//
//	if w.Code != http.StatusOK {
//		t.Errorf("Expected status 200, got %d", w.Code)
//	}
//}
//
//func TestHandleTasks_DELETE(t *testing.T) {
//	handler := setup()
//	req := httptest.NewRequest(http.MethodDelete, "/task?id=1", nil)
//	w := httptest.NewRecorder()
//	handler.HandleTasks(w, req)
//
//	if w.Code != http.StatusOK {
//		t.Errorf("Expected status 200, got %d", w.Code)
//	}
//}
//
//func TestHandleTaskByID(t *testing.T) {
//	handler := setup()
//
//	req := httptest.NewRequest(http.MethodGet, "/task/1", nil)
//	w := httptest.NewRecorder()
//
//	handler.HandleTaskByID(w, req)
//
//	if w.Code != http.StatusOK {
//		t.Errorf("Expected status 200, got %d", w.Code)
//	}
//}

package task

import (
	"ThreeLayeredArchitecture/models"
	"bytes"
	"encoding/json"
	"errors"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTaskHandler_HandleTasks(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := NewMockTaskServiceInterface(ctrl)
	handler := NewTaskHandler(mockService)

	type testCase struct {
		id         int
		desc       string
		method     string
		url        string
		body       interface{}
		mockFunc   func()
		wantStatus int
	}

	testCases := []testCase{
		{
			id:     1,
			desc:   "POST valid input",
			method: http.MethodPost,
			url:    "/tasks",
			body:   models.MYTask{Description: "Clean Room"},
			mockFunc: func() {
				mockService.EXPECT().CreateTask("Clean Room").
					Return(models.MYTask{ID: 1, Description: "Clean Room", Completed: false}, nil)
			},
			wantStatus: http.StatusOK,
		},
		{
			id:     2,
			desc:   "GET tasks",
			method: http.MethodGet,
			url:    "/tasks",
			mockFunc: func() {
				mockService.EXPECT().GetPendingTasks().
					Return([]models.MYTask{
						{ID: 1, Description: "Task1", Completed: false},
					}, nil)
			},
			wantStatus: http.StatusOK,
		},
		{
			id:     3,
			desc:   "PUT complete task",
			method: http.MethodPut,
			url:    "/tasks?id=1",
			mockFunc: func() {
				mockService.EXPECT().CompleteTask(1).Return(nil)
			},
			wantStatus: http.StatusOK,
		},
		{
			id:     4,
			desc:   "DELETE task failure",
			method: http.MethodDelete,
			url:    "/tasks?id=2",
			mockFunc: func() {
				var err = errors.New("fail")
				mockService.EXPECT().DeleteTask(2).Return(err)
			},
			wantStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		var reqBody *bytes.Buffer

		if tc.body != nil {
			bodyBytes, _ := json.Marshal(tc.body)
			reqBody = bytes.NewBuffer(bodyBytes)
		} else {
			reqBody = &bytes.Buffer{}
		}

		req := httptest.NewRequest(tc.method, tc.url, reqBody)
		w := httptest.NewRecorder()

		if tc.mockFunc != nil {
			tc.mockFunc()
		}

		handler.HandleTasks(w, req)

		if w.Code != tc.wantStatus {
			t.Errorf("Test ID %d - %s: expected status %d, got %d", tc.id, tc.desc, tc.wantStatus, w.Code)
		}
	}
}

func TestTaskHandler_HandleTaskByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := NewMockTaskServiceInterface(ctrl)
	handler := NewTaskHandler(mockService)

	type testCase struct {
		id         int
		desc       string
		method     string
		url        string
		mockFunc   func()
		wantStatus int
	}

	testCases := []testCase{
		{
			id:     1,
			desc:   "GET valid task",
			method: http.MethodGet,
			url:    "/task?id=1",
			mockFunc: func() {
				mockService.EXPECT().GetTask(1).Return(models.MYTask{ID: 1, Description: "Study"}, nil)
			},
			wantStatus: http.StatusOK,
		},
		{
			id:         2,
			desc:       "Invalid method used",
			method:     http.MethodPost,
			url:        "/task?id=1",
			mockFunc:   nil,
			wantStatus: http.StatusMethodNotAllowed,
		},
		{
			id:         3,
			desc:       "Invalid ID format",
			method:     http.MethodGet,
			url:        "/task?id=abc",
			mockFunc:   nil,
			wantStatus: http.StatusBadRequest,
		},
		{
			id:     4,
			desc:   "Task not found",
			method: http.MethodGet,
			url:    "/task?id=5",
			mockFunc: func() {
				mockService.EXPECT().GetTask(5).Return(models.MYTask{}, errors.New("not found"))
			},
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest(tc.method, tc.url, nil)
		w := httptest.NewRecorder()

		if tc.mockFunc != nil {
			tc.mockFunc()
		}

		handler.HandleTaskByID(w, req)

		if w.Code != tc.wantStatus {
			t.Errorf("Test ID %d - %s: expected status %d, got %d", tc.id, tc.desc, tc.wantStatus, w.Code)
		}
	}
}
