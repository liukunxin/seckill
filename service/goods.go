package service

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"seckill/dao"
	"seckill/dao/model"
	redis2 "seckill/storge/redis"
	"time"
)

// 单一的动作
type Goods interface {
	GetGoods(ctx context.Context, goodsID string) (*model.Goods, error)
}

type goods struct {
	goodsDao    dao.Goods
	redisClient *redis.Client
}

func NewGoods() Goods {
	return &goods{
		goodsDao:    dao.New(),
		redisClient: redis2.GetRedisClient(),
	}
}

// / 业务功能单一的方法
func (s *goods) GetGoods(ctx context.Context, goodsID string) (*model.Goods, error) {
	goodsCache, _ := s.redisClient.Get(ctx, goodsID).Result()
	if goodsCache == "" {
		data, err := s.goodsDao.GetGoods(ctx, goodsID)
		if err != nil {
			return nil, err
		}
		// data 进行redis 缓存
		dataBytes, _ := json.Marshal(data)
		s.redisClient.Set(ctx, goodsID, dataBytes, 10*time.Second)
	}
	var goodsCacheTmp *model.Goods
	json.Unmarshal([]byte(goodsCache), &goodsCacheTmp)
	return goodsCacheTmp, nil

}
