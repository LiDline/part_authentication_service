package models

type LoginRequest struct {
	Guid     string `json:"guid" validate:"required,uuid4"`
	Password string `json:"password" validate:"required"`
	Ip       string `json:"ip" validate:"required,ip"`
}
