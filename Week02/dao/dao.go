package dao

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

type Item struct {
	ID   int
	Name string
}

var errOther = errors.New("dao: other errors")

func GetItem(id int) (*Item, error) {
	sqlStr := fmt.Sprintf("SELECT * FROM items WHERE id = %d", id)
	switch id {
	case 0:
		return nil, errors.Wrap(sql.ErrNoRows, sqlStr)
	case 1:
		return nil, errors.Wrap(errOther, sqlStr)
	default:
		return &Item{
			ID:   id,
			Name: fmt.Sprintf("item-%d", id),
		}, nil
	}
}
