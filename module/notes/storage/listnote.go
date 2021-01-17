package storage

import "gin-project/module/notes/model"

//define methods of storage
func (store *storageMySql) ListNote() ([]model.Note, error) {
	var notes []model.Note
	store.Db.Find(&notes)
	return notes, nil
}
