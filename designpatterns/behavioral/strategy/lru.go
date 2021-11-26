package strategy

type Lru struct{}

func (l *Lru) evict(c *Cache) string {
	return "'lru' eviction policy"
}
