package strategy

import "fmt"

type Cache struct {
	Storage        map[string]string
	EvictionPolicy evictionPolicy
	Capacity       int
	MaxCapacity    int
}

func initCache(policy evictionPolicy) *Cache {
	storage := make(map[string]string)
	return &Cache{
		Storage:        storage,
		EvictionPolicy: policy,
		Capacity:       0,
		MaxCapacity:    5,
	}
}

func (c *Cache) setEvictionPolicy(e evictionPolicy) {
	c.EvictionPolicy = e
}

func (c *Cache) evict() {
	c.EvictionPolicy.evict(c)
	c.Capacity--
}

func (c *Cache) Add(key string, value string) string {
	if c.Capacity >= c.MaxCapacity {
		c.evict()
	}
	c.Capacity++
	c.Storage[key] = value
	return fmt.Sprintf("%d", c.Capacity)
}

func (c *Cache) Get(key string) string {
	value, _ := c.Storage[key]
	delete(c.Storage, key)
	return value
}
