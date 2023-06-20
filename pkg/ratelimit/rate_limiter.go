package ratelimit

import (
	"Song_API/pkg/cache"
	"Song_API/pkg/ratelimit/bucket"
	"encoding/json"
)

// Rule struct defines the fields for rules of a rate limiter. It has the following fields:
// Capacity: The maximum number of tokens that can be stored in the bucket
// Rate: The rate at which the request is allowed to be made
// Path: The path of the API endpoint
// Method: The HTTP method of the API endpoint
type Rule struct {
	Capacity int
	Rate     int
	Path     string
	Method   string
}

// GetBucket function gets the token bucket from the cache if it exists, otherwise it creates a new token bucket
func GetBucket(key string, rule Rule, bucketCache cache.Cache) *bucket.Bucket {
	var tokenBucket *bucket.Bucket
	if b, err := bucketCache.Get(key); err == nil {
		json.Unmarshal([]byte(b), &tokenBucket)
	} else {
		tokenBucket = bucket.NewBucket(rule.Capacity, rule.Rate)
	}
	return tokenBucket
}
