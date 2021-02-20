package usertransport

import (
	"gin-project/appctx"
	"gin-project/appctx/hasher"
	"gin-project/appctx/tokenprovider/jwt"
	"gin-project/common"
	"gin-project/module/user/userbusiness"
	"gin-project/module/user/usermodel"
	"gin-project/module/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginUser(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := ctx.ShouldBind(&loginUserData); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		db := appCtx.GetDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewMySQLStore(db)
		md5 := hasher.NewMd5Hash()

		business := userbusiness.NewLoginUseCase(store, tokenProvider, md5, appctx.NewTokenConfig())
		account, err := business.Login(ctx.Request.Context(), &loginUserData)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
