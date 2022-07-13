package usecase

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/widyan/go-codebase/helper"
	"github.com/widyan/go-codebase/middleware/entity"
	"github.com/widyan/go-codebase/middleware/interfaces"
	"github.com/widyan/go-codebase/middleware/model"
	gmodel "github.com/widyan/go-codebase/model"
	"github.com/widyan/go-codebase/responses"
	rspn "github.com/widyan/go-codebase/responses"
)

type Usecase struct {
	Repository      interfaces.RepositoryInterface
	Logger          *logrus.Logger
	Tools           interfaces.ToolsInterface
	PrivateKey      *rsa.PrivateKey
	PublicKey       *rsa.PublicKey
	Responses       rspn.GinResponses
	ExpToken        int
	ExpRefreshToken int
}

func CreateUsecase(repo interfaces.RepositoryInterface, tools interfaces.ToolsInterface, response rspn.GinResponses, logger *logrus.Logger, expToken, expRefreshToken int, privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) interfaces.UsecaseInterface {
	return &Usecase{
		Repository:      repo,
		Logger:          logger,
		Tools:           tools,
		Responses:       response,
		ExpToken:        expToken,
		ExpRefreshToken: expRefreshToken,
		PrivateKey:      privateKey,
		PublicKey:       publicKey,
	}
}

func (u *Usecase) CreateTokenServices(ctx context.Context, request model.RequestToken) (responses model.ResponsesToken, err error) {
	users, err := u.Repository.GetUserBasedOnEmail(ctx, request.Email)
	if err != nil {
		return
	}

	if len(users) == 0 {
		err = helper.SetCaptureError(gmodel.CaptureError{
			HttpCode:  http.StatusForbidden,
			ErrorCode: rspn.UserNotAllowedAccess,
		})
		return
	}

	expJwtToken, tkn := u.createTokenJwt(ctx, request, users)
	expRefreshToken, rfToken := u.createRefreshTokenJwt(ctx, request)

	return model.ResponsesToken{
		Token:               tkn,
		RefreshToken:        rfToken,
		ExpiredToken:        expJwtToken,
		ExpiredRefreshToken: expRefreshToken,
	}, nil
}

func (u *Usecase) createTokenJwt(ctx context.Context, request model.RequestToken, users []entity.User) (expJwtToken int64, tkn string) {
	var role []string = []string{}
	for _, user := range users {
		role = append(role, user.Role)
	}

	expJwtToken = u.Tools.GetTimeNowUnix(u.ExpToken)
	claims := &model.Auth{
		request.Email,
		role,
		jwt.StandardClaims{
			ExpiresAt: expJwtToken,
			Issuer:    "Dana",
			IssuedAt:  u.Tools.GetTimeNowUnixIssued(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tkn, _ = token.SignedString(u.PrivateKey)
	return
}

func (u *Usecase) createRefreshTokenJwt(ctx context.Context, request model.RequestToken) (expRefreshToken int64, rtkn string) {
	expRefreshToken = u.Tools.GetTimeNowUnix(u.ExpRefreshToken)
	claims := &model.AuthRefreshToken{
		request.Email,
		jwt.StandardClaims{
			ExpiresAt: expRefreshToken,
			Issuer:    "Dana",
			IssuedAt:  u.Tools.GetTimeNowUnixIssued(),
		},
	}

	rftoken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	rtkn, _ = rftoken.SignedString(u.PrivateKey)
	return
}

func (u *Usecase) AddUser(ctx context.Context, user model.RequestUser) (err error) {
	err = u.Repository.AddUser(ctx, entity.User{
		Email:    user.Email,
		ID:       u.Tools.GetUUID(),
		Name:     user.Name,
		Role:     user.Role,
		IsActive: user.IsActive,
	})
	if err != nil {
		return
	}
	return
}

func (u *Usecase) VerifyAutorizationToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			u.Responses.AbortWithStatusJSONAndErrorCode(c, http.StatusBadRequest, responses.TokenIsNotAllowedEmpty)
			return
		}

		// get authorization token
		tokenString := c.GetHeader("Authorization")
		strArr := strings.Split(tokenString, " ")
		if len(strArr) == 2 {
			tokenString = strArr[1]
		} else {
			u.Responses.AbortWithStatusJSONAndErrorCode(c, http.StatusForbidden, responses.InvalidToken)
			return
		}

		var claim model.Auth
		token, err := jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
			return u.PublicKey, nil
		})

		if err != nil {
			if strings.Contains(err.Error(), "token is expired") {
				u.Responses.AbortWithStatusJSONAndErrorCode(c, http.StatusForbidden, responses.TokenExpired)
				return
			}
			c.AbortWithError(http.StatusForbidden, err)
			return
		}

		if !token.Valid {
			u.Responses.AbortWithStatusJSONAndErrorCode(c, http.StatusForbidden, responses.InvalidToken)
			return
		}

		byt, _ := json.Marshal(claim)
		c.Set("bind", byt)
	}
}
