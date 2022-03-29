package model

type Logger struct {
	Time       string `json:"time"`
	Method     string `json:"method"`
	Path       string `json:"path"`
	StatusCode int    `json:"status_code"`
	Latency    string `json:"latency"`
	Body       string `json:"body"`
	Header     string `json:"header"`
	Error      string `json:"error"`
	UserAgent  string `json:"user_agent"`
	ClientIP   string `json:"client_ip"`
	Responses  string `json:"responses"`
}

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

type CaptureError struct {
	Type       string `json:"type"`
	HttpCode   int    `json:"http_code"`
	ErrorCode  int    `json:"error_code"`
	ResultCode string `json:"resultCode"`
}
