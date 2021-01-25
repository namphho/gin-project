package storage

import (
	"gin-project/module/common"
	"gin-project/module/notes/model"
)

//define methods of storage
func (store *storageMySql) ListNote(filter *model.Filter, paging *common.Paging) ([]model.Note, error) {
	db := store.Db
	var notes []model.Note

	db = db.Table(model.Note{}.TableName()).Where("status not in (0)")

	if v := filter; v != nil {
		if v.CategoryId > 0 {
			db.Where("category_id = ?", v.CategoryId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}
	db.Limit(paging.Limit)
	db.Offset((paging.Page - 1) * paging.Limit)

	if err := db.Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
}
