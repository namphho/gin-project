package userbusiness

import (
	"context"
	"gin-project/appctx/hasher"
	"gin-project/common"
	"gin-project/module/user/usermodel"
)

type IRegisterUserCase interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}


type registerUseCase struct {
	usecase IRegisterUserCase
	hasher hasher.Hasher
}

func NewRegisterUseCaseInstance(usecase IRegisterUserCase, hasher hasher.Hasher) *registerUseCase {
	return &registerUseCase{usecase: usecase, hasher: hasher}
}

func (u *registerUseCase) RegisterUser(ctx context.Context, data *usermodel.UserCreate)  error {
	user, err := u.usecase.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return common.ErrEntityExisted(usermodel.EntityName, err)
	}

	salt := common.GenSalt(50)
	data.Password = u.hasher.Hash(data.Password + salt)
	data.Salt = salt

	role := common.RoleUser
	data.Role = &role

	if err := u.usecase.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	return nil
}