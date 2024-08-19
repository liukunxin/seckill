package model

type Goods struct {
	GoodsID     uint64 `json:"goods_id" gorm:"goods_id"`
	Description string `json:"description"`
}
