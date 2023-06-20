package bucket

import (
	"sync"
	"time"
)

type Bucket struct {
	capacity   int
	rate       int
	token      int
	lastRefill time.Time
	mutex      sync.Mutex
}

func NewBucket(capacity, rate int) *Bucket {
	return &Bucket{
		capacity:   capacity,
		rate:       rate,
		token:      capacity,
		lastRefill: time.Now(),
	}
}

func (b *Bucket) refill() {
	now := time.Now()
	if b.token < b.capacity {
		elapsed := now.Sub(b.lastRefill)
		b.token += int(elapsed.Seconds()) * b.rate
		if b.token > b.capacity {
			b.token = b.capacity
		}
		b.lastRefill = now
	}
}

func (b *Bucket) Allow() bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.refill()
	if b.token > 0 {
		b.token--
		return true
	}
	return false
}
