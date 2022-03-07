package api

import (
	"net/http"
	"os"
	"strings"
)

func (a *Api) SendErrorToTelegram(nameService, message string) {
	token := ""
	chatID := ""
	if os.Getenv("MODE") == "development" {
		token = os.Getenv("TOKEN_BOT_TELEGRAM_TESTING")
		chatID = os.Getenv("CHAT_ID_TESTING")
	} else {
		token = os.Getenv("TOKEN_BOT_TELEGRAM")
		chatID = os.Getenv("CHAT_ID")
	}

	owner := os.Getenv("PROJECT_NAME")

	url := "https://api.telegram.org/" + token + "/sendMessage"
	method := "POST"

	text := "service: *" + nameService + "*\n" + "owner: *" + owner + "*\n" + "messages: *" + message + "*"

	payload := strings.NewReader("chat_id=" + chatID + "&text=" + text + "&parse_mode=Markdown")
	if strings.Contains(message, "*") {
		payload = strings.NewReader("chat_id=" + chatID + "&text=" + text)
	}

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, payload)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		// http.Error(w, "failed to query backend", 500)
		return
	} else {
		defer res.Body.Close()
	}
}
