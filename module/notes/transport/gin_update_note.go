package transport

import (
	"gin-project/appctx"
	"gin-project/common"
	"gin-project/module/notes/business"
	"gin-project/module/notes/model"
	"gin-project/module/notes/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateNote(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idString := ctx.Param("note-id")
		uid, err := common.FromBase58(idString)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data model.NoteUpdate
		if err := ctx.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		noteStorage := storage.NewMySqlStorageInstance(db)
		updateNoteUseCase := business.NewInstanceUpdateNoteUseCase(noteStorage)

		if err := updateNoteUseCase.UpdateNoteById(int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, gin.H{"note-id": idString})
	}
}
