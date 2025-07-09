package user

import (
	"ThreeLayeredArchitecture/models"
	"gofr.dev/pkg/gofr"
)

type userStoreInterface interface {
	GetAllUsers(ctx *gofr.Context) ([]models.User, error)
	GetUserByID(ctx *gofr.Context, id int) (models.User, error)
	CreateUser(ctx *gofr.Context, user models.User) (models.User, error)
}
