package user

import (
	"ThreeLayeredArchitecture/models"
	"gofr.dev/pkg/gofr"
)

type UserService struct {
	Store userStoreInterface
}

func NewUserService(store userStoreInterface) *UserService {
	return &UserService{Store: store}
}

func (s *UserService) GetAllUsers(ctx *gofr.Context) ([]models.User, error) {
	return s.Store.GetAllUsers(ctx)
}

func (s *UserService) GetUserByID(ctx *gofr.Context, id int) (models.User, error) {
	return s.Store.GetUserByID(ctx, id)
}

func (s *UserService) CreateUser(ctx *gofr.Context, user models.User) (models.User, error) {
	return s.Store.CreateUser(ctx, user)
}
