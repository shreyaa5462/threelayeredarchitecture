package user

import (
	"ThreeLayeredArchitecture/models"
	"ThreeLayeredArchitecture/store/user"
)

type UserService struct {
	Store *user.UserStore
}

func NewUserService(store *user.UserStore) *UserService {
	return &UserService{Store: store}
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	return s.Store.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Store.GetAllUsers()
}

func (s *UserService) GetUserByID(id int) (models.User, error) {
	return s.Store.GetUserByID(id)
}
