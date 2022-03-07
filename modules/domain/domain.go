package domain

import (
	"codebase/go-codebase/entity"
	"codebase/go-codebase/model"
	"context"
)

type Usecase_Interface interface {
	VerifikasiToken(ctx context.Context, token string) (codes int, vrf model.VerifikasiToken, err error)
	InsertUser(ctx context.Context, user entity.Users) (err error)
	SendErrorToTelegram(nameService, message string)
	GetOneUser(ctx context.Context) (user entity.Users, err error)
	GetAllUsers(ctx context.Context) (users []entity.Users, err error)
	UpdateUserByID(ctx context.Context, id int, fullname string) (err error)
}

type ToolsUsecase_Interface interface {
	SendEmails(from, pass, to, identity, msg, smtpMail, port string) (err error)
}

type Repository_Interface interface {
	InsertUser(ctx context.Context, user entity.Users) (err error)
	GetOneUser(ctx context.Context) (user entity.Users, err error)
	GetAllUsers(ctx context.Context) (users []entity.Users, err error)
	UpdateUserByID(ctx context.Context, id int, fullname string) (err error)
}

type API_Interface interface {
	VerifikasiToken(ctx context.Context, token string) (codes int, vrf model.VerifikasiToken, err error)
	SendErrorToTelegram(nameService, message string)
}

type Worker_Interface interface {
}

type Tools_Interface interface {
}
