package model

type NoteCreate struct {
	Id int `json:"id" gorm:"column:id;"`
	Title string `json:"title" gorm:"column:title;"`
	Content string `json:"content" gorm:"column:content;"`
}

func (note NoteCreate) TableName() string {
	return Note{}.TableName()
}
