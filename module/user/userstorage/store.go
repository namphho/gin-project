package userstorage

import "gorm.io/gorm"

type mySQLStore struct {
	db *gorm.DB
}

func NewMySQLStore(db *gorm.DB) *mySQLStore {
	return &mySQLStore{db: db}
}