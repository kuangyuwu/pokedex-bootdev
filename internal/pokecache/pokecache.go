package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	lastUseAt time.Time
	val       []byte
}

type Cache struct {
	mutex    *sync.Mutex
	interval time.Duration
	entries  map[string]*cacheEntry
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		mutex:    new(sync.Mutex),
		interval: interval,
		entries:  make(map[string]*cacheEntry),
	}
	go c.reapLoop()
	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	e := cacheEntry{
		lastUseAt: time.Now(),
		val:       val,
	}
	c.entries[key] = &e
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	e, ok := c.entries[key]
	if ok {
		c.entries[key].lastUseAt = time.Now()
		return e.val, true
	} else {
		return nil, false
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.reap()
	}
}

func (c *Cache) reap() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for key, e := range c.entries {
		if time.Since(e.lastUseAt) > c.interval {
			delete(c.entries, key)
		}
	}
}
