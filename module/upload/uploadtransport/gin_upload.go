package uploadtransport

import (
	"bytes"
	"errors"
	"fmt"
	"gin-project/appctx"
	"gin-project/common"
	"gin-project/module/upload/uploadmodel"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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


//S3_BUCKET_NAME=g02-food-delivery;
//S3_REGION=ap-southeast-1;
//S3_API_KEY=AKIAWLDHV2FT2ZORCW2B;
//S3_SECRET=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;
//S3_DOMAIN=d124uuz544hfp3.cloudfront.net;

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

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		_ = file.Close()

		//create byte buffer
		fileBytes := bytes.NewBuffer(dataBytes)

		//check file is image or note
		err = IsImage(fileBytes)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		//get dimension
		w, h, err := GetImageDimension(bytes.NewBuffer(dataBytes))
		if err != nil {
			panic(uploadmodel.ErrNoFileConfig)
		}

		fileExt := filepath.Ext(fileHeader.Filename)
		fileName := fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

		//if err := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf(LOCAL_STORAGE_IMAGE_PATH, fileName)); err != nil {
		//	panic(common.ErrInvalidRequest(err))
		//}

		//upload to s3
		//create a single AWS session (we can re use this if we're uploading many files)
		mySession, err := session.NewSession(&aws.Config{
			Region: aws.String("ap-southeast-1"),
			Credentials: credentials.NewStaticCredentials(
				"",
				"",
				"",
				),
		})

		_ , err = s3.New(mySession).PutObject(&s3.PutObjectInput{
			Bucket: aws.String("g02-food-delivery"),
			Key: aws.String(fileName),
			ACL: aws.String("private"),
			Body: bytes.NewReader(dataBytes),
			ContentLength: aws.Int64(fileHeader.Size),
		})

		if err != nil {
			panic(common.ErrInvalidRequest(errors.New("can't now upload")))
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
