package authservices

import (
	"context"
	"test/internal/db"
	models "test/internal/types"

	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateTokens(req models.LoginRequest) (string, error) {
	errAuth := checkPassword(req)

	if errAuth != nil {
		return "", errAuth
	}

	return req.GUID, nil
}

func getUserByGUID(guid string) (string, error) {
	query := "SELECT id, password FROM users WHERE id = $1"

	row := db.Conn.QueryRow(context.Background(), query, guid)

	var id string
	var hashedPassword string

	err := row.Scan(&id, &hashedPassword)

	if err == pgx.ErrNoRows {
		return "", err
	}

	return hashedPassword, nil
}

func checkPassword(req models.LoginRequest) error {
	hashedPassword, errBd := getUserByGUID(req.GUID)

	if errBd != nil {
		return errBd
	}

	errPassword := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))

	if errPassword != nil {
		return errPassword
	}

	return nil
}
