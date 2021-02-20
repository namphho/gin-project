package userstorage

import (
	"context"
	"gin-project/common"
	"gin-project/module/user/usermodel"
)

func (s *mySQLStore) CreateUser(ctx context.Context, data *usermodel.UserCreate)  error {
	db := s.db.Begin()

	if err := db.Table(data.TableName()).Create(&data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}