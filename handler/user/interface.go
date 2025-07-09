package user

import (
	models "ThreeLayeredArchitecture/models/user"
	"gofr.dev/pkg/gofr"
)

type UserService interface {
	CreateUser(*gofr.Context, models.User) (models.User, error)
	GetAllUsers(*gofr.Context) ([]models.User, error)
	GetUserByID(*gofr.Context, int) (models.User, error)
}
