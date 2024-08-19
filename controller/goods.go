package controller

import (
	"net/http"
	"seckill/logic"
)
import "github.com/gin-gonic/gin"

type GoodsController interface {
	GoodsInfo(ctx *gin.Context)
	StockInfo(ctx *gin.Context)
	Order(ctx *gin.Context)
}

type goods struct {
	goodsLogic logic.Goods
}

func NewGoods() GoodsController {
	return &goods{
		goodsLogic: logic.NewGoods(),
	}
}

// 采用go validate 做参数校验

func (c *goods) GoodsInfo(ctx *gin.Context) {
	//var param params.PopReq
	//if err := ctx.ShouldBindJSON(&param); err != nil {
	//
	//}
	ctx.JSON(http.StatusOK, gin.H{"code": 100, "msg": "param is err"})
}

func (c *goods) StockInfo(ctx *gin.Context) {
	//var param params.PopReq
	//if err := ctx.ShouldBindJSON(&param); err != nil {
	//
	//}
	ctx.JSON(http.StatusOK, gin.H{"code": 100, "msg": "param is err"})
}

func (c *goods) Order(ctx *gin.Context) {
	//var param params.PopReq
	//if err := ctx.ShouldBindJSON(&param); err != nil {
	//
	//}
	ctx.JSON(http.StatusOK, gin.H{"code": 100, "msg": "param is err"})
}
