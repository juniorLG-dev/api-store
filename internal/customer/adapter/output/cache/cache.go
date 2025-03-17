package cache 

import (
	"loja/internal/customer/adapter/output/model/customer_cache"

	"context"
	"encoding/json"
	"time"

	redis "github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type cache struct {
	rdb *redis.Client
}

func NewCustomerCache(rdb *redis.Client) *cache {
	return &cache{
		rdb: rdb,
	}
}

type PortCache interface {
	SetCache(customer_cache.InfoCustomer) error
	GetCache(string) (*customer_cache.InfoCustomer, error)
}

func (c *cache) SetCache(infoCustomer customer_cache.InfoCustomer) error {
	infoCustomerJSON, err := json.Marshal(infoCustomer)

	err = c.rdb.Set(ctx, infoCustomer.Email, infoCustomerJSON, 5*time.Minute).Err()

	return err
}

func (c *cache) GetCache(email string) (*customer_cache.InfoCustomer, error) {
	customerCache, err := c.rdb.Get(ctx, email).Result()
	
	var infoCustomer customer_cache.InfoCustomer

	err = json.Unmarshal([]byte(customerCache), &infoCustomer)

	return &infoCustomer, err
}