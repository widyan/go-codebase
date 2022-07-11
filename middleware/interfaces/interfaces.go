package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/widyan/go-codebase/middleware/entity"
	"github.com/widyan/go-codebase/middleware/model"
)

type UsecaseInterface interface {
	CreateTokenServices(ctx context.Context, request model.RequestToken) (responses model.ResponsesToken, err error)
	AddUser(ctx context.Context, user model.RequestUser) (err error)
	VerifyAutorizationToken() gin.HandlerFunc
}

type RepositoryInterface interface {
	AddUser(ctx context.Context, user entity.User) (err error)
	GetUserBasedOnEmail(ctx context.Context, email string) (users []entity.User, err error)
}

type ToolsInterface interface {
	GetUUID() (uid string)
	GetTimeNowUnix(hour int) int64
	GetTimeNowUnixIssued() int64
}
