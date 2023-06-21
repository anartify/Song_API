package bucket

import (
	"time"
)

// Bucket struct defines the fields for implementation of token bucket algorithm.
type Bucket struct {
	Capacity   int
	Rate       int
	Token      int
	LastRefill time.Time
}

// NewBucket function instantiates a new token bucket
func NewBucket(capacity, rate int) *Bucket {
	return &Bucket{
		Capacity:   capacity,
		Rate:       rate,
		Token:      capacity,
		LastRefill: time.Now(),
	}
}

// refill function refills the token bucket according to the rate and elapsed time between requests
func (b *Bucket) refill() {
	now := time.Now()
	if b.Token < b.Capacity {
		elapsed := now.Sub(b.LastRefill)
		b.Token += int(elapsed.Seconds()) * b.Rate
		if b.Token > b.Capacity {
			b.Token = b.Capacity
		}
		b.LastRefill = now
	}
}

// Allow function checks if the token bucket has enough tokens to allow a request
func (b *Bucket) Allow() bool {
	b.refill()
	if b.Token > 0 {
		b.Token--
		return true
	}
	return false
}
