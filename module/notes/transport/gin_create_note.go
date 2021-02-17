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

func CreateNote(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var data model.NoteCreate
		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		noteStorage := storage.NewMySqlStorageInstance(db)
		useCase := business.NewInstanceCreateNoteUseCase(noteStorage)

		err := useCase.CreateNote(&data)

		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, gin.H{"id": data.Id})
	}
}
