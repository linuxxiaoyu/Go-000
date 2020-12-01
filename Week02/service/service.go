package service

import (
	"github.com/linuxxiaoyu/Go-000/Week02/dao"
)

func Item(id int) (*dao.Item, error) {
	return dao.GetItem(id)
}
