package usecase

import (
	"context"
	"github.com/widyan/go-codebase/modules/domain/entity"
	"github.com/widyan/go-codebase/modules/domain/interfaces"

	"github.com/sirupsen/logrus"
)

type Usecase struct {
	Repository interfaces.Repository_Interface
	Logger     *logrus.Logger
}

func CreateUsecase(repo interfaces.Repository_Interface, logger *logrus.Logger) *Usecase {
	return &Usecase{
		Repository: repo,
		Logger:     logger,
	}
}

func (b *Usecase) InsertUser(ctx context.Context, user entity.Users) (err error) {
	return b.Repository.InsertUser(ctx, user)
}

func (b *Usecase) GetOneUser(ctx context.Context) (user entity.Users, err error) {
	return b.Repository.GetOneUser(ctx)
}

func (b *Usecase) GetAllUsers(ctx context.Context) (users []entity.Users, err error) {
	return b.Repository.GetAllUsers(ctx)
}

func (b *Usecase) UpdateUserByID(ctx context.Context, id int, fullname string) (err error) {
	return b.Repository.UpdateUserByID(ctx, id, fullname)
}

func (b *Usecase) GetOneUserByID(ctx context.Context, id int) (user entity.Users, err error) {
	return b.Repository.GetOneUserByID(ctx, id)
}
