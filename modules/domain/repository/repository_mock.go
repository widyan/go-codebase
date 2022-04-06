package repository

import (
	"codebase/go-codebase/modules/domain/entity"
	"codebase/go-codebase/modules/domain/interfaces"
	"context"

	"github.com/stretchr/testify/mock"
)

type MockRepositoryUserImpl struct {
	mock.Mock
}

func CreateMockRepository() interfaces.Repository_Interface {
	return &MockRepositoryUserImpl{}
}

func (r *MockRepositoryUserImpl) GetAllUsers(ctx context.Context) (user []entity.Users, err error) {
	args := r.Called(ctx)
	return args.Get(0).([]entity.Users), args.Error(1)
}

func (r *MockRepositoryUserImpl) GetOneUser(ctx context.Context) (user entity.Users, err error) {
	args := r.Called(ctx)
	return args.Get(0).(entity.Users), args.Error(1)
}

func (r *MockRepositoryUserImpl) InsertUser(ctx context.Context, user entity.Users) (err error) {
	args := r.Called(ctx, user)
	return args.Error(0)
}

func (r *MockRepositoryUserImpl) UpdateUserByID(ctx context.Context, id int, fullname string) (err error) {
	args := r.Called(ctx, id, fullname)
	return args.Error(0)
}

func (r *MockRepositoryUserImpl) GetOneUserByID(ctx context.Context, id int) (user entity.Users, err error) {
	args := r.Called(ctx, id)
	return args.Get(0).(entity.Users), args.Error(1)
}
