package service

import (
	"github.com/vera-byte/cool-admin-go/cool"
	"github.com/vera-byte/cool-admin-go/modules/demo/model"
)

type DemoGoodsService struct {
	*cool.Service
}

func NewDemoGoodsService() *DemoGoodsService {
	return &DemoGoodsService{
		&cool.Service{
			Model: model.NewDemoGoods(),
			ListQueryOp: &cool.QueryOp{

				Join: []*cool.JoinOp{},
			},
		},
	}
}
