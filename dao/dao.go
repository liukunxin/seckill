package dao

import (
	"context"
	"seckill/dao/model"
)

type Goods interface {
	GetGoods(ctx context.Context, goodsID string) (model.Goods, error)
}

type goods struct {
}

func New() Goods {
	return &goods{}
}

func (d *goods) GetGoods(ctx context.Context, goodsID string) (model.Goods, error) {
	// mysql  gorm 框架去实现mysql 的读写
	return model.Goods{}, nil
}
