package user

import (
	"ThreeLayeredArchitecture/models"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http/response"
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
	users, err := u.Service.CreateUser(ctx, input)
	if err != nil {
		return nil, err
	}
	return response.Raw{Data: users}, nil
}
func (u *UserHandler) GetAllUsers(ctx *gofr.Context) (any, error) {
	users, err := u.Service.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	return response.Raw{Data: users}, nil
}
func (u *UserHandler) GetUserByID(ctx *gofr.Context) (any, error) {
	id, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		return nil, err
	}
	users, err := u.Service.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return response.Raw{Data: users}, nil
}
