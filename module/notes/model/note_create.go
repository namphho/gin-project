package model

import "gin-project/common"

type NoteCreate struct {
	common.SQLModelCreate `json:",inline"`
	Title                 string         `json:"title" gorm:"column:title;"`
	Content               string         `json:"content" gorm:"column:content;"`
	Cover                 *common.Image  `json:"cover" gorm:"column:cover;"`
	Photos                *common.Images `json:"photos" gorm:"column:photos;"`
}

func (model NoteCreate) TableName() string {
	return Note{}.TableName()
}