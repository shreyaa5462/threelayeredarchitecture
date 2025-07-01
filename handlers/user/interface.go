package user

import "ThreeLayeredArchitecture/models"

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
}
