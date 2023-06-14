package repository

import (
	"database/sql"

	"github.com/ikadgzl/inventory/authsvc/internal/model"
)

type AuthRepository interface {
	FindOne(email string) (*model.AuthUser, error)
	Create(name, email string, password []byte) error
}

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}
