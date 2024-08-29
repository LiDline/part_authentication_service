package models

type LoginRequest struct {
	Guid string `json:"guid" validate:"required,uuid4"`
	Ip   string `json:"ip" validate:"required,ip"`
}
