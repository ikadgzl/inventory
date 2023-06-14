package repository

import "fmt"

func (r *authRepository) Create(name, email string, password []byte) error {
	s := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)"
	stmt, err := r.db.Prepare(s)
	if err != nil {
		return fmt.Errorf("failed to prepare statement at Create: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, email, password)
	if err != nil {
		return fmt.Errorf("failed to execute query at Create: %v", err)
	}

	return nil
}
