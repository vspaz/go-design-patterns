package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSingletonOk(t *testing.T) {
	singleton_1 := getSingleton()
	assert.Equal(t, singleton_1.CreateCount, 1)
	assert.Equal(t, singleton_1.MethodInvocationCount, 1)

	singleton_2 := getSingleton()
	assert.Equal(t, singleton_2.CreateCount, 1)
	assert.Equal(t, singleton_2.MethodInvocationCount, 2)
}
