package filewatcher

import (
	"os"
	"time"
)

type Event struct { Path string }

type Watcher struct { path string; interval time.Duration }

func New(path string, interval time.Duration) *Watcher { return &Watcher{path:path, interval:interval} }

// Start polls file mod time; sends Event on change until stop channel closed.
func (w *Watcher) Start(stop <-chan struct{}) <-chan Event {
	ch := make(chan Event)
	go func(){
		defer close(ch)
		var lastMod time.Time
		for {
			select { case <-stop: return; default: }
			fi, err := os.Stat(w.path)
			if err==nil {
				if fi.ModTime().After(lastMod) { lastMod = fi.ModTime(); ch <- Event{Path:w.path} }
			}
			time.Sleep(w.interval)
		}
	}()
	return ch
}
