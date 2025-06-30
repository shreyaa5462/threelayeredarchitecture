package user_test

import (
	"errors"
	"testing"

	"ThreeLayeredArchitecture/models"
	userservice "ThreeLayeredArchitecture/services/user"
)

// Mock store that satisfies the userStoreInterface
type MockUserStore struct{}

func (m *MockUserStore) GetAllUsers() ([]models.User, error) {
	return []models.User{
		{ID: 1, Name: "Shreya"},
		{ID: 2, Name: "Aman"},
	}, nil
}

func (m *MockUserStore) GetUserByID(id int) (models.User, error) {
	if id == 1 {
		return models.User{ID: 1, Name: "Shreya"}, nil
	}
	return models.User{}, errors.New("user not found")
}

func (m *MockUserStore) CreateUser(user models.User) (models.User, error) {
	user.ID = 3
	return user, nil
}

func TestGetAllUsers(t *testing.T) {
	mockStore := &MockUserStore{}
	service := userservice.NewUserService(mockStore)

	users, err := service.GetAllUsers()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}
}

func TestGetUserByID_Found(t *testing.T) {
	mockStore := &MockUserStore{}
	service := userservice.NewUserService(mockStore)

	user, err := service.GetUserByID(1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if user.Name != "Shreya" {
		t.Errorf("Expected name Shreya, got %s", user.Name)
	}
}

func TestGetUserByID_NotFound(t *testing.T) {
	mockStore := &MockUserStore{}
	service := userservice.NewUserService(mockStore)

	_, err := service.GetUserByID(99)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestCreateUser(t *testing.T) {
	mockStore := &MockUserStore{}
	service := userservice.NewUserService(mockStore)

	input := models.User{Name: "NewUser"}
	createdUser, err := service.CreateUser(input)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if createdUser.ID != 3 {
		t.Errorf("Expected user ID 3, got %d", createdUser.ID)
	}
}
