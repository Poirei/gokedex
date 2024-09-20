package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	Val       []byte
}

type Cache struct {
	store map[string]cacheEntry
	mutex sync.RWMutex
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		store: make(map[string]cacheEntry),
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.store[key] = cacheEntry{
		createdAt: time.Now(),
		Val:       val,
	}
}

func (c *Cache) Get(key string) (val []byte, isPresent bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	entry, ok := c.store[key]

	if !ok {
		return nil, false
	}

	return entry.Val, true
}

func (c *Cache) reap() {
	c.mutex.Lock()

	defer c.mutex.Unlock()

	now := time.Now().UnixNano()

	for key, val := range c.store {
		if val.createdAt.UnixNano() < now {
			delete(c.store, key)
		}
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for {
		<-ticker.C
		c.reap()
	}
}
