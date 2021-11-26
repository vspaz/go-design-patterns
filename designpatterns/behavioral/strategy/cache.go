package strategy

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

func (c *Cache) Add(key string, value string) {
	if c.Capacity == c.MaxCapacity {
		c.evict()
	}
}

func (c *Cache) Get(key string) {
	delete(c.Storage, key)
}
