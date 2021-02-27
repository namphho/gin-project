package uploadmodel

import (
	"errors"
	"gin-project/common"
)

const UploadName = "Upload"

type Upload struct {
	common.SQLModel `json:",inline"`
	common.Image    `json:",inline"`
}

func (Upload) TableName() string {
	return "upload"
}

func (u *Upload) Mask(isAdmin bool) {
	u.GenUID(common.DBTypeUser, common.ShardId)
}

var (
	ErrFileTooLarge = common.NewCustomError(
		errors.New("file is too large"),
		"file is too large",
		"ErrFileIsTooLarge")

	ErrNoFileConfig = common.NewCustomError(
		errors.New("can't decode file config"),
		"can't decode file config",
		"ErrDecodeFileConfig")
)
