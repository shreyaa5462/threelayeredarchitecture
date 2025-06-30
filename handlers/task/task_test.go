package task_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"ThreeLayeredArchitecture/handlers/task"
	"ThreeLayeredArchitecture/models"
)

type MockTaskService struct{}

func (m *MockTaskService) CreateTask(description string) (models.MYTask, error) {
	return models.MYTask{ID: 1, Description: description, Completed: false}, nil
}

func (m *MockTaskService) GetPendingTasks() ([]models.MYTask, error) {
	return []models.MYTask{
		{ID: 1, Description: "Mock Task 1", Completed: false},
		{ID: 2, Description: "Mock Task 2", Completed: false},
	}, nil
}

func (m *MockTaskService) GetTask(id int) (models.MYTask, error) {
	return models.MYTask{ID: id, Description: fmt.Sprintf("Mock Task %d", id), Completed: false}, nil
}

func (m *MockTaskService) CompleteTask(id int) error {
	return nil
}

func (m *MockTaskService) DeleteTask(id int) error {
	return nil
}

func setup() *task.TaskHandler {
	return task.NewTaskHandler(&MockTaskService{})
}

func TestHandleTasks_POST(t *testing.T) {
	handler := setup()
	taskData := models.MYTask{Description: "Test Task"}
	body, _ := json.Marshal(taskData)

	req := httptest.NewRequest(http.MethodPost, "/task", bytes.NewReader(body))
	w := httptest.NewRecorder()
	handler.HandleTasks(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestHandleTasks_GET(t *testing.T) {
	handler := setup()
	req := httptest.NewRequest(http.MethodGet, "/task", nil)
	w := httptest.NewRecorder()
	handler.HandleTasks(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestHandleTasks_PUT(t *testing.T) {
	handler := setup()
	req := httptest.NewRequest(http.MethodPut, "/task?id=1", nil)
	w := httptest.NewRecorder()
	handler.HandleTasks(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestHandleTasks_DELETE(t *testing.T) {
	handler := setup()
	req := httptest.NewRequest(http.MethodDelete, "/task?id=1", nil)
	w := httptest.NewRecorder()
	handler.HandleTasks(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestHandleTaskByID(t *testing.T) {
	handler := setup()

	req := httptest.NewRequest(http.MethodGet, "/task/1", nil)
	w := httptest.NewRecorder()

	handler.HandleTaskByID(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}
