package model

type NoteUpdate struct {
	Title *string `json:"title" gorm:"column:title;"`
	Content *string `json:"content" gorm:"column:content;"`
}

func (note NoteUpdate) TableName() string {
	return Note{}.TableName()
}
