package common

import (
	"time"
)

type SQLModel struct {
	Id        int        `json:"-" gorm:"column:id;"`
	FakeId    *UID       `json:"id" gorm:"-"`
	Status    int        `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (model *SQLModel) GenUID(objectType int, shardId uint32)  {
	uid := NewUID(uint32(model.Id), objectType, shardId)
	model.FakeId = &uid
}

type SQLModelCreate struct {
	Id        int        `json:"-" gorm:"column:id;"`
	FakeId    *UID       `json:"id" gorm:"-"`
	Status    int        `json:"status" gorm:"column:status;default:1"`
}

func (model *SQLModelCreate) GenUID(objectType int, shardId uint32) {
	uid := NewUID(uint32(model.Id), objectType, shardId)
	model.FakeId = &uid
}

