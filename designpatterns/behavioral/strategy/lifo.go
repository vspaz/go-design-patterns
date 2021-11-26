package strategy

type Lifo struct{}

func (l *Lifo) evict(c *Cache) string {
	return "'lifo' eviction policy"
}
