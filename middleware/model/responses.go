package model

type ResponsesToken struct {
	Token               string `json:"token"`
	RefreshToken        string `json:"refresh_token"`
	ExpiredToken        int64  `json:"expired_token"`
	ExpiredRefreshToken int64  `json:"expired_refresh_token"`
}
