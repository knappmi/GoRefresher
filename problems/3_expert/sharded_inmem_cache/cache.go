package shardedcache

// TODO: implement sharded cache with TTL & LFU

type Cache struct{}

func New(n uint64) *Cache { return &Cache{} }
func (c *Cache) Set(k string, v any) { /* TODO */ }
func (c *Cache) Get(k string) (any, bool) { return nil, false }
