package user

import (
	"ThreeLayeredArchitecture/models"
	"gofr.dev/pkg/gofr"
	"strconv"
)

type UserHandler struct {
	Service UserServiceInterface
}

func NewUserHandler(service UserServiceInterface) *UserHandler {
	return &UserHandler{Service: service}
}

func (u *UserHandler) CreateUser(ctx *gofr.Context) (any, error) {
	var input models.User
	err := ctx.Bind(&input)
	if err != nil {
		return nil, err
	}
	users, err := u.Service.CreateUser(input)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (u *UserHandler) GetAllUsers(ctx *gofr.Context) (any, error) {
	users, err := u.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (u *UserHandler) GetUserByID(ctx *gofr.Context) (any, error) {
	id, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		return nil, err
	}
	users, err := u.Service.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return users, nil
}
