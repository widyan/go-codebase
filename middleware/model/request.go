package model

type RequestToken struct {
	Email string `json:"email" validate:"required"`
}

type RequestUser struct {
	Email    string `json:"email" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Role     string `json:"role" validate:"required"`
	IsActive bool   `json:"is_active" validate:"required"`
}
