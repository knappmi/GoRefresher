package shardedcache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	c := New(8)
	c.Set("a", 1)
	v, ok := c.Get("a"); if !ok || v.(int)!=1 { t.Fatal("missing a") }
	c.SetTTL("b", 2, 50*time.Millisecond)
	if _, ok := c.Get("b"); !ok { t.Fatal("missing b early") }
	time.Sleep(80*time.Millisecond)
	if _, ok := c.Get("b"); ok { t.Fatal("b should expire") }
}
