package helper

import (
	"codebase/go-codebase/model"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/smtp"
	"time"

	"go.elastic.co/apm/module/apmhttp"
)

var tracingClient = apmhttp.WrapClient(http.DefaultClient)

// RandomWords is
func RandomWords(sumRandom int) string {
	const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, sumRandom)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func SendEmailUsingSMTP(from, pass, to, identity, msg, smtpMail, port string) (err error) {
	err = smtp.SendMail(smtpMail+":"+port,
		smtp.PlainAuth(identity, from, pass, smtpMail),
		from, []string{to}, []byte(msg))
	if err != nil {
		return
	}
	return
}

func ConvertTzToNormal(timestamp string) (date time.Time, err error) {
	t, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	date = t.In(loc)
	return
}

func SetCaptureError(capt model.CaptureError) error {
	capt.Type = "capture error"
	byteCapt, _ := json.Marshal(capt)
	return fmt.Errorf(string(byteCapt))
}

func UniqueString(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func UniqueInt(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// func SendEmailUsingSMTP() (err error) {
// 	from := "xxx@gmail.com"
// 	pass := "xxxxx"
// 	to := "xxx@gmail.com"

// 	msg := "From: " + from + "\n" +
// 		"To: " + to + "\n" +
// 		"Subject: Hello there\n\n"

// 	err = smtp.SendMail("smtp.gmail.com:587",
// 		smtp.PlainAuth("xxxxx", from, pass, "smtp.gmail.com"),
// 		from, []string{to}, []byte(msg))

// 	if err != nil {
// 		log.Printf("smtp error: %s", err)
// 		return
// 	}
// }
