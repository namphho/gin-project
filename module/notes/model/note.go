package model

type Note struct {
	Id      int    `json:"id"`
	Title   string `json:"title" gorm:"column:title;"`
	Status  int    `json:"status" gorm:"column:status;"`
	Content string `json:"content" gorm:"column:content;"`
}
