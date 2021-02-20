package usertransport

import (
	"gin-project/appctx"
	"gin-project/appctx/hasher"
	"gin-project/common"
	"gin-project/module/user/userbusiness"
	"gin-project/module/user/usermodel"
	"gin-project/module/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUser(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var user usermodel.UserCreate

		if err := ctx.ShouldBind(&user); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		store := userstorage.NewMySQLStore(db)
		usecase := userbusiness.NewRegisterUseCaseInstance(store, hasher.NewMd5Hash())

		if err := usecase.RegisterUser(ctx.Request.Context(), &user); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		user.GenUID(common.DBTypeUser, common.ShardId)

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(user.FakeId))
	}
}
