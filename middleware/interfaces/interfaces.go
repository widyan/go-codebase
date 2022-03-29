package interfaces

import (
	"codebase/go-codebase/middleware/model"
	"context"

	"github.com/gin-gonic/gin"
)

type ApisMiddleware interface {
	VerifikasiToken(ctx context.Context, token string) (codes int, vrf model.VerifikasiToken, err error)
}

type UsecaseMiddleware interface {
	VerifyAutorizationToken() gin.HandlerFunc
}
