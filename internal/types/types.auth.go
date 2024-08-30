package models

type LoginRequest struct {
	Guid string `json:"guid" validate:"required,uuid4"`
	Ip   string `json:"ip" validate:"required,ip"`
}

type LoginResponse struct {
	Access_token  string `json:"access_token" validate:"required"`
	Refresh_token string `json:"refresh_token" validate:"required"`
}
