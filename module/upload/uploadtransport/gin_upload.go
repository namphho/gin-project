package uploadtransport

import (
	"errors"
	"fmt"
	"gin-project/appctx"
	"gin-project/common"
	"gin-project/module/upload/uploadmodel"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strings"
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
		defer file.Close()

		err = IsImage(file)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		clientFile, _ := fileHeader.Open()
		defer clientFile.Close()
		w, h, err := GetImageDimension(clientFile)
		if err != nil {
			panic(uploadmodel.ErrNoFileConfig)
		}

		fileExt := filepath.Ext(fileHeader.Filename)
		fileName := fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

		if err := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf(LOCAL_STORAGE_IMAGE_PATH, fileName)); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		resp := common.Image{
			Url: fmt.Sprintf(FILE_UPLOAD_PATH, fileName),
			Width: w,
			Height: h,
			CloudName: "localhost",
			Extension: fileExt,
		}

		ctx.JSON(http.StatusOK, gin.H{"status": resp})
	}
}

func GetImageDimension(reader io.Reader) (int, int, error) {
	config, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Printf("err: %s", err.Error())
		return 0, 0, err
	}
	return config.Width, config.Height, nil
}

//38:25
func IsImage(reader io.Reader) error{
	fileType, err := DetectFileType(reader)
	if err != nil {
		return err
	}
	isJpg := strings.Contains(fileType, "jpg")
	isPng := strings.Contains(fileType, "png")
	isJpeg := strings.Contains(fileType, "jpeg")

	if !isJpg && !isPng && !isJpeg {
		return errors.New("file is not image")
	}
	return nil

}

func DetectFileType(reader io.Reader) (string, error) {
	mime, err := mimetype.DetectReader(reader)
	if err != nil {
		log.Printf(">>>>>>>err: ", err.Error())
		return "", err
	}
	return mime.Extension(), nil
}
