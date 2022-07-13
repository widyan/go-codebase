package interfaces

import (
	"context"

	"github.com/widyan/go-codebase/modules/domain/entity"
)

type Usecase_Interface interface {
	InsertUser(ctx context.Context, user entity.Users) (err error)
	GetOneUser(ctx context.Context) (user entity.Users, err error)
	GetAllUsers(ctx context.Context) (users []entity.Users, err error)
	UpdateUserByID(ctx context.Context, id int, fullname string) (err error)
	GetOneUserByID(ctx context.Context, id int) (user entity.Users, err error)
}

type Repository_Interface interface {
	InsertUser(ctx context.Context, user entity.Users) (err error)
	GetOneUser(ctx context.Context) (user entity.Users, err error)
	GetAllUsers(ctx context.Context) (users []entity.Users, err error)
	UpdateUserByID(ctx context.Context, id int, fullname string) (err error)
	GetOneUserByID(ctx context.Context, id int) (user entity.Users, err error)
}

type Scheduller interface {
	TestScheduller()
}
