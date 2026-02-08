package cache

import (
	"context"
	"fmt"
	"sync"
	"time"
	. "web/conf/cache"

	"github.com/redis/go-redis/v9"
)

// RedisClient 封装 Redis 客户端
type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

var (
	instance *RedisClient
	once     sync.Once
)

// NewRedisClient 初始化 Redis
func NewRedisClient() (*RedisClient, error) {
	cfg := GetCacheInfo()
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		DialTimeout:  time.Duration(cfg.DialTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
	})

	// 测试连接
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis connect failed: %w", err)
	}

	instance = &RedisClient{
		client: rdb,
		ctx:    ctx,
	}

	return instance, nil
}

// Set 设置 key
func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(r.ctx, key, value, expiration).Err()
}

// Get 获取 key
func (r *RedisClient) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

// Del 删除 key
func (r *RedisClient) Del(key string) error {
	return r.client.Del(r.ctx, key).Err()
}

// Incr 自增
func (r *RedisClient) Incr(key string) (int64, error) {
	return r.client.Incr(r.ctx, key).Result()
}

// Exists 判断 key 是否存在
func (r *RedisClient) Exists(key string) (bool, error) {
	cnt, err := r.client.Exists(r.ctx, key).Result()
	return cnt > 0, err
}

// Close 关闭连接
func (r *RedisClient) Close() error {
	return r.client.Close()
}
