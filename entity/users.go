package entity

type Users struct {
	ID        int    `json:"id"`
	Fullname  string `json:"fullname"`
	NoHP      string `json:"no_hp"`
	IsAttend  bool   `json:"is_attend"`
	CreatedAt string `json:"created_at"`
}
