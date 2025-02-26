package pokecache 

import(
	"time" 
	"sync" 
)
type Cache struct {
	Entries		map[string]cacheEntry 
	mu			sync.Mutex 
	interval 	time.Duration
} 

type cacheEntry struct {
	createdAt	time.Time
	Val 		[]byte 
} 

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		Entries: make(map[string]cacheEntry),
		interval: interval,
	}

	go c.reapLoop()
	
	return c
} 

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	e := cacheEntry{ 
	createdAt: 	time.Now(), 
	Val: 		val,
	}  
	c.Entries[key] = e 
	}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock() 
	defer c.mu.Unlock()  

	if entry, ok := c.Entries[key]; !ok {
		return nil, false 
	} else if time.Now().Sub(entry.createdAt) > c.interval { 
		return nil, false
	} else {
		return entry.Val, true 
	}
}
	
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval) 
	defer ticker.Stop()
	for {
		_ = <-ticker.C //can this be blocked without the variable??? 
		for k, entry := range c.Entries {
			if time.Now().Sub(entry.createdAt) > c.interval {
				delete(c.Entries, k)
			}
		}
	}
}
