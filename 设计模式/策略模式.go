package main

import "fmt"

type cache struct {
	storage      map[string]string
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e evictionAlgo) *cache {
	storage := make(map[string]string)
	return &cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *cache) setEvictionAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) get(key string) {
	delete(c.storage, key)
}

func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

type evictionAlgo interface {
	evict(c *cache)
}

type fifo struct{}

func (l *fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo strtegy")
}

type lru struct{}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by lru strtegy")
}

type lfu struct{}

func (l *lfu) evict(c *cache) {
	fmt.Println("Evicting by lfu strtegy")
}

func main() {
	lfu := &lfu{}
	cache := initCache(lfu)
	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")

	lru := &lru{}
	cache.setEvictionAlgo(lru)
	cache.add("d", "4")

	fifo := &fifo{}
	cache.setEvictionAlgo(fifo)
	cache.add("e", "5")
}
