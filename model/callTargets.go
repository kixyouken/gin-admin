package model

import (
	"time"

	"gorm.io/gorm"
)

type CallTargetsModel struct {
	gorm.Model
	CallID    uint
	UserID    uint
	CreatedAt time.Time      `gorm:"-"`
	UpdatedAt time.Time      `gorm:"-"`
	DeletedAt gorm.DeletedAt `gorm:"-"`
}

func (CallTargetsModel) TableName() string {
	return "call_targets"
}
