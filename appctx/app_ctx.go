package appctx

import "gorm.io/gorm"

type AppContext interface {
	GetDBConnection() *gorm.DB
}

type appContext struct {
	db *gorm.DB
}

func NewInstance(db *gorm.DB) *appContext {
	return &appContext{db: db}
}

func (appCtx *appContext) GetDBConnection() *gorm.DB {
	return appCtx.db.Session(&gorm.Session{NewDB: true})
}
