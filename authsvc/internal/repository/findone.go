package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/ikadgzl/inventory/authsvc/internal/model"
	"github.com/ikadgzl/inventory/authsvc/internal/util"
)

func (r *authRepository) FindOne(email string) (*model.AuthUser, error) {
	q := `SELECT id, password FROM users WHERE email = $1`
	stmt, err := r.db.Prepare(q)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement at FindOne: %v", err)
	}
	defer stmt.Close()

	au := &model.AuthUser{}
	err = stmt.QueryRow(email).Scan(&au.ID, &au.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, util.ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to execute query at FindOne: %v", err)
	}

	return au, nil
}
