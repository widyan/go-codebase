package helper

import (
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

func SendEmailUsingSMTP(logger CustomLogger, from, pass, to, identity, msg, smtpMail, port string) (err error) {
	err = smtp.SendMail(smtpMail+":"+port,
		smtp.PlainAuth(identity, from, pass, smtpMail),
		from, []string{to}, []byte(msg))
	if err != nil {
		logger.Error(err)
		return
	}
	return
}

func ConvertTzToNormal(timestamp string) (date time.Time, err error) {
	t, err := time.Parse(time.RFC3339, timestamp)
	loc, err := time.LoadLocation("Local")
	if err != nil {
		return
	}
	date = t.In(loc)
	return
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
