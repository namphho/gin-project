package storage

import "gorm.io/gorm"

//define storage sql
type storageMySql struct {
	Db *gorm.DB
}

func NewInstance(db *gorm.DB) *storageMySql {
	return &storageMySql{db}
}
