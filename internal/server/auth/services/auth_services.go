package authservices

import (
	"context"
	"test/internal/db"
	models "test/internal/types"

	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateTokens(req models.LoginRequest) (string, error) {
	errBd := getUserByGUID(req)

	if errBd != nil {
		return "", errBd
	}

	return req.GUID, nil
}

func getUserByGUID(req models.LoginRequest) error {
	query := "SELECT id, password FROM users WHERE id = $1"

	row := db.Conn.QueryRow(context.Background(), query, req.GUID)

	var id string
	var hashedPassword string

	err := row.Scan(&id, &hashedPassword)

	if err == pgx.ErrNoRows {
		return err
	}

	errPassword := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))

	if errPassword != nil {
		return errPassword
	}

	return nil
}
