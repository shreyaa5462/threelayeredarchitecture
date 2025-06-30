package userstore

import (
	"ThreeLayeredArchitecture/models"
	"database/sql"
	"log"
)

type UserStore struct {
	DB *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{DB: db}
}

func (s *UserStore) GetAllUsers() ([]models.User, error) {
	query := "SELECT id, name FROM users"
	rows, err := s.DB.Query(query)
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

func (s *UserStore) GetUserByID(id int) (models.User, error) {
	query := "SELECT id, name FROM users WHERE id = ?"
	row := s.DB.QueryRow(query, id)
	var u models.User
	err := row.Scan(&u.ID, &u.Name)
	return u, err
}

func (s *UserStore) CreateUser(user models.User) (models.User, error) {
	query := "INSERT INTO users (name) VALUES (?)"
	result, err := s.DB.Exec(query, user.Name)
	if err != nil {
		return models.User{}, err
	}
	id, _ := result.LastInsertId()
	user.ID = int(id)
	return user, nil
}
