package logic

import (
	"context"
	"seckill/service"
)

type Goods interface {
	// 获取商品信息
	GetGoods(ctx context.Context, goodsID string)

	// 获取库存信息
	GetStorage(ctx context.Context, goodsID string)

	// 下单
	Order(ctx context.Context, goodsID string, uid string, payPrice float64)
}

type goods struct {
	goodsSrv service.Goods
}

func NewGoods() Goods {
	return &goods{
		goodsSrv: service.NewGoods(),
	}
}

func (l *goods) GetGoods(ctx context.Context, goodsID string) {

}

func (l *goods) GetStorage(ctx context.Context, goodsID string) {

}

func (l *goods) Order(ctx context.Context, goodsID string, uid string, payPrice float64) {
	// 商品service

	// 订单service

	// 库存的service
}
