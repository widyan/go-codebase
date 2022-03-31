package entity

type Users struct {
	ID        int    `json:"id"`
	Fullname  string `json:"fullname" validate:"required"`
	NoHP      string `json:"no_hp" validate:"required"`
	IsAttend  bool   `json:"is_attend"`
	CreatedAt string `json:"created_at"`
}
