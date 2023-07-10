package models

import (
	"fmt"
	. "github.com/Sadio94/micro_service/internal/intialize"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DBClient *gorm.DB

type BaseModel struct {
	Id       uint `gorm:"primaryKey"`
	CreateAt time.Time
	UpdateAt time.Time
	CreateBy string
	UpdateBy string
	Version  int16
	isDel    int8 `gorm:"index"`
}

// InitDb init db
// get connection
func InitDb(dbConf *MySQLConfig) error {
	// by config
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DB)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DBClient = db

	// 额外的连接配置
	sqlDB, err := db.DB() // database/sql.DB
	if err != nil {
		return err
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(dbConf.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(dbConf.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}
