package storage

import "gorm.io/gorm"

//define storage sql
type storageMySql struct {
	Db *gorm.DB
}

func NewMySqlStorageInstance(db *gorm.DB) *storageMySql {
	return &storageMySql{db}
}
