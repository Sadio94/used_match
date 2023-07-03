package models

import "gorm.io/gorm"

type Submit struct {
	gorm.Model
	Id 			string  	`gorm:"column:id;type:varchar(32);" json:"id"`     						// uuid
	ProblemId   string		`gorm:"column:problem_id;type:varchar(32);" json:"problemId"`
	UserId   	string		`gorm:"column:user_id;type:varchar(32);" json:"userId"`
	Path  		string 		`gorm:"column:path;type:varchar(255);" json:"path"`						// 代码存放路径
}





func (this *Submit) TableName() string{
	return "submit"
}
