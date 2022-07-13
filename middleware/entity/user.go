package entity

import "database/sql"

type User struct {
	ID        string       `json:"id"`
	Email     string       `json:"email"`
	Name      string       `json:"name"`
	Role      string       `json:"role"`
	IsActive  bool         `json:"is_active"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}
