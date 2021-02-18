package transport

import (
	"gin-project/module/appctx"
	"gin-project/module/common"
	"gin-project/module/notes/business"
	"gin-project/module/notes/model"
	"gin-project/module/notes/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteNote(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idString := ctx.Param("note-id")
		uid, err := common.FromBase58(idString)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		mysqlStorage := storage.NewMySqlStorageInstance(db)
		useCase := business.NewInstanceDeleteUseCase(mysqlStorage)
		if err := useCase.DeleteNote(int(uid.GetLocalID())); err != nil {
			panic(common.ErrCannotDeleteEntity(model.Note{}.TableName(), err))
		}
		ctx.JSON(http.StatusOK, gin.H{"data": "okay"})
	}
}
