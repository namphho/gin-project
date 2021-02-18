package transport

import (
	"gin-project/module/appctx"
	"gin-project/module/common"
	"gin-project/module/notes/business"
	"gin-project/module/notes/model"
	"gin-project/module/notes/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateNote(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idString := ctx.Param("note-id")
		noteId, err := strconv.Atoi(idString)

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

		if err := updateNoteUseCase.UpdateNoteById(noteId, &data); err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, gin.H{"note-id": noteId})
	}
}
