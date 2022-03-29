package logger

import (
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func SendToTelgram(nameService, message string) {
	if strings.Contains(message, "Error 1062: Duplicate entry") {
		return
	}

	_, file, line, _ := runtime.Caller(1)

	token := ""
	chatID := ""
	if os.Getenv("MODE") == "development" {
		token = os.Getenv("TOKEN_BOT_TELEGRAM_TESTING")
		chatID = os.Getenv("CHAT_ID_TESTING")
	} else {
		token = os.Getenv("TOKEN_BOT_TELEGRAM")
		chatID = os.Getenv("CHAT_ID")
	}

	url := "https://api.telegram.org/" + token + "/sendMessage"
	method := "POST"

	location, _ := time.LoadLocation("Local")
	tms := time.Now().In(location)

	text := "service: *" + nameService + "*\n" + "owner: * Scheduller *\n" + "timestamp: *" + tms.Format("2006-01-02 15:04:05") + "*\n" + "line error: *" + file + ": " + strconv.Itoa(line) + "*\n" + "messages: *" + message + "*"

	payload := strings.NewReader("chat_id=" + chatID + "&text=" + text + "&parse_mode=Markdown")

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
}
