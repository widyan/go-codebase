package usecase

import (
	"codebase/go-codebase/modules/domain/entity"
	"context"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type TestUsecase struct {
	Mock   *MockRepositoryUserImpl
	Logger *logrus.Logger
	Ctx    context.Context
}

func CreateInitiationTest() *TestUsecase {
	var logger = logrus.New()
	logger.SetReportCaller(true)
	return &TestUsecase{Mock: new(MockRepositoryUserImpl), Logger: logger, Ctx: context.Background()}
}

func TestInsertUser(t *testing.T) {
	u := CreateInitiationTest()

	var tms time.Time = time.Now()
	entityUser := entity.Users{
		ID:        1,
		Fullname:  "Test User",
		NoHP:      "08123456789",
		CreatedAt: tms.Format("2006-01-02 15:04:05"),
	}

	u.Mock.On("InsertUser", u.Ctx, entityUser).Return(nil)
	usecase := CreateUsecase(u.Mock, u.Logger)
	err := usecase.InsertUser(u.Ctx, entityUser)
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
	usecase := CreateUsecase(u.Mock, u.Logger)
	userall, err := usecase.GetAllUsers(u.Ctx)
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
	usecase := CreateUsecase(u.Mock, u.Logger)
	user, err := usecase.GetOneUser(u.Ctx)
	assert.NoError(t, err)
	assert.Equal(t, user, entityUser)
}

func TestUpdateUserByID(t *testing.T) {
	u := CreateInitiationTest()

	u.Mock.On("UpdateUserByID", u.Ctx, 1, "Test").Return(nil)
	usecase := CreateUsecase(u.Mock, u.Logger)
	err := usecase.UpdateUserByID(u.Ctx, 1, "Test")
	assert.NoError(t, err)
}
