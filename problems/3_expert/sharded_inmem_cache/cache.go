package shardedcache

import (
	"hash/fnv"
	"sync"
	"time"
)

type entry struct { val any; exp int64 }

type shard struct { mu sync.RWMutex; m map[string]entry }

type Cache struct { shards []shard }

func New(n uint64) *Cache { if n==0 { n=1 }; c := &Cache{ shards: make([]shard, n) }; for i:=range c.shards { c.shards[i].m = make(map[string]entry) }; return c }

func (c *Cache) getShard(key string) *shard { h := fnv.New32a(); h.Write([]byte(key)); idx := uint(h.Sum32()) % uint(len(c.shards)); return &c.shards[idx] }

func (c *Cache) Set(k string, v any) { c.SetTTL(k,v,0) }
func (c *Cache) SetTTL(k string, v any, ttl time.Duration) { s := c.getShard(k); s.mu.Lock(); defer s.mu.Unlock(); exp := int64(0); if ttl>0 { exp = time.Now().Add(ttl).UnixNano() }; s.m[k] = entry{val:v, exp:exp} }

func (c *Cache) Get(k string) (any, bool) { s := c.getShard(k); s.mu.RLock(); e, ok := s.m[k]; s.mu.RUnlock(); if !ok { return nil, false }; if e.exp>0 && time.Now().UnixNano() > e.exp { s.mu.Lock(); delete(s.m,k); s.mu.Unlock(); return nil, false }; return e.val, true }
