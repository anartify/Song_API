package ratelimit

import "Song_API/pkg/ratelimit/bucket"

type Rule struct {
	Capacity int
	Rate     int
	Path     string
	Method   string
}

var clientBucket = make(map[string]*bucket.Bucket)

func GetBucket(key string, rule Rule) *bucket.Bucket {
	if _, ok := clientBucket[key]; !ok {
		clientBucket[key] = bucket.NewBucket(rule.Capacity, rule.Rate)
	}
	return clientBucket[key]
}
