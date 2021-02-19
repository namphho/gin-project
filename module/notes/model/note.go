package model

import "gin-project/common"

type Note struct {
	common.SQLModel `json:",inline"`
	Title           string         `json:"title" gorm:"column:title;"`
	Content         string         `json:"content" gorm:"column:content;"`
	Cover           *common.Image  `json:"cover" gorm:"column:cover;"`
	Photos          *common.Images `json:"photos" gorm:"column:photos;"`
}

func (n Note) TableName() string {
	return "notes"
}

func (n Note) EntityName() string {
	return n.TableName()
}
