package storage

import "gin-project/module/notes/model"

func (store *storageMySql) FindNoteWithKeys(id int, moreKeys ...string) (*model.Note, error) {
	var note model.Note
	db := store.Db

	for _, key := range moreKeys {
		db = db.Preload(key)
	}

	if err := db.Where("id = ?", id).First(&note).Error; err != nil {
		return nil, err
	}
	return &note, nil
}
