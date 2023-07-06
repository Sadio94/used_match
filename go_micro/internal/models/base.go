package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DBClient *gorm.DB

type BaseModel struct {
	Id       uint      `gorm:"column:id;primaryKey" json:"id"`
	CreateAt time.Time `gorm:"column:create_at;autoCreateTime;" json:"createAt"`
	CreateBy string
	UpdateAt time.Time `gorm:"column:update_at;autoUpdateTime;" json:"updateAt"`
	UpdateBy string
	Version  uint64
	IsDel    bool `gorm:"column:id_del;index;" json:"idDel"`
}

// InitDb init db
// get connection
func InitDb() error {
	// by config
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "admin", "123456", 3306, "db1")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DBClient = db

	return nil
}
