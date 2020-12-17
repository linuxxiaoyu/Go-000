package biz

import (
	"github.com/linuxxiaoyu/Go-000/Week04/internal/data"
	"github.com/pkg/errors"
)

func AddItem(name string, price float32) (int64, error) {
	_, err := data.GetItem(name)
	if err == nil {
		return 0, errors.WithStack(errors.New("Alreay Exist"))
	}
	if err != nil && data.IsNoRows(err) {
		return data.AddItem(name, price)
	}
	return 0, err
}
