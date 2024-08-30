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

type checkRefreshTokenProps struct {
	models.LoginRequest
	Refresh_tokens_id int
}

func UpdateTokens(req models.RefreshRequest) (models.LoginResponse, error) {
	payload := payloadToken(req.Access_token)

	obj, errQueryRefreshToken := checkRefreshToken(payload, req)

	if errQueryRefreshToken != nil {
		return models.LoginResponse{}, errQueryRefreshToken
	}

	res, errGenerated := generateNewTokens(obj)

	if errGenerated != nil {
		return models.LoginResponse{}, errGenerated
	}

	return res, nil
}

func payloadToken(tokenString string) payloadTokenProps {
	token, _ := jwt.ParseWithClaims(tokenString, &tokenPayload{}, func(token *jwt.Token) (interface{}, error) {

		return []byte(config.Secret), nil
	})

	claims, _ := token.Claims.(*tokenPayload)

	return payloadTokenProps{Sub: claims.Sub, Ip: claims.Ip, Iat: claims.Iat}
}

func checkRefreshToken(payload payloadTokenProps, req models.RefreshRequest) (checkRefreshTokenProps, error) {
	query := `
    SELECT 
        refresh_token, 
        ip,
		email,
		refresh_tokens.id AS refresh_tokens_id
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
	var refresh_tokens_id int

	errQuery := row.Scan(&hashRefresh, &ip, &email, &refresh_tokens_id)

	if errQuery == pgx.ErrNoRows {
		return checkRefreshTokenProps{}, errQuery
	}

	errHashRefresh := bcrypt.CompareHashAndPassword([]byte(hashRefresh), []byte(req.Refresh_token))

	if errHashRefresh != nil {
		return checkRefreshTokenProps{}, errHashRefresh
	}

	if payload.Ip != req.Ip {
		sendEmail(email)
	}

	res := checkRefreshTokenProps{
		LoginRequest: models.LoginRequest{
			Guid: payload.Sub,
			Ip:   req.Ip,
		},
		Refresh_tokens_id: refresh_tokens_id,
	}

	return res, nil
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

func generateNewTokens(obj checkRefreshTokenProps) (models.LoginResponse, error) {
	newTokens, errGenerated := generateTokens(obj.LoginRequest)

	if errGenerated != nil {
		return models.LoginResponse{}, errGenerated
	}

	sql := "DELETE FROM refresh_tokens WHERE id = $1"

	_, err := db.Conn.Exec(context.Background(), sql, obj.Refresh_tokens_id)

	if err != nil {
		return models.LoginResponse{}, err
	}

	return newTokens, nil
}
