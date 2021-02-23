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

func GetNotes(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var paging common.Paging
		var filter model.Filter

		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := ctx.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		db := appCtx.GetDBConnection()
		noteStorage := storage.NewMySqlStorageInstance(db)
		listNoteUseCase := business.NewInstance(noteStorage)

		//create paging usermodel

		notes, err := listNoteUseCase.GetAllNotes(&filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range notes {
			notes[i].Mask(false)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(notes, paging, filter))
	}
}
