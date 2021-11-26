package strategy

type Fifo struct{}

func (f *Fifo) evict(c *Cache) string {
	return "fifo eviction policy"
}
