package authservices

import (
	"context"
	"test/internal/db"
	models "test/internal/types"

	"github.com/jackc/pgx/v4"
)

func CreateTokens(req models.LoginRequest) (string, error) {
	errAuth := getUserByGUID(req.Guid)

	if errAuth != nil {
		return "", errAuth
	}

	return req.Guid, nil
}

// ----------------------------Check BD----------------------------

func getUserByGUID(guid string) error {
	query := "SELECT id, password FROM users WHERE id = $1"

	row := db.Conn.QueryRow(context.Background(), query, guid)

	var id string

	err := row.Scan(&id)

	if err == pgx.ErrNoRows {
		return err
	}

	return nil
}

// ----------------------------Generate tokens----------------------------
