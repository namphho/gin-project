package storage

import "gin-project/module/notes/model"

func (store *storageMySql) Create(content *model.NoteCreate) error{
	db := store.Db
	db.Where(model.NoteCreate{}.TableName()).Create(&content)
	return nil
}