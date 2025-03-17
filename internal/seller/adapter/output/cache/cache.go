package cache

import (
	redis "github.com/redis/go-redis/v9"
	"loja/internal/seller/adapter/output/model/seller_cache"

	"time"
	"context"
	"encoding/json"
)

var ctx = context.Background()

type cache struct {
	rdb *redis.Client
}

func NewCache(rdb *redis.Client) *cache {
	return &cache{
		rdb: rdb,
	}
}

type PortCache interface {
	SetCache(seller_cache.InfoSeller) error
	GetCache(string) (*seller_cache.InfoSeller, error)
}

func (c *cache) SetCache(infoSeller seller_cache.InfoSeller) error {
	infoSellerJSON, err := json.Marshal(infoSeller)

	err = c.rdb.Set(ctx, infoSeller.Email, infoSellerJSON, 5*time.Minute).Err()

	return err
}

func (c *cache) GetCache(email string) (*seller_cache.InfoSeller, error) {
	sellerCache, err := c.rdb.Get(ctx, email).Result()

	var infoSeller seller_cache.InfoSeller

	err = json.Unmarshal([]byte(sellerCache), &infoSeller)

	return &infoSeller, err
}