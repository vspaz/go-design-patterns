package strategy

type evictionPolicy interface {
	evict(c *Cache) string
}
