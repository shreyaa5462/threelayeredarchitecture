package user

import (
	"ThreeLayeredArchitecture/models"
	"gofr.dev/pkg/gofr"
	"log"
)

type UserStore struct {
}

func NewUserStore() *UserStore {
	return &UserStore{}
}

func (s *UserStore) GetAllUsers(ctx *gofr.Context) ([]models.User, error) {
	query := "SELECT id, name FROM users"
	rows, err := ctx.SQL.Query(query)

	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	var users []models.User

	for rows.Next() {
		var u models.User
		err = rows.Scan(&u.ID, &u.Name)

		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

func (s *UserStore) GetUserByID(ctx *gofr.Context, id int) (models.User, error) {
	query := "SELECT id, name FROM users WHERE id = ?"
	row := ctx.SQL.QueryRow(query, id)

	var u models.User
	err := row.Scan(&u.ID, &u.Name)

	return u, err
}

func (s *UserStore) CreateUser(ctx *gofr.Context, user models.User) (models.User, error) {
	query := "INSERT INTO users (name) VALUES (?)"
	result, err := ctx.SQL.Exec(query, user.Name)

	if err != nil {
		return models.User{}, err
	}

	id, _ := result.LastInsertId()
	user.ID = int(id)

	return user, nil
}
