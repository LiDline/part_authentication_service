package models

type LoginRequest struct {
	GUID     string `json:"guid"`
	Password string `json:"password"`
}
