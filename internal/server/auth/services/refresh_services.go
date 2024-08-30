package authservices

import (
	"fmt"
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

func UpdateTokens(req models.LoginResponse) (models.LoginResponse, error) {
	verifyToken(req.Access_token)

	// if errToken != nil {
	// 	return models.LoginResponse{}, errToken
	// }
	return models.LoginResponse{}, nil
}

func verifyToken(tokenString string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return config.Secret, nil
	})
	if err != nil {
		log.Fatalf("Error parsing token: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		for key, value := range claims {
			fmt.Printf("%s: %v\n", key, value)
		}
	} else {
		fmt.Println("Invalid token")
	}
}
