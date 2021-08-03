package services

import (
	"github.com/HatsuneMikuLab/hrbrain-challenge/models"
	"database/sql"
	"context"
)

type IUsersService interface {
	GetByID(ctx context.Context, id string) (*models.User, error)
	Add(ctx context.Context, data *models.User) ([]string error)
}

type usersService struct {
	DB *sql.DB
	TableName string
}

func NewUserService(db *sql.DB, tableName string) *usersService {
	return &usersService{ DB: db, TableName: tableName }
}

func (us *usersService) GetByID(ctx context.Context, id string) (*models.User, error) {
	user := &models.User{}
	row := us.DB.QueryRowContext(ctx, "SELECT * FROM ? WHERE user = ?", us.TableName, id)
	err := row.Scan(user)
	return user, err
}

func (us *usersService) Add(ctx context.Context, data *models.User) ([]string, error) {
	validationErrors := data.Validate()
	if len(validationErrors) != 0 {
		return validationErrors, nil
	}
	_, err := us.DB.ExecContext(ctx, "INSERT INTO ? (user, email) VALUES (?, ?)", data.User, data.Email)
	return nil, err
}