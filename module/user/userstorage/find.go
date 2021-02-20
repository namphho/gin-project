package userstorage


import (
	"context"
	"gin-project/common"
	"gin-project/module/user/usermodel"
)

func (s *mySQLStore) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)  {
	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreInfo{
		db = db.Preload(moreInfo[i])
	}

	var user usermodel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &user, nil
}