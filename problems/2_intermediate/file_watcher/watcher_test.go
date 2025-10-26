package filewatcher

import (
	"os"
	"testing"
	"time"
)

func TestWatcher(t *testing.T) {
	fname := t.TempDir()+"/f.txt"
	os.WriteFile(fname, []byte("a"), 0644)
	w := New(fname, 10*time.Millisecond)
	stop := make(chan struct{})
	ch := w.Start(stop)
	os.WriteFile(fname, []byte("b"), 0644)
	select {
	case <-ch: // got event
		close(stop)
	case <-time.After(200*time.Millisecond):
		close(stop); t.Fatal("no event")
	}
}
