package storage

import (
	"gin-project/module/notes/model"
)

func (store *storageMySql) UpdateNote(noteId int, content *model.NoteUpdate) error{
	db := store.Db
	if err := db.Table(model.NoteUpdate{}.TableName()).Where("id = ?", noteId).Updates(content).Error; err != nil {
		return err
	}
	return nil
}