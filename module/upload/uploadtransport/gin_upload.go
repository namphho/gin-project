package uploadtransport

import (
	"gin-project/appctx"
	"gin-project/common"
	"gin-project/module/upload/uploadbusiness"
	"gin-project/module/upload/uploadstore"
	"github.com/gin-gonic/gin"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
)

func Upload(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		fileHeader, err := ctx.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := ctx.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		_ = file.Close()

		//upload to s3
		store := uploadstore.NewSQLStore(appCtx.GetDBConnection())
		useCase := uploadbusiness.NewUploadUseCase(appCtx.UploadProvider(), store)
		img, err := useCase.Upload(ctx.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(img.Id))
	}
}
