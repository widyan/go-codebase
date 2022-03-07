package tools

import (
	"codebase/go-codebase/helper"
	"net/smtp"
)

type Tools struct {
	Logger *helper.CustomLogger
}

func CreateTools(logger *helper.CustomLogger) *Tools {
	return &Tools{logger}
}

func (a *Tools) SendEmails(from, pass, to, identity, msg, smtpMail, port string) (err error) {
	err = smtp.SendMail(smtpMail+":"+port,
		smtp.PlainAuth(identity, from, pass, smtpMail),
		from, []string{to}, []byte(msg))
	if err != nil {
		a.Logger.Error(err.Error())
		return
	}
	return
}
