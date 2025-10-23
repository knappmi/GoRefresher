package circuit

// TODO: implement circuit breaker states & backoff

type Breaker struct{}

func New() *Breaker { return &Breaker{} }
func (b *Breaker) Allow() bool { return true }
func (b *Breaker) Fail() {}
func (b *Breaker) Reset() {}

func Backoff(attempt int, base int64) int64 { // TODO: implement exponential backoff
	return 0 }
