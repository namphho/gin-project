package usertransport

import (
	"gin-project/appctx"
	"gin-project/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProfile(appCtx appctx.AppContext) func(ctx *gin.Context){
	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser).(common.Masker)
		data.Mask(true)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
