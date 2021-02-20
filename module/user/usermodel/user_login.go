package usermodel

import "gin-project/appctx/tokenprovider"

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"-"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

type Account struct {
	AccessToken *tokenprovider.Token `json:"access_token"`
	RefreshToken *tokenprovider.Token `json:"refresh_token"`
}

func NewAccount(accessToken, refreshToken *tokenprovider.Token)  *Account{
	return &Account{
		accessToken,
		refreshToken,
	}
}