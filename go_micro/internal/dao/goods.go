/**
 * @Author: yinjie_luo
 * @Description:
 * @File:  goods
 * @Date: 2023/07/06 10:47
 */

package dao

import "github.com/Sadio94/go_micro/internal/models"

var GoodsOperate = GoodsDao{}

type GoodsDao struct{}

func (this *GoodsDao) Insert(goods *models.Goods) (uint, error) {
	result := models.DBClient.Create(goods)
	if result.Error != nil {
		return 0, result.Error
	}

	return goods.BaseModel.Id, nil
}
