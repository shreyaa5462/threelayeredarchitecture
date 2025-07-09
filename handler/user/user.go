package user

import (
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http/response"
	"strconv"

	"ThreeLayeredArchitecture/models/user"
)

type UserHandler struct {
	Service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{Service: service}
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  Add a new user to the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user body models.User true "User object"
// @Success      201 {object} models.User
// @Failure      400 {string} string "Invalid input"
// @Failure      500 {string} string "Internal Server Error"
// @Router       /users [post]
func (h *UserHandler) CreateUser(ctx *gofr.Context) (any, error) {
	var user models.User

	err := ctx.Bind(&user)
	if err != nil {
		return nil, err
	}

	newUser, err := h.Service.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return response.Raw{Data: newUser}, nil
}

// GetAllUsers godoc
// @Summary      Get all users
// @Description  Retrieve all users from the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200 {array} models.User
// @Failure      500 {string} string "Internal Server Error"
// @Router       /users [get]
func (h *UserHandler) GetAllUsers(ctx *gofr.Context) (any, error) {
	users, err := h.Service.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return response.Raw{Data: users}, nil
}

// GetUserByID godoc
// @Summary      Get user by ID
// @Description  Retrieve a specific user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id path int true "User ID"
// @Success      200 {object} models.User
// @Failure      400 {string} string "Invalid ID"
// @Failure      404 {string} string "User not found"
// @Router       /users/{id} [get]
func (h *UserHandler) GetUserbyID(ctx *gofr.Context) (any, error) {
	//idStr := r.PathValue("id")
	id, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		return nil, err
	}

	user, err := h.Service.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return response.Raw{Data: user}, nil
}
