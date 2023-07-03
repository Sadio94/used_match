package models

import "gorm.io/gorm"


type Problem struct {
	gorm.Model
	Id 			string `gorm:"column:id;type:varchar(32);" json:"id"`     						// uuid
	CategoryId	string `gorm:"column:category_id;type:varchar(255);" json:"categoryId"`			// 分类id 以逗号分隔
	Title		string `gorm:"column:title;type:varchar(255);" json:"title"`					// 文章标题
	Content     string `gorm:"content:title;type:varchar(255);" json:"content"`					// 文章内容
	MaxRunTime  int    `gorm:"content:max_runTime;type:int(11);" json:"maxRunTime"`       		// 最大运行时间
	MaxMem     int 	   `gorm:"content:max_mem;type:int(11);" json:"maxMem"`						// 最大运行内存
}


func (this *Problem) TableName() string{
	return "problem"
}
