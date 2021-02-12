package transport

import (
	"gin-project/module/appctx"
	"gin-project/module/notes/business"
	"gin-project/module/notes/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteNote(appCtx appctx.AppContext) func(ctx *gin.Context){
	return func(ctx *gin.Context) {
		idString := ctx.Param("note-id")
		id, _ := strconv.Atoi(idString)

		db := appCtx.GetDBConnection()
		mysqlStorage := storage.NewMySqlStorageInstance(db)
		useCase := business.NewInstanceDeleteUseCase(mysqlStorage)
		if err := useCase.DeleteNote(id); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": "okay"})
	}
}