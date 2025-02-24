package pokecache 

import(
	"time" 
	"sync"
)
type Cache struct {
	entries		map[string]cacheEntry 
	mu			sync.Mutex 
	interval 	time.Duration
} 

type cacheEntry struct {
	createdAt	Time 
	val 		[]byte 
} 

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]cacheEntry),
	}

	//start the background reaping process
	go c.reapLoop(interval)

	
	return c
} 

func (c *Cache) Add(key string, val []byte) 
	c.mu.Lock()
	defer c.mu.Unlock()

	e := cacheEntry{ 
	createdAt: 	time.Now(), 
	val: 		val,
	}  
	c.entries[key] = e 

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock() 
	defer c.mu.Unlock() 
	

	if entry, ok := c.entries[key]; !ok {
		return nil, false 
	} else if time.Now().Sub(entry.createdAt) > c.interval { 
		return nil, false
	} else {
		return entry.val, true 
	}
}
	
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval) 
	defer ticker.Stop()
	for {
		t := <-ticker.C 
		for _, entry := range c.entries {
			if time.Now().Sub(entry.createdAt) > c.interval {
				delete(c.entries, entry)
			}
		}
	}
}




