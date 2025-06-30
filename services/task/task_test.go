package task

import (
	"testing"

	"ThreeLayeredArchitecture/models"
)

type MockTaskStore struct{}

func (m *MockTaskStore) CreateTask(description string) (models.MYTask, error) {
	return models.MYTask{ID: 1, Description: description, Completed: false}, nil
}

func (m *MockTaskStore) GetPendingTasks() ([]models.MYTask, error) {
	return []models.MYTask{
		{ID: 1, Description: "Test Task", Completed: false},
	}, nil
}

func (m *MockTaskStore) GetTaskByID(id int) (models.MYTask, error) {
	return models.MYTask{ID: id, Description: "Test Task", Completed: false}, nil
}

func (m *MockTaskStore) MarkTaskCompleted(id int) error {
	return nil
}

func (m *MockTaskStore) DeleteTask(id int) error {
	return nil
}

func TestCreateTask(t *testing.T) {
	service := NewTaskService(&MockTaskStore{})
	task, err := service.CreateTask("Sample Task")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if task.Description != "Sample Task" {
		t.Errorf("Expected task description 'Sample Task', got '%s'", task.Description)
	}
}

func TestGetPendingTasks(t *testing.T) {
	service := NewTaskService(&MockTaskStore{})
	tasks, err := service.GetPendingTasks()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(tasks) == 0 {
		t.Errorf("Expected at least 1 task, got 0")
	}
}

func TestGetTask(t *testing.T) {
	service := NewTaskService(&MockTaskStore{})
	task, err := service.GetTask(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if task.ID != 1 {
		t.Errorf("Expected ID 1, got %d", task.ID)
	}
}

func TestCompleteTask(t *testing.T) {
	service := NewTaskService(&MockTaskStore{})
	err := service.CompleteTask(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestDeleteTask(t *testing.T) {
	service := NewTaskService(&MockTaskStore{})
	err := service.DeleteTask(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
