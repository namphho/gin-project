package model

import "gin-project/module/common"

type NoteUpdate struct {
	Title   *string        `json:"title" gorm:"column:title;"`
	Content *string        `json:"content" gorm:"column:content;"`
	Cover   *common.Image  `json:"cover" gorm:"column:cover;"`
	Photos  *common.Images `json:"photos" gorm:"column:photos;"`
}

func (note NoteUpdate) TableName() string {
	return Note{}.TableName()
}
