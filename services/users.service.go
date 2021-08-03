package services

import (
	"regexp"
	"github.com/HatsuneMikuLab/hrbrain-challenge/models"
	"database/sql"
)

type IUsersService interface {
	GetUserByID(id string) (*models.User, error)
	AddUser(data *models.User) ([]string, bool, error)
}

type usersService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) *usersService {
	return &usersService{ DB: db }
}

func (us *usersService) GetUserByID(id string) (*models.User, error) {
	user := &models.User{}
	row := us.DB.QueryRow("SELECT * FROM users WHERE id = $1", id)
	err := row.Scan(&user.ID, &user.Email)
	if err != nil {
		isNoRowsError, err := regexp.MatchString("sql.*", err.Error())
		if err == nil && isNoRowsError {
			return nil, nil
		}
		return nil, err
	}
	return user, err
}

func (us *usersService) AddUser(data *models.User) ([]string, bool, error) {
	validationErrors := data.Validate()
	if len(validationErrors) > 0 {
		return validationErrors, false, nil
	}
	_, err := us.DB.Exec("INSERT INTO users (id, email) VALUES ($1, $2)", data.ID, data.Email)
	if err != nil {
		isDuplicateKeyError, err := regexp.MatchString("pq.*", err.Error())
		if err == nil && isDuplicateKeyError {
			return nil, true, nil
		}
		return nil, false, err
	}
	return nil, false, err
}