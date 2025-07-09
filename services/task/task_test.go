package task

import (
	"context"
	"errors"
	"testing"

	"ThreeLayeredArchitecture/models"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/container"
)

func getTestContext(t *testing.T) *gofr.Context {
	mockContainer, _ := container.NewMockContainer(t)
	return &gofr.Context{
		Context:   context.Background(),
		Request:   nil,
		Container: mockContainer,
	}
}

func TestTaskService_CreateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockTaskStoreInterface(ctrl)
	service := NewTaskService(mockStore)
	ctx := getTestContext(t)

	input := models.MYTask{Description: "New Task", Completed: false}
	expected := models.MYTask{ID: 1, Description: "New Task", Completed: false}

	mockStore.EXPECT().CreateTask(ctx, input).Return(expected, nil)

	result, err := service.CreateTask(ctx, input)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestTaskService_GetPendingTasks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockTaskStoreInterface(ctrl)
	service := NewTaskService(mockStore)
	ctx := getTestContext(t)

	expected := []models.MYTask{
		{ID: 1, Description: "Task1", Completed: false},
		{ID: 2, Description: "Task2", Completed: false},
	}

	mockStore.EXPECT().GetPendingTasks(ctx).Return(expected, nil)

	result, err := service.GetPendingTasks(ctx)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestTaskService_GetTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockTaskStoreInterface(ctrl)
	service := NewTaskService(mockStore)
	ctx := getTestContext(t)

	expected := models.MYTask{ID: 1, Description: "Read", Completed: false}

	mockStore.EXPECT().GetTaskByID(ctx, 1).Return(expected, nil)

	result, err := service.GetTask(ctx, 1)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestTaskService_CompleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockTaskStoreInterface(ctrl)
	service := NewTaskService(mockStore)
	ctx := getTestContext(t)

	mockStore.EXPECT().MarkTaskCompleted(ctx, 1).Return(nil)

	err := service.CompleteTask(ctx, 1)
	assert.Nil(t, err)
}

func TestTaskService_DeleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockTaskStoreInterface(ctrl)
	service := NewTaskService(mockStore)
	ctx := getTestContext(t)

	mockStore.EXPECT().DeleteTask(ctx, 1).Return(errors.New("delete error"))

	err := service.DeleteTask(ctx, 1)
	assert.EqualError(t, err, "delete error")
}
