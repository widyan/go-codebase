package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
)

type ApisMiddleware interface {
	VerifikasiToken(ctx context.Context, token string) (codes int, vrf EntityVerifikasiToken, err error)
}

type UsecaseMiddleware interface {
	VerifyAutorizationToken() gin.HandlerFunc
}
