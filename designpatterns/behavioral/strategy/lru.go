package strategy

type Lru struct{}

func (l *Lru) evict() string {
	return "'lru' eviction policy"
}
