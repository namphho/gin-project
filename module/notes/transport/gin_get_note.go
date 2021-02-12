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

func GetNotes(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var paging common.Paging
		var filter model.Filter

		ctx.ShouldBind(&paging)
		ctx.ShouldBind(&filter)

		paging.Fulfill()

		db := appCtx.GetDBConnection()
		noteStorage := storage.NewInstance(db)
		listNoteUseCase := business.NewInstance(noteStorage)

		//create paging model

		notes, err := listNoteUseCase.GetAllNotes(&filter, &paging)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(notes, paging, filter))
	}
}
