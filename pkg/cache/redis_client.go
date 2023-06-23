package cache

import (
	"Song_API/pkg/apperror"
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

// Cache interface defines the required methods for a cache client
type Cache interface {
	Get(key string, value interface{}) error
	Set(key string, value interface{}, exp ...time.Duration) error
	Delete(key string) error
	AcquireLock(bucketKey string) error
	ReleaseLock(bucketKey string) error
}

// Redis implements the Cache interface
type Redis struct {
	host        string
	port        int
	db          int
	expire      time.Duration
	password    string
	lockTimeOut time.Duration
}

// NewClient instantiates a new Redis object
func NewClient(host string, port int, db int, expire int, password string) Cache {
	return &Redis{
		host:        host,
		port:        port,
		db:          db,
		expire:      time.Duration(expire) * time.Second,
		password:    password,
		lockTimeOut: 1 * time.Second,
	}
}

// getClient method returns a new redis client
func (r *Redis) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     r.host + ":" + strconv.Itoa(r.port),
		Password: r.password,
		DB:       r.db,
	})
}

// Get method takes key and interface as arguements and populates the interface with the value of the given key
func (r *Redis) Get(key string, value interface{}) error {
	client := r.getClient()
	resp, err := client.Get(context.Background(), key).Result()
	json.Unmarshal([]byte(resp), value)
	return err
}

// Set method inserts the value of the given key
func (r *Redis) Set(key string, value interface{}, exp ...time.Duration) error {
	client := r.getClient()
	expiration := r.expire
	if len(exp) > 0 {
		expiration = exp[0]
	}
	valBytes, errMarshal := json.Marshal(value)
	if errMarshal != nil {
		return &apperror.CustomError{Message: "Failed to marshal data"}
	}
	err := client.Set(context.Background(), key, valBytes, expiration).Err()
	return err
}

// Delete method deletes the value of the given key
func (r *Redis) Delete(key string) error {
	client := r.getClient()
	err := client.Del(context.Background(), key).Err()
	return err
}

// AcquireLock method acquires a lock on the given bucket
func (r *Redis) AcquireLock(bucketKey string) error {
	client := r.getClient()
	attempts := 5 // Number of times to try to acquire lock
	for i := 0; i < attempts; i++ {
		status, err := client.SetNX(context.Background(), bucketKey, "locked", r.lockTimeOut).Result()
		if status {
			return nil
		}
		if err != nil {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	return &apperror.CustomError{Message: "Failed to acquire lock"}
}

// ReleaseLock method releases the lock on the given bucket
func (r *Redis) ReleaseLock(bucketKey string) error {
	client := r.getClient()
	err := client.Del(context.Background(), bucketKey).Err()
	return err
}
