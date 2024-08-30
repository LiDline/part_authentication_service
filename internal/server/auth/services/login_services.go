package authservices

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"log"
	"test/config"
	"test/internal/constants"
	"test/internal/db"
	models "test/internal/types"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateTokens(req models.LoginRequest) (models.LoginResponse, error) {
	errAuth := getUserByGUID(req.Guid)

	if errAuth != nil {
		return models.LoginResponse{}, errAuth
	}

	accessToken, timeGenerateToken, accessTokenErr := GenerateAccessToken(req)

	if accessTokenErr != nil {
		return models.LoginResponse{}, accessTokenErr
	}

	refreshToken, refreshTokenErr := GenerateRefreshToken(req, timeGenerateToken)

	if refreshTokenErr != nil {
		return models.LoginResponse{}, refreshTokenErr
	}

	return models.LoginResponse{Access_token: accessToken, Refresh_token: refreshToken}, nil
}

// ----------------------------Check BD----------------------------

func getUserByGUID(guid string) error {
	query := "SELECT id FROM users WHERE id = $1"

	row := db.Conn.QueryRow(context.Background(), query, guid)

	var id string

	err := row.Scan(&id)

	if err == pgx.ErrNoRows {
		return err
	}

	return nil
}

// ----------------------------Generate tokens----------------------------

func GenerateAccessToken(req models.LoginRequest) (string, int, error) {
	timeNow := time.Now().UTC().Nanosecond()

	log.Print(timeNow)

	log.Print(time.Now().Add(constants.EXP_ACCESS_TOKEN * 3600000).UTC().Nanosecond())

	payload := jwt.MapClaims{
		"sub": req.Guid,
		"ip":  req.Ip,
		"exp": time.Now().Add(constants.EXP_ACCESS_TOKEN * 3600000).UTC().Nanosecond(),
		"iat": timeNow,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, payload)

	accessToken, err := token.SignedString([]byte(config.Secret))

	return accessToken, timeNow, err
}

func GenerateRefreshToken(req models.LoginRequest, timeGenerateToken int) (string, error) {
	bytes := make([]byte, 16)

	rand.Read(bytes)

	refresh := base64.StdEncoding.EncodeToString(bytes)

	refreshBcrypt, _ := bcrypt.GenerateFromPassword([]byte(refresh), bcrypt.DefaultCost)

	sqlString := "INSERT INTO refresh_tokens (refresh_token, created_at, ip, id) VALUES ($1, $2, $3, $4)"

	_, err := db.Conn.Exec(context.Background(), sqlString, string(refreshBcrypt), timeGenerateToken, req.Ip, req.Guid)

	if err != nil {
		return "", err
	}

	return refresh, nil
}
