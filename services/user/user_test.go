package user

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

func getMockContext(t *testing.T) *gofr.Context {
	mockContainer, _ := container.NewMockContainer(t)
	return &gofr.Context{
		Context:   context.Background(),
		Request:   nil,
		Container: mockContainer,
	}
}

func TestUserService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockuserStoreInterface(ctrl)
	service := NewUserService(mockStore)
	ctx := getMockContext(t)

	input := models.User{Name: "Shreya"}
	expected := models.User{ID: 1, Name: "Shreya"}

	mockStore.EXPECT().CreateUser(ctx, input).Return(expected, nil)

	result, err := service.CreateUser(ctx, input)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestUserService_GetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockuserStoreInterface(ctrl)
	service := NewUserService(mockStore)
	ctx := getMockContext(t)

	expected := models.User{ID: 1, Name: "Shreya"}

	mockStore.EXPECT().GetUserByID(ctx, 1).Return(expected, nil)

	result, err := service.GetUserByID(ctx, 1)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestUserService_GetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockuserStoreInterface(ctrl)
	service := NewUserService(mockStore)
	ctx := getMockContext(t)

	expected := []models.User{
		{ID: 1, Name: "Shreya"},
		{ID: 2, Name: "Aman"},
	}

	mockStore.EXPECT().GetAllUsers(ctx).Return(expected, nil)

	result, err := service.GetAllUsers(ctx)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestUserService_Errors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockuserStoreInterface(ctrl)
	service := NewUserService(mockStore)
	ctx := getMockContext(t)

	t.Run("CreateUser error", func(t *testing.T) {
		mockStore.EXPECT().CreateUser(ctx, models.User{Name: "X"}).Return(models.User{}, errors.New("insert error"))
		_, err := service.CreateUser(ctx, models.User{Name: "X"})
		assert.EqualError(t, err, "insert error")
	})

	t.Run("GetUserByID error", func(t *testing.T) {
		mockStore.EXPECT().GetUserByID(ctx, 100).Return(models.User{}, errors.New("not found"))
		_, err := service.GetUserByID(ctx, 100)
		assert.EqualError(t, err, "not found")
	})

	t.Run("GetAllUsers error", func(t *testing.T) {
		mockStore.EXPECT().GetAllUsers(ctx).Return(nil, errors.New("db error"))
		_, err := service.GetAllUsers(ctx)
		assert.EqualError(t, err, "db error")
	})
}
