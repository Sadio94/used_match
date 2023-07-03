package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id 			string  `gorm:"column:id;type:varchar(32);" json:"id"`     						// uuid
	Name   		string  `gorm:"column:name;type:varchar(100);" json:"name"`
	Password    string  `gorm:"column:password;type:varchar(100);" json:"password"`
	Phone       string  `gorm:"column:phone;type:varchar(20);" json:"phone"`
	Mail        string  `gorm:"column:mail;type:varchar(100);" json:"mail"`
}


func (this *User) TableName() string{
	return "user"
}
