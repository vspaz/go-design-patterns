package strategy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCache(t *testing.T) {
	lruPolicy := &Lru{}
	cache := initCache(lruPolicy)
	cache.Add("a", "a")
	cache.Add("b", "b")
	assert.Equal(t, "a", cache.Get("a"))
	cache.setEvictionPolicy(&Lifo{})
}
