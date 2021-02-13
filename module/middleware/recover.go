package middleware

import (
	"errors"
	"fmt"
	"gin-project/module/appctx"
	"gin-project/module/common"
	"github.com/gin-gonic/gin"
)

func Recover(ctx appctx.AppContext) gin.HandlerFunc{
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				context.Header("Content-type", "application/json")
				if appErr, ok := err.(*common.AppError); ok {
					context.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
					return
				}

				appErr := common.ErrInternal(errors.New(fmt.Sprintf("%v", err)))
				context.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
				return
			}

		}()
		context.Next()
	}
}