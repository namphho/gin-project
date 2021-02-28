package appctx

import (
	"gin-project/appctx/uploadprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetDBConnection() *gorm.DB
	SecretKey() string
	UploadProvider() uploadprovider.UploadProvider
}

type appContext struct {
	db             *gorm.DB
	secretKey      string
	uploadProvider uploadprovider.UploadProvider
}

func NewInstance(db *gorm.DB, secretKey string, uploadProvider uploadprovider.UploadProvider) *appContext {
	return &appContext{db: db, secretKey: secretKey, uploadProvider: uploadProvider}
}

func (appCtx *appContext) GetDBConnection() *gorm.DB {
	return appCtx.db.Session(&gorm.Session{NewDB: true})
}

func (appCtx *appContext) SecretKey() string {
	return appCtx.secretKey
}

func (appCtx *appContext) UploadProvider() uploadprovider.UploadProvider{
	return appCtx.uploadProvider
}

type tokenExpiry struct {
	atExp int
	rtExp int
}

func NewTokenConfig() tokenExpiry {
	return tokenExpiry{
		atExp: 60 * 60 * 24 * 7,     //7 days
		rtExp: 60 * 60 * 24 * 7 * 2, // 14 days
	}
}

func (tk tokenExpiry) GetAccessTokenExp() int {
	return tk.atExp
}

func (tk tokenExpiry) GetRefreshTokenExp() int {
	return tk.rtExp
}
