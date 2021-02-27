package uploadtransport

import (
	"fmt"
	"gin-project/appctx"
	"gin-project/common"
	"gin-project/module/upload/uploadmodel"
	"github.com/gin-gonic/gin"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

const FILE_UPLOAD_PATH = "http://localhost:8080/upload/%s"
const LOCAL_STORAGE_IMAGE_PATH = "./temp/%s"

func Upload(appCtx appctx.AppContext) func(ctx *gin.Context){
	return func(ctx *gin.Context) {
		fileHeader, err := ctx.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		file, err := fileHeader.Open()
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		w, h, err := GetImageDimension(file)
		if err != nil {
			panic(uploadmodel.ErrNoFileConfig)
		}

		fileName := fmt.Sprintf("%d%s", time.Now().Nanosecond(), filepath.Ext(fileHeader.Filename))

		if err := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf(LOCAL_STORAGE_IMAGE_PATH, fileName)); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		resp := common.Image{
			Url: fmt.Sprintf(FILE_UPLOAD_PATH, fileName),
			Width: w,
			Height: h,
			CloudName: "",
			Extension: "",
		}

		ctx.JSON(http.StatusOK, gin.H{"status": resp})
	}
}

func GetImageDimension(reader io.Reader) (int, int, error) {
	config, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Printf("err: ", err)
		return 0, 0, err
	}
	return config.Width, config.Height, nil
}
