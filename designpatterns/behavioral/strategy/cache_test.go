package strategy

import "testing"

func TestCache(t *testing.T) {
	lruPolicy := &Lru{}
	lruCache := initCache(lruPolicy)
	lruCache.Add("a", "a")
	lruCache.Add("b", "b")
}
