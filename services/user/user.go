package user

import (
	"ThreeLayeredArchitecture/models"
)

type UserService struct {
	Store userStoreInterface
}

func NewUserService(store userStoreInterface) *UserService {
	return &UserService{Store: store}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Store.GetAllUsers()
}

func (s *UserService) GetUserByID(id int) (models.User, error) {
	return s.Store.GetUserByID(id)
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	return s.Store.CreateUser(user)
}
