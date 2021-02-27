package uploadtransport

import (
	"gin-project/appctx"
	"gin-project/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(appCtx appctx.AppContext) func(ctx *gin.Context){
	return func(ctx *gin.Context) {
		fileHeader, err := ctx.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		err = ctx.SaveUploadedFile(fileHeader, "./" + fileHeader.Filename)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
