package model

type SendRequest struct {
	Email string `json:"email" validate:"required,email"`
}
