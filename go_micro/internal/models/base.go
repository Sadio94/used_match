package models

import "time"

type BaseModel struct {
	Id uint `gorm:"primaryKey"`
	CreateAt time.Time
	UpdateAt time.Time
}
