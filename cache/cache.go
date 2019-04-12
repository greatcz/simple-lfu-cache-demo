package cache

import (
	"github.com/greatcz/simple-lfu-cache-demo/cache/lfu"
)

func New() *lfu.LfuCache {
	return lfu.New()
}

type CacheInterface interface {
	Set(key string, value interface{})
	Get(key string) interface{}
}
