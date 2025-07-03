// // import (
// //
// //	"errors"
// //	"testing"
// //
// //	"ThreeLayeredArchitecture/models"
// //	userservice "ThreeLayeredArchitecture/services/user"
// //
// // )
// //
// // // Mock store that satisfies the userStoreInterface
// // type MockUserStore struct{}
// //
// //	func (m *MockUserStore) GetAllUsers() ([]models.User, error) {
// //		return []models.User{
// //			{ID: 1, Name: "Shreya"},
// //			{ID: 2, Name: "Singh"},
// //		}, nil
// //	}
// //
// //	func (m *MockUserStore) GetUserByID(id int) (models.User, error) {
// //		if id == 1 {
// //			return models.User{ID: 1, Name: "Shreya"}, nil
// //		}
// //		return models.User{}, errors.New("user not found")
// //	}
// //
// //	func (m *MockUserStore) CreateUser(user models.User) (models.User, error) {
// //		user.ID = 3
// //		return user, nil
// //	}
// //
// //	func TestGetAllUsers(t *testing.T) {
// //		mockStore := &MockUserStore{}
// //		service := userservice.NewUserService(mockStore)
// //
// //		users, err := service.GetAllUsers()
// //		if err != nil {
// //			t.Fatalf("Expected no error, got %v", err)
// //		}
// //		if len(users) != 2 {
// //			t.Errorf("Expected 2 users, got %d", len(users))
// //		}
// //	}
// //
// //	func TestGetUserByID_Found(t *testing.T) {
// //		mockStore := &MockUserStore{}
// //		service := userservice.NewUserService(mockStore)
// //
// //		user, err := service.GetUserByID(1)
// //		if err != nil {
// //			t.Fatalf("Expected no error, got %v", err)
// //		}
// //		if user.Name != "Shreya" {
// //			t.Errorf("Expected name Shreya, got %s", user.Name)
// //		}
// //	}
// //
// //	func TestGetUserByID_NotFound(t *testing.T) {
// //		mockStore := &MockUserStore{}
// //		service := userservice.NewUserService(mockStore)
// //
// //		_, err := service.GetUserByID(99)
// //		if err == nil {
// //			t.Error("Expected error, got nil")
// //		}
// //	}
// //
// //	func TestCreateUser(t *testing.T) {
// //		mockStore := &MockUserStore{}
// //		service := userservice.NewUserService(mockStore)
// //
// //		input := models.User{Name: "NewUser"}
// //		createdUser, err := service.CreateUser(input)
// //		if err != nil {
// //			t.Fatalf("Expected no error, got %v", err)
// //		}
// //		if createdUser.ID != 3 {
// //			t.Errorf("Expected user ID 3, got %d", createdUser.ID)
// //		}
// //	}
package user_test

import (
	"ThreeLayeredArchitecture/services/user"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"ThreeLayeredArchitecture/models"

	"go.uber.org/mock/gomock"
)

func TestGetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := user.NewMockuserStoreInterface(ctrl)
	testCases := []struct {
		name      string
		mockResp  []models.User
		mockErr   error
		expected  []models.User
		expectErr error
	}{
		{
			name: "returns users successfully",
			mockResp: []models.User{
				{ID: 1, Name: "Ram"},
				{ID: 2, Name: "Shyam"},
			},
			mockErr: nil,
			expected: []models.User{
				{ID: 1, Name: "Ram"},
				{ID: 2, Name: "Shyam"},
			},
			expectErr: nil,
		},
		{
			name:      "returns error when store fails",
			mockResp:  []models.User{},
			mockErr:   models.CustomError{Code: http.StatusInternalServerError, Message: "DB error"},
			expected:  []models.User{},
			expectErr: models.CustomError{Code: http.StatusInternalServerError, Message: "DB error"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockStore.EXPECT().GetAllUsers().Return(tc.mockResp, tc.mockErr)

			svc := user.NewUserService(mockStore)

			result, err := svc.GetAllUsers()

			if !errors.Is(err, tc.expectErr) {
				t.Errorf("Expected error %v, got %v", tc.expectErr, err)
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected result %v, got %v", tc.expected, result)
			}
		})
	}
}
