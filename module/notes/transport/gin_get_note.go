package transport

import (
	"gin-project/module/appctx"
	"gin-project/module/notes/business"
	"gin-project/module/notes/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetNotes(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		db := appCtx.GetDBConnection()
		mysqlStorage := storage.NewInstance(db)
		listNoteUseCase := business.NewInstance(mysqlStorage)
		notes, _ := listNoteUseCase.GetAllNotes()

		ctx.JSON(http.StatusOK, notes)
	}
}
