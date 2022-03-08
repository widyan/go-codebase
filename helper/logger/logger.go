package logger

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"codebase/go-codebase/helper"
	"codebase/go-codebase/model"
	"strconv"
	"strings"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
)

func SendLoggerToElasticSearch() {
	var param model.Logger

	if helper.Buf.String() != "" {

		param.StatusCode = 400

		logs := logrus.New()
		// logs.SetOutput(ioutil.Discard)
		logs.Out = ioutil.Discard
		client, err := elastic.NewClient(elastic.SetURL(os.Getenv("URL_ELASTIC_SEARCH")))
		if err != nil {
			// logs.Panic(err)
			logs.Println(err.Error())
			return
		}
		hook, err := elogrus.NewAsyncElasticHook(client, "localhost", logs.Level, os.Getenv("INDEX_LOGGER"))
		if err != nil {
			// logs.Panic(err)
			logs.Println(err.Error())
			return
		}
		logs.Hooks.Add(hook)

		logs.SetFormatter(&logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime: "@timestamp",
				logrus.FieldKeyMsg:  "message",
			},
		})
		logs.SetLevel(logs.Level)

		wkt := time.Now().Format("2006-01-02 15:04:05")
		fields := logrus.Fields{
			"service":     "Scheduller API Gateway",
			"time":        wkt,
			"client_ip":   param.ClientIP,
			"method":      param.Method,
			"path":        param.Path,
			"status_code": param.StatusCode,
			"latency":     param.Latency,
			"body":        param.Body,
			"header":      param.Header,
			"user_agent":  param.UserAgent,
			"responses":   param.Responses,
		}

		if helper.Buf.String() != "" {
			logs.WithFields(fields).Error(helper.Buf.String())
			go SendErrorToTelegram(helper.Buf.String(), "Scheduller API Gateway", wkt, fields)
		}

	}

	helper.Buf = bytes.Buffer{}
}

func SendErrorToTelegram(message, nameService, wkt string, lg logrus.Fields) {
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

	text := "service: *" + nameService + "*\n" + "timestamp: *" + wkt + "* \n" + "path: *" + lg["path"].(string) + "*\n" + "messages: *" + message + "*"

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
