package domain

import (
	"codebase/go-codebase/entity"
	"codebase/go-codebase/helper"
	"codebase/go-codebase/model"
	"context"
)

type Usecase struct {
	Repository Repository_Interface
	Api        API_Interface
	Worker     Worker_Interface
	Tools      Tools_Interface
	Logger     *helper.CustomLogger
}

func CreateUsecase(repo Repository_Interface, api API_Interface, worker Worker_Interface, tools Tools_Interface, logger *helper.CustomLogger) *Usecase {
	return &Usecase{
		Repository: repo,
		Api:        api,
		Worker:     worker,
		Tools:      tools,
		Logger:     logger,
	}
}

func (b Usecase) VerifikasiToken(ctx context.Context, token string) (codes int, vrf model.VerifikasiToken, err error) {
	return b.Api.VerifikasiToken(ctx, token)
}

func (b Usecase) InsertUser(ctx context.Context, user entity.Users) (err error) {
	return b.Repository.InsertUser(ctx, user)
}

func (b Usecase) SendErrorToTelegram(nameService, message string) {
	b.Api.SendErrorToTelegram(nameService, message)
}

func (b Usecase) GetOneUser(ctx context.Context) (user entity.Users, err error) {

	return b.Repository.GetOneUser(ctx)
}

func (b Usecase) GetAllUsers(ctx context.Context) (users []entity.Users, err error) {
	return b.Repository.GetAllUsers(ctx)
}

func (b Usecase) UpdateUserByID(ctx context.Context, id int, fullname string) (err error) {
	return b.Repository.UpdateUserByID(ctx, id, fullname)
}
