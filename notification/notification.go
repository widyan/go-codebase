package notification

import (
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/widyan/go-codebase/helper"
)

type Notification struct {
	API      helper.API_Interface
	Service  string
	TokenBot string
	ChatID   string
}

func CreateNotification(api helper.API_Interface, service, tokenBot, chatID string) *Notification {
	return &Notification{
		API:      api,
		Service:  service,
		TokenBot: tokenBot,
		ChatID:   chatID,
	}
}

func (s *Notification) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.PanicLevel, logrus.FatalLevel, logrus.WarnLevel}
}

func (s *Notification) Fire(entry *logrus.Entry) error {
	location, _ := time.LoadLocation("Local")
	tms := time.Now().In(location)
	var isContainArstik bool = false
	if strings.Contains(entry.Message, "*") {
		isContainArstik = true
	}
	text := "service: *" + s.Service + "*\n" + "timestamp: *" + tms.Format("2006-01-02 15:04:05") + "*\n" + "line error: *" + entry.Caller.File + ":" + strconv.Itoa(entry.Caller.Line) + "*\n" + "messages: *" + entry.Message + "*"
	go s.API.SendToTelegram("https://api.telegram.org/", "POST", s.TokenBot, s.ChatID, text, isContainArstik)
	return nil
}
