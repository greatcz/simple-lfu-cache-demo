package lfu

import (
	"testing"
)

func TestLfuCache_Set(t *testing.T) {
	key := "test_key"
	value := "test_value"
	cache := New()
	cache.Set(key, value)
	getVal := cache.Get(key)
	if getVal != value {
		t.Errorf("get %v want test_value", getVal)
	}
}
