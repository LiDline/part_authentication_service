package authservices

import (
	"context"
	"test/config"
	"test/internal/db"
	models "test/internal/types"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v4"
)

func CreateTokens(req models.LoginRequest) (string, error) {
	errAuth := getUserByGUID(req.Guid)

	if errAuth != nil {
		return "", errAuth
	}

	accessToken, accessTokenErr := GenerateAccessToken(req.Guid, req.Ip)

	if accessTokenErr != nil {
		return "", accessTokenErr
	}

	return accessToken, nil
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

func GenerateAccessToken(guid string, ip string) (string, error) {
	payload := jwt.MapClaims{
		"guid": guid,
		"ip":   ip,
		"exp":  time.Now().Add(1 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, payload)

	accessToken, err := token.SignedString([]byte(config.Secret))

	return accessToken, err
}
