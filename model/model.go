package model

import "codebase/go-codebase/entity"

// Header is
type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Responses is
type Responses struct {
	Code      int         `json:"code"`
	Status    string      `json:"status"`
	Message   string      `json:"message"`
	ErrorCode int         `json:"error_code"`
	Data      interface{} `json:"data"`
}

type FormData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// VerifikasiToken is
type VerifikasiToken struct {
	Code      int                 `json:"code"`
	Status    string              `json:"status"`
	Message   string              `json:"message"`
	ErrorCode int                 `json:"error_code"`
	Data      DataVerifikasiToken `json:"data"`
}

type DataVerifikasiToken struct {
	DeviceID       string `json:"deviceId"`
	Email          string `json:"email"`
	Exp            int    `json:"exp"`
	Iat            int    `json:"iat"`
	Iss            string `json:"iss"`
	LoginStatus    bool   `json:"loginStatus"`
	SubscriberID   string `json:"subscriber_id"`
	UserID         string `json:"userId"`
	UserActive     int    `json:"user_active"`
	Fullname       string `json:"fullname"`
}

type ResponsesRedisVerfikasiToken struct {
	Verifytoken struct {
		Deviceid       string `json:"deviceId"`
		DeviceName     string `json:"device_name"`
		DeviceType     string `json:"device_type"`
		Email          string `json:"email"`
		Exp            int    `json:"exp"`
		FcmToken       string `json:"fcm_token"`
		Fullname       string `json:"fullname"`
		Iat            int    `json:"iat"`
		Iss            string `json:"iss"`
		Keygen         string `json:"keygen"`
		KeygenToken    string `json:"keygen_token"`
		Loginstatus    bool   `json:"loginStatus"`
		RefreshToken   string `json:"refresh_token"`
		SubscriberID   string `json:"subscriber_id"`
		Userid         string `json:"userId"`
		UserActive     int    `json:"user_active"`
	} `json:"VerifyToken"`
	ID int `json:"id"`
}

type ModifyUsers struct {
	entity.Users
	IsActive bool
}
