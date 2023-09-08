package limiter

import (
	"testing"
	"time"
)

func TestLeakyBucket(t *testing.T) {
	l := NewLeakyBucket(10, time.Millisecond)
	for i := 0; i < 10; i++ {
		l.AddCount()
	}
	time.Sleep(time.Millisecond)
	if l.IsAllowed() {
		t.Errorf("LeakyBucket should not be allowed")
	}
	time.Sleep(11 * time.Millisecond)
	if !l.IsAllowed() {
		t.Errorf("LeakyBucket should be allowed with limit %d and current count %d", 10, l.GetCount())
	}
}
