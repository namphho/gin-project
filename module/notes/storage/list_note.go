package storage

import (
	"gin-project/common"
	"gin-project/module/notes/model"
)

//define methods of userstorage
func (store *storageMySql) ListNote(filter *model.Filter, paging *common.Paging, moreKeys ...string) ([]model.Note, error) {
	db := store.Db
	var notes []model.Note

	db = db.Table(model.Note{}.TableName()).Where("status not in (0)")

	if v := filter; v != nil {
		if v.CategoryId > 0 {
			db.Where("category_id = ?", v.CategoryId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	db.Limit(paging.Limit)

	for _, k := range moreKeys {
		db = db.Preload(k)
	}

	db.Offset((paging.Page - 1) * paging.Limit)

	if err := db.Order("id desc").Find(&notes).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return notes, nil
}
