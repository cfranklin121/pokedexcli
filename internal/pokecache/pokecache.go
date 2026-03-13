package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mu    sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	if c.cache == nil {
		c.cache = make(map[string]cacheEntry)
	}

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if entry, ok := c.cache[key]; ok {

		return entry.val, true

	}
	return nil, false
}

/*
func NewCache(interval int) Cache {
	wg := sync.WaitGroup{}
	wg.Add(interval)
	new := Cache{}
	return new
}
*/
