package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Id 			string  `gorm:"column:id;type:varchar(32);" json:"id"`     						// uuid
	Name   		string  `gorm:"column:name;type:varchar(100);" json:"name"`    		// 分类名称
	ParentId    string  `gorm:"column:parent_id;type:int(11);" json:"parentId"`	// 父级id
}


func (this *Category) TableName() string{
	return "category"
}

