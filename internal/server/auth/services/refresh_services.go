package authservices

import (
	"log"
	"test/config"
	models "test/internal/types"

	"github.com/golang-jwt/jwt/v5"
)

type tokenPayload struct {
	sub string
	ip  string
	exp string
	iat string
}

func UpdateTokens(req models.RefreshRequest) (models.LoginResponse, error) {
	verifyToken(req.Access_token)

	// if errToken != nil {
	// 	return models.LoginResponse{}, errToken
	// }
	return models.LoginResponse{}, nil
}

func verifyToken(tokenString string) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		return []byte(config.Secret), nil
	})

	log.Print(token.Valid)

}
