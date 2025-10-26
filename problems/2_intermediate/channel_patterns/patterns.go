package channelpatterns

// FanIn merges two input channels into one output until both are closed.
func FanIn[T any](a, b <-chan T) <-chan T {
	out := make(chan T)
	go func(){
		defer close(out)
		openA, openB := true, true
		for openA || openB {
			select {
			case v, ok := <-a:
				if ok { out<-v } else { openA = false }
			case v, ok := <-b:
				if ok { out<-v } else { openB = false }
			}
		}
	}()
	return out
}
