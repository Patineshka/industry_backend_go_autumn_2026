package main

import (
	"sync"
	"time"
)

type Clock interface{ Now() time.Time }
type Limiter struct {
	mu     sync.Mutex
	clock  Clock
	rate   float64
	burst  int
	tokens float64
	last   time.Time
}

func NewLimiter(clock Clock, ratePerSec float64, burst int) *Limiter {
	limiter := &Limiter{clock: clock, rate: ratePerSec, burst: burst}
	if burst <= 0 || clock == nil {
		return limiter
	}
	limiter.last = clock.Now()
	limiter.tokens = float64(burst)
	return limiter
}

// refill is called with mu held. Moving the clock backwards cannot earn tokens.
func (l *Limiter) refill() {
	now := l.clock.Now()
	if !now.After(l.last) {
		return
	}
	elapsed := now.Sub(l.last).Seconds()
	l.last = now
	if l.rate <= 0 {
		return
	}
	missing := float64(l.burst) - l.tokens
	// Compare against the time to fill before multiplying a potentially huge rate.
	if elapsed >= missing/l.rate {
		l.tokens = float64(l.burst)
	} else {
		l.tokens += elapsed * l.rate
	}
}

func (l *Limiter) AllowN(n int) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if n <= 0 || n > l.burst || l.burst <= 0 || l.clock == nil {
		return false
	}
	l.refill()
	cost := float64(n)
	if cost > l.tokens {
		return false
	}
	l.tokens -= cost
	return true
}
