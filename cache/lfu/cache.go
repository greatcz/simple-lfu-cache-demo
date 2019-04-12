package lfu

import (
	"container/list"
)

// wiki:http://dhruvbird.com/lfu.pdf
// reference:https://ieftimov.com/post/when-why-least-frequently-used-cache-implementation-golang/
const (
	initFreq             = 1
	defaultWeepOutAmount = 2
	defaultCapacity      = 1024
)

type CacheItem struct {
	key              string
	value            interface{}
	frequencyPointer *list.Element
}

type FrequencyItem struct {
	holds     map[*CacheItem]bool
	frequency int
}

type LfuCache struct {
	byKey    map[string]*CacheItem // hash map
	freqs    *list.List            // frequency list
	capacity int
	size     int
}

func New() *LfuCache {
	cache := LfuCache{
		byKey:    make(map[string]*CacheItem),
		freqs:    list.New(),
		capacity: defaultCapacity,
		size:     0,
	}

	return &cache
}

func (cache *LfuCache) Set(key string, value interface{}) {
	if item, ok := cache.byKey[key]; ok {
		item.value = value
		cache.increment(item)
	} else {
		item := new(CacheItem)
		item.key = key
		item.value = value
		cache.byKey[key] = item
		cache.size++
		if cache.size > cache.capacity {
			cache.weedOut(defaultWeepOutAmount)
		}
		cache.increment(item)
	}
}

func (cache *LfuCache) Get(key string) interface{} {
	if element, ok := cache.byKey[key]; ok {
		cache.increment(element)
		return element.value
	}

	return nil
}

// weed out amount of element
func (cache *LfuCache) weedOut(amount int) {
	for i := 0; i < amount; {
		if item := cache.freqs.Front(); item != nil {
			for entry, _ := range item.Value.(*FrequencyItem).holds {
				if i < amount {
					delete(cache.byKey, entry.key)
					cache.remove(item, entry)
					cache.size--
					i++
				}
			}
		}
	}

	return
}

func (cache *LfuCache) increment(item *CacheItem) {
	currentFreqPointer := item.frequencyPointer
	var (
		newFreq        int
		newFreqPointer *list.Element
	)

	if currentFreqPointer == nil {
		newFreq = initFreq
		newFreqPointer = cache.freqs.Front()
	} else {
		newFreq = currentFreqPointer.Value.(*FrequencyItem).frequency + 1
		newFreqPointer = currentFreqPointer.Next()
	}

	if newFreqPointer == nil || newFreqPointer.Value.(*FrequencyItem).frequency != newFreq {
		freqItem := new(FrequencyItem)
		freqItem.frequency = newFreq
		freqItem.holds = make(map[*CacheItem]bool)
		if freqItem.frequency == initFreq {
			newFreqPointer = cache.freqs.PushFront(freqItem)
		} else {
			newFreqPointer = cache.freqs.InsertAfter(freqItem, currentFreqPointer)
		}
	}

	item.frequencyPointer = newFreqPointer
	newFreqPointer.Value.(*FrequencyItem).holds[item] = true
	if currentFreqPointer != nil {
		cache.remove(currentFreqPointer, item)
	}
}

// remove one element from old freqPointer
func (cache *LfuCache) remove(freqItem *list.Element, item *CacheItem) {
	delete(freqItem.Value.(*FrequencyItem).holds, item)
	if len(freqItem.Value.(*FrequencyItem).holds) == 0 {
		cache.freqs.Remove(freqItem)
	}
}

// not increment
func (cache *LfuCache) GetFreq(key string) int {
	if element, ok := cache.byKey[key]; ok {
		return element.frequencyPointer.Value.(*FrequencyItem).frequency
	}

	return 0
}
