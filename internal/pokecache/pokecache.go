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

/*
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	//Add code
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	//Add code
	return []byte{}, false
}
*/
/*
func NewCache(interval int) Cache {
	wg := sync.WaitGroup{}
	wg.Add(interval)
	new := Cache{}
	return new
}
*/
