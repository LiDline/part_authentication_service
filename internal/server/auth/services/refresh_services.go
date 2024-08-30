package authservices

import (
	"context"
	"log"
	"net/smtp"
	"test/config"
	"test/internal/db"
	models "test/internal/types"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

type tokenPayload struct {
	jwt.RegisteredClaims
	Sub string `json:"sub"`
	Ip  string `json:"ip"`
	Exp int64  `json:"exp"`
	Iat int64  `json:"iat"`
}

type payloadTokenProps struct {
	Sub string `json:"sub"`
	Ip  string `json:"ip"`
	Iat int64  `json:"iat"`
}

func UpdateTokens(req models.RefreshRequest) (models.LoginResponse, error) {
	payload := payloadToken(req.Access_token)

	errQueryRefreshToken := checkRefreshToken(payload, req)

	if errQueryRefreshToken != nil {
		return models.LoginResponse{}, errQueryRefreshToken
	}

	return models.LoginResponse{}, nil
}

func payloadToken(tokenString string) payloadTokenProps {
	token, _ := jwt.ParseWithClaims(tokenString, &tokenPayload{}, func(token *jwt.Token) (interface{}, error) {

		return []byte(config.Secret), nil
	})

	claims, _ := token.Claims.(*tokenPayload)

	return payloadTokenProps{Sub: claims.Sub, Ip: claims.Ip, Iat: claims.Iat}
}

func checkRefreshToken(payload payloadTokenProps, req models.RefreshRequest) error {
	query := `
    SELECT 
        refresh_token, 
        ip,
		email 
    FROM 
        refresh_tokens 
	INNER JOIN users ON users.id = refresh_tokens.user_id
    WHERE 
        user_id = $1 
        AND created_at = $2
`

	row := db.Conn.QueryRow(context.Background(), query, payload.Sub, payload.Iat)

	var hashRefresh string
	var ip string
	var email string

	errQuery := row.Scan(&hashRefresh, &ip, &email)

	if errQuery == pgx.ErrNoRows {
		return errQuery
	}

	errHashRefresh := bcrypt.CompareHashAndPassword([]byte(hashRefresh), []byte(req.Refresh_token))

	if errHashRefresh != nil {
		return errHashRefresh
	}

	if payload.Ip != req.Ip {
		sendEmail(email)
	}

	return nil
}

func sendEmail(toEmailAddress string) {
	from := config.Email
	password := config.PasswordEmail

	to := []string{toEmailAddress}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	subject := "Subject: This is the subject of the mail\n"
	body := "This is the body of the mail"
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)

	if err != nil {
		log.Print(err)
	}
}
