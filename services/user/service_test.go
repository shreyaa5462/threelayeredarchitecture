package user

import (
	model "ThreeLayeredArchitecture/models/user"
	"errors"
	"go.uber.org/mock/gomock"
	"gofr.dev/pkg/gofr"
	"testing"
)

var ctx *gofr.Context

func TestUserService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockStore := NewMockStore(ctrl)
	service := NewUserService(mockStore)

	testCases := []struct {
		id       int
		desc     string
		input    model.User
		expected model.User
		mockErr  error
	}{
		{
			id:       1,
			desc:     "Success creating user",
			input:    model.User{ID: 1, Name: "Ram"},
			expected: model.User{ID: 1, Name: "Ram"},
			mockErr:  nil,
		},
		{
			id:       2,
			desc:     "Failure creating user",
			input:    model.User{ID: 2, Name: "bhim"},
			expected: model.User{},
			mockErr:  errors.New("unable to create user"),
		},
	}

	for _, tc := range testCases {
		mockStore.EXPECT().CreateUser(ctx, tc.input).Return(tc.expected, tc.mockErr)
		result, err := service.CreateUser(ctx, tc.input)

		if (err == nil && tc.mockErr != nil) || (err != nil && tc.mockErr == nil) {
			t.Errorf("[Test ID %d] %s: expected error %v, got %v", tc.id, tc.desc, tc.mockErr, err)
		}

		if result.ID != tc.expected.ID || result.Name != tc.expected.Name {
			t.Errorf("[Test ID %d] %s: expected user %+v, got %+v", tc.id, tc.desc, tc.expected, result)
		}
	}
}

func TestUserService_GetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockStore(ctrl)
	service := NewUserService(mockStore)

	expectedUsers := []model.User{
		{ID: 1, Name: "Ram"},
		{ID: 2, Name: "Bhim"},
	}

	mockStore.EXPECT().GetAllUsers(ctx).Return(expectedUsers, nil)
	users, err := service.GetAllUsers(ctx)

	if err != nil {
		t.Errorf("[Success case] unexpected error: %v", err)
	}

	if len(users) != len(expectedUsers) {
		t.Errorf("[Success case] expected %d users, got %d", len(expectedUsers), len(users))
	} else {
		for i := range users {
			if users[i].ID != expectedUsers[i].ID || users[i].Name != expectedUsers[i].Name {
				t.Errorf("[Success case] expected user %+v, got %+v", expectedUsers[i], users[i])
			}
		}
	}

	mockErr := errors.New("database error")
	mockStore.EXPECT().GetAllUsers(ctx).Return(nil, mockErr)

	usersFail, errFail := service.GetAllUsers(ctx)
	if errFail == nil || errFail.Error() != mockErr.Error() {
		t.Errorf("[Failure case] expected error %v, got %v", mockErr, errFail)
	}

	if usersFail != nil {
		t.Errorf("[Failure case] expected nil users, got %v", usersFail)
	}
}

func TestUserService_GetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockStore(ctrl)
	service := NewUserService(mockStore)

	testCases := []struct {
		id       int
		desc     string
		inputID  int
		expected model.User
		mockErr  error
	}{
		{
			id:       1,
			desc:     "User found",
			inputID:  1,
			expected: model.User{ID: 1, Name: "Ram"},
			mockErr:  nil,
		},
		{
			id:       2,
			desc:     "User not found",
			inputID:  99,
			expected: model.User{},
			mockErr:  errors.New("user not found"),
		},
	}

	for _, tc := range testCases {
		mockStore.EXPECT().GetUserByID(ctx, tc.inputID).Return(tc.expected, tc.mockErr)
		result, err := service.GetUserByID(ctx, tc.inputID)

		if (err == nil && tc.mockErr != nil) ||
			(err != nil && tc.mockErr == nil) ||
			(err != nil && tc.mockErr != nil && err.Error() != tc.mockErr.Error()) {
			t.Errorf("[Test ID %d] %s: expected error %v, got %v", tc.id, tc.desc, tc.mockErr, err)
		}

		if result.ID != tc.expected.ID || result.Name != tc.expected.Name {
			t.Errorf("[Test ID %d] %s: expected user %+v, got %+v", tc.id, tc.desc, tc.expected, result)
		}
	}
}
