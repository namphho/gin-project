package transport

import (
	"gin-project/module/appctx"
	"gin-project/module/common"
	"gin-project/module/notes/business"
	"gin-project/module/notes/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetNotes(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		var paging common.Paging
		ctx.ShouldBind(&paging)
		paging.Fulfill()

		db := appCtx.GetDBConnection()
		mysqlStorage := storage.NewInstance(db)
		listNoteUseCase := business.NewInstance(mysqlStorage)

		//create paging model

		notes, err := listNoteUseCase.GetAllNotes(&paging)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(notes, paging, nil))
	}
}
