package transport

import (
	"gin-project/appctx"
	"gin-project/common"
	"gin-project/module/notes/business"
	"gin-project/module/notes/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetNoteById(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		idString := ctx.Param("note-id")
		uid, err := common.FromBase58(idString)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetDBConnection()
		noteStorage := storage.NewMySqlStorageInstance(db)
		getNoteUseCase := business.NewInstanceGetNoteUseCase(noteStorage)

		note, err := getNoteUseCase.GetNote(int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}
		note.GenUID(common.DBTypeNote, common.ShardId)
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(note))
	}
}
