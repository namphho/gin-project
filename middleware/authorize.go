package middleware

import (
	"errors"
	"fmt"
	"gin-project/appctx"
	"gin-project/appctx/tokenprovider/jwt"
	"gin-project/common"
	"gin-project/module/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	AUTHORIZATION = "Authorization"
	BEARER        = "Bearer"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(value string) (string, error) {
	parts := strings.Split(value, " ")

	if parts[0] != BEARER || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}
	return parts[1], nil
}

func RequiredAuth(appCtx appctx.AppContext) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader(AUTHORIZATION))

		if err != nil {
			panic(err)
		}

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		db := appCtx.GetDBConnection()
		store := userstorage.NewMySQLStore(db)

		user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("status is zero")))
		}

		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
