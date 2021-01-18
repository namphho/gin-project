package transport

import (
	"gin-project/module/appctx"
	"gin-project/module/notes/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateNote(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		db := appCtx.GetDBConnection()
		var note model.Note
		if err := ctx.ShouldBindJSON(&note); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&note)
		ctx.JSON(http.StatusOK, note)
	}
}
