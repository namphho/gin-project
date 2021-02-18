package transport

import (
	"gin-project/module/appctx"
	"gin-project/module/common"
	"gin-project/module/notes/business"
	"gin-project/module/notes/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetNoteById(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idString := ctx.Param("note-id")
		noteId, err := strconv.Atoi(idString)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		noteStorage := storage.NewMySqlStorageInstance(db)
		getNoteUseCase := business.NewInstanceGetNoteUseCase(noteStorage)

		note, err := getNoteUseCase.GetNote(noteId)
		if err != nil {
			panic(err)
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(note))
	}
}
