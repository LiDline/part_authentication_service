package models

type LoginRequest struct {
	Guid string `json:"guid" validate:"required,uuid4"`
	Ip   string `json:"ip" validate:"required,ip"`
}

type LoginResponse struct {
	Access_token  string
	Refresh_token string
}
