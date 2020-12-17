package service

import (
	"context"

	"github.com/linuxxiaoyu/Go-000/Week04/api"
	"github.com/linuxxiaoyu/Go-000/Week04/internal/biz"
)

type ItemService struct{}

func (item *ItemService) AddItem(ctx context.Context, r *api.ItemRequest) (*api.ItemResponse, error) {
	var resp api.ItemResponse
	id, err := biz.AddItem(r.Name, r.Price)
	if err == nil {
		resp.Id = id
	}
	return &resp, err
}
