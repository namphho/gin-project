package uploadbusiness

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"gin-project/appctx/uploadprovider"
	"gin-project/common"
	"gin-project/module/upload/uploadmodel"
	"github.com/gabriel-vasile/mimetype"
	"image"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type UploadUseCase struct {
	provider uploadprovider.UploadProvider
}

func NewUploadUseCase(provider uploadprovider.UploadProvider) *UploadUseCase {
	return &UploadUseCase{provider: provider}
}

func (usecase *UploadUseCase) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {

	//check file is image
	err := IsImage(bytes.NewReader(data))
	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	//detect size
	w, h, err := GetImageDimension(bytes.NewBuffer(data))
	if err != nil {
		return nil, uploadmodel.ErrNoFileConfig
	}

	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	//upload
	img, err := usecase.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	img.Extension = fileExt

	return img, nil
}

func GetImageDimension(reader io.Reader) (int, int, error) {
	config, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Printf("err: %s", err.Error())
		return 0, 0, err
	}
	return config.Width, config.Height, nil
}

func IsImage(reader io.Reader) error {
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
