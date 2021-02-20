
package userbusiness

import (
	"context"
	"gin-project/appctx"
	"gin-project/appctx/hasher"
	"gin-project/appctx/tokenprovider"
	"gin-project/common"
	"gin-project/module/user/usermodel"
)

type ILoginUseCase interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type TokenConfig interface {
	GetAccessTokenExp() int
	GetRefreshTokenExp() int
}

type LoginUseCase struct {
	appCtx appctx.AppContext
	useCase ILoginUseCase
	provider tokenprovider.Provider
	hashser hasher.Hasher
	tokenCfg TokenConfig
}

func NewLoginUseCase(useCase ILoginUseCase, provider tokenprovider.Provider, hasher hasher.Hasher, config TokenConfig) *LoginUseCase {
	return &LoginUseCase{
		useCase: useCase,
		provider: provider,
		hashser: hasher,
		tokenCfg: config,
	}
}

func (u *LoginUseCase) Login(ctx context.Context, data *usermodel.UserLogin) (*usermodel.Account, error){

	user, err := u.useCase.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, common.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	passHashed := u.hashser.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role: user.Role.String(),
	}

	accessToken, err := u.provider.Generate(payload, u.tokenCfg.GetAccessTokenExp())
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, err := u.provider.Generate(payload, u.tokenCfg.GetRefreshTokenExp())
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := usermodel.NewAccount(accessToken, refreshToken)

	return account, nil

}
