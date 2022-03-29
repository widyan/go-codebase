package middleware

type EntityVerifikasiToken struct {
	Code      int                       `json:"code"`
	Status    string                    `json:"status"`
	Message   string                    `json:"message"`
	ErrorCode int                       `json:"error_code"`
	Data      EntityDataVerifikasiToken `json:"data"`
}

type EntityDataVerifikasiToken struct {
	DeviceID       string `json:"deviceId"`
	Email          string `json:"email"`
	Exp            int    `json:"exp"`
	Iat            int    `json:"iat"`
	IsIndiboxUser  bool   `json:"is_indibox_user"`
	IsIndihomeUser bool   `json:"is_indihome_user"`
	IsUseeTvUser   bool   `json:"is_useeTv_user"`
	Iss            string `json:"iss"`
	LoginStatus    bool   `json:"loginStatus"`
	SubscriberID   string `json:"subscriber_id"`
	UserID         string `json:"userId"`
	UserActive     int    `json:"user_active"`
	Fullname       string `json:"fullname"`
	IsSuperuser    bool   `json:"is_superuser"`
}
