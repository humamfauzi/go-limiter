package limiter

import (
	"time"
)

type leakyBucket struct {
	limit    int
	duration time.Duration

	count  *int
	ticker <-chan time.Time
}

// NewLeakyBucket returns a new Limiter that allows up to limit events to happen in duration time.
// If more events happen, IsAllowed will return false.
// duration represent event/duration so if you set duration to 1 second, you will have 1 event per second.
// longest limitation duration is limit multiplied by duration.
func NewLeakyBucket(limit int, duration time.Duration) Limiter {
	ticker := time.NewTicker(duration).C
	count := new(int)
	go func(count *int) {
		for {
			<-ticker
			if *count > 0 {
				*count--
			}
		}
	}(count)
	return &leakyBucket{
		limit:    limit,
		duration: duration,
		ticker:   ticker,
		count:    count,
	}
}

func (l *leakyBucket) GetCount() int {
	return *l.count
}

func (l *leakyBucket) AddCount() {
	*l.count++
}

func (l *leakyBucket) IsAllowed() bool {
	return *l.count < l.limit
}
