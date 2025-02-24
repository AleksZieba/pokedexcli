package pokecache 

import(
	"time" 
	"sync"
)
type Cache struct {
	entries		map[string]cacheEntry 
	mu			sync.Mutex
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

	ticker := time.NewTicker(interval)
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
	
	if c.entries[key] != nil {
		return c.entries[key].val, true 
	} else {
		return nil, false 
	}
}
	




