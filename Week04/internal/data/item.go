package data

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"
)

type Item struct {
	ID    int64
	Name  string
	Price float32
}

func IsNoRows(err error) bool {
	return err == sql.ErrNoRows
}

func GetItem(name string) (Item, error) {
	item := Item{}
	switch name {
	case "":
		return item, errors.Wrap(sql.ErrNoRows, "name: "+name)
	case "err":
		return item, errors.Wrap(errors.New("other error"), "name: "+name)
	default:
		item.ID = time.Now().Unix()
		item.Name = name
		item.Price = 3099
		return item, nil
	}
}

func AddItem(name string, price float32) (int64, error) {
	return time.Now().Unix(), nil
}
