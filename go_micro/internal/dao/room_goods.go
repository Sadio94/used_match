/**
 * @Author: yinjie_luo
 * @Description:
 * @File:  room_goods
 * @Date: 2023/07/06 10:47
 */

package dao

import "github.com/Sadio94/go_micro/internal/models"

var RoomGoodsDaoOperate = RoomGoodsDao{}

type RoomGoodsDao struct{}

func (this *RoomGoodsDao) Insert(goods *models.RoomGoods) (string, error) {

	return "", nil
}
