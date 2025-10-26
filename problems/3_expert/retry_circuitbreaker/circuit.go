package circuit

import "time"

// Breaker states
const (
	stateClosed = iota
	stateOpen
	stateHalf
)

type Breaker struct {
	failures int
	threshold int
	state int
	openedAt time.Time
	cooldown time.Duration
}

func New() *Breaker { return &Breaker{ threshold: 3, state: stateClosed, cooldown: time.Second } }

func (b *Breaker) Allow() bool {
	switch b.state {
	case stateClosed: return true
	case stateOpen:
		if time.Since(b.openedAt) > b.cooldown { b.state = stateHalf; return true }
		return false
	case stateHalf: return true
	}
	return true
}

func (b *Breaker) Fail() { b.failures++; if b.state==stateHalf || b.failures>=b.threshold { b.state = stateOpen; b.openedAt = time.Now(); b.failures = 0 } }
func (b *Breaker) Reset() { b.failures = 0; b.state = stateClosed }

func Backoff(attempt int, base int64) int64 { if attempt < 0 { attempt = 0 }; return base << uint(attempt) }
