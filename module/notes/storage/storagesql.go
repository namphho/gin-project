package storage

import "gorm.io/gorm"

//define userstorage sql
type storageMySql struct {
	Db *gorm.DB
}

func NewMySqlStorageInstance(db *gorm.DB) *storageMySql {
	return &storageMySql{db}
}
