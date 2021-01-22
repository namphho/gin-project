package storage

func (store *storageMySql) DeleteNote(id int) error {
	db := store.Db
	if err := db.Table("notes").Where("id = ?", id).Update("status", 0).Error; err != nil {
		return err
	}
	return nil
}
