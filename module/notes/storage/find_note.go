package storage

import "gin-project/module/notes/model"

func (store *storageMySql) FindNote(id int) (*model.Note, error) {
	var note model.Note
	if err := store.Db.Where("id = ?", id).First(&note).Error; err != nil {
		return nil, err
	}
	return &note, nil
}
