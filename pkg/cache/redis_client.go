package cache

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

// Cache interface defines the required methods for a cache client
type Cache interface {
	Get(key string) (string, error)
	Set(key string, value string, exp ...time.Duration) error
	Delete(key string) error
}

// RedisCache implements the Cache interface
type RedisCache struct {
	host     string
	port     int
	db       int
	expire   time.Duration
	password string
}

// NewCacheClient instantiates a new RedisCache object
func NewCacheClient(host string, port int, db int, expire int, password string) Cache {
	return &RedisCache{
		host:     host,
		port:     port,
		db:       db,
		expire:   time.Duration(expire) * time.Second,
		password: password,
	}
}

// getClient method returns a new redis client
func (r *RedisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     r.host + ":" + strconv.Itoa(r.port),
		Password: r.password,
		DB:       r.db,
	})
}

// Get method returns the value of the given key
func (r *RedisCache) Get(key string) (string, error) {
	client := r.getClient()
	resp, err := client.Get(context.Background(), key).Result()
	return resp, err
}

// Set method inserts the value of the given key
func (r *RedisCache) Set(key string, value string, exp ...time.Duration) error {
	client := r.getClient()
	expiration := r.expire
	if len(exp) > 0 {
		expiration = exp[0]
	}
	err := client.Set(context.Background(), key, value, expiration).Err()
	return err
}

// Delete method deletes the value of the given key
func (r *RedisCache) Delete(key string) error {
	client := r.getClient()
	err := client.Del(context.Background(), key).Err()
	return err
}
