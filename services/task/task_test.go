package task

import (
	"ThreeLayeredArchitecture/models"
	"errors"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestTaskService_CreateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockTaskStoreInterface(ctrl)
	service := NewTaskService(mockStore)

	testCases := []struct {
		id       int
		desc     string
		input    string
		expected models.MYTask
		mockErr  error
	}{
		{1, "Success case", "Clean Room", models.MYTask{ID: 1, Description: "Clean Room", Completed: false}, nil},
		{2, "Failure case", "Cook Food", models.MYTask{}, errors.New("failed to create task")},
	}

	for _, tc := range testCases {
		mockStore.EXPECT().CreateTask(tc.input).Return(tc.expected, tc.mockErr)
		result, err := service.CreateTask(tc.input)

		if (err == nil && tc.mockErr != nil) || (err != nil && tc.mockErr == nil) {
			t.Errorf("[Test ID %d] %s: expected error %v, got %v", tc.id, tc.desc, tc.mockErr, err)
		}

		if result != tc.expected {
			t.Errorf("[Test ID %d] %s: expected result %+v, got %+v", tc.id, tc.desc, tc.expected, result)
		}
	}
}

func TestTaskService_GetPendingTasks(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockTaskStoreInterface(ctrl)
	service := NewTaskService(mockStore)

	expected := []models.MYTask{
		{ID: 1, Description: "Clean Room", Completed: false},
		{ID: 2, Description: "Wash Clothes", Completed: false},
	}

	mockStore.EXPECT().GetPendingTasks().Return(expected, nil)
	tasks, err := service.GetPendingTasks()

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(tasks) != len(expected) {
		t.Errorf("Expected %d tasks, got %d", len(expected), len(tasks))
	}

	for i := range expected {
		if tasks[i] != expected[i] {
			t.Errorf("Mismatch at index %d: expected %+v, got %+v", i, expected[i], tasks[i])
		}
	}

	mockErr := errors.New("db error")
	mockStore.EXPECT().GetPendingTasks().Return(nil, mockErr)
	result, err := service.GetPendingTasks()

	if err == nil || err.Error() != mockErr.Error() {
		t.Errorf("Expected error %v, got %v", mockErr, err)
	}

	if result != nil {
		t.Errorf("Expected nil result, got %+v", result)
	}
}

func TestTaskService_GetTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockTaskStoreInterface(ctrl)
	service := NewTaskService(mockStore)

	testCases := []struct {
		id       int
		desc     string
		inputID  int
		expected models.MYTask
		mockErr  error
	}{
		{1, "Task found", 1, models.MYTask{ID: 1, Description: "Clean", Completed: false}, nil},
		{2, "Task not found", 99, models.MYTask{}, errors.New("task not found")},
	}

	for _, tc := range testCases {
		mockStore.EXPECT().GetTaskByID(tc.inputID).Return(tc.expected, tc.mockErr)
		result, err := service.GetTask(tc.inputID)

		if (err == nil && tc.mockErr != nil) || (err != nil && tc.mockErr == nil) {
			t.Errorf("[Test ID %d] %s: expected error %v, got %v", tc.id, tc.desc, tc.mockErr, err)
		}

		if result != tc.expected {
			t.Errorf("[Test ID %d] %s: expected %+v, got %+v", tc.id, tc.desc, tc.expected, result)
		}
	}
}

func TestTaskService_CompleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockTaskStoreInterface(ctrl)
	service := NewTaskService(mockStore)

	mockStore.EXPECT().MarkTaskCompleted(1).Return(nil)

	err := service.CompleteTask(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	mockErr := errors.New("failed to update task")
	mockStore.EXPECT().MarkTaskCompleted(2).Return(mockErr)

	err = service.CompleteTask(2)

	if err == nil || err.Error() != mockErr.Error() {
		t.Errorf("Expected error %v, got %v", mockErr, err)
	}
}

func TestTaskService_DeleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockTaskStoreInterface(ctrl)
	service := NewTaskService(mockStore)

	mockStore.EXPECT().DeleteTask(1).Return(nil)
	err := service.DeleteTask(1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	mockErr := errors.New("failed to delete task")
	mockStore.EXPECT().DeleteTask(2).Return(mockErr)
	err = service.DeleteTask(2)

	if err == nil || err.Error() != mockErr.Error() {
		t.Errorf("Expected error %v, got %v", mockErr, err)
	}
}
