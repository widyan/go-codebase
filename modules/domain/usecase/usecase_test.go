package usecase

import (
	"context"
	"github.com/widyan/go-codebase/modules/domain/entity"
	"github.com/widyan/go-codebase/modules/domain/repository"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TestUsecase struct {
	Mock    *repository.MockRepositoryUserImpl
	Service Usecase
	Logger  *logrus.Logger
	Ctx     context.Context
}

func CreateInitiationTest() *TestUsecase {
	var logger = logrus.New()
	repositoryUser := &repository.MockRepositoryUserImpl{Mock: mock.Mock{}}
	serviceUser := Usecase{repositoryUser, logger}
	return &TestUsecase{Mock: repositoryUser, Service: serviceUser, Logger: logger, Ctx: context.Background()}
}

func TestInsertUser(t *testing.T) {
	u := CreateInitiationTest()
	CreateUsecase(u.Mock, u.Logger)

	var tms time.Time = time.Now()
	entityUser := entity.Users{
		ID:        1,
		Fullname:  "Test User",
		NoHP:      "08123456789",
		CreatedAt: tms.Format("2006-01-02 15:04:05"),
	}

	u.Mock.On("InsertUser", u.Ctx, entityUser).Return(nil)
	err := u.Service.InsertUser(u.Ctx, entityUser)
	assert.NoError(t, err)
}

func TestGetAllUsers(t *testing.T) {
	u := CreateInitiationTest()
	var tms time.Time = time.Now()

	users := []entity.Users{
		{
			ID:        1,
			Fullname:  "Test User",
			NoHP:      "08123456789",
			CreatedAt: tms.Format("2006-01-02 15:04:05"),
		}, {
			ID:        2,
			Fullname:  "Test User",
			NoHP:      "08123456789",
			CreatedAt: tms.Format("2006-01-02 15:04:05"),
		},
	}

	u.Mock.On("GetAllUsers", u.Ctx).Return(users, nil)
	userall, err := u.Service.GetAllUsers(u.Ctx)
	assert.NoError(t, err)

	assert.Equal(t, userall, users)
}

func TestGetOneUser(t *testing.T) {
	u := CreateInitiationTest()

	var tms time.Time = time.Now()
	entityUser := entity.Users{
		ID:        1,
		Fullname:  "Test User",
		NoHP:      "08123456789",
		CreatedAt: tms.Format("2006-01-02 15:04:05"),
	}

	u.Mock.On("GetOneUser", u.Ctx).Return(entityUser, nil)
	user, err := u.Service.GetOneUser(u.Ctx)
	assert.NoError(t, err)
	assert.Equal(t, user, entityUser)
}

func TestUpdateUserByID(t *testing.T) {
	u := CreateInitiationTest()

	u.Mock.On("UpdateUserByID", u.Ctx, 1, "Test").Return(nil)
	err := u.Service.UpdateUserByID(u.Ctx, 1, "Test")
	assert.NoError(t, err)
}
