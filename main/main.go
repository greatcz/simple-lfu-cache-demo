package main

import (
	"fmt"
	"strconv"

	"github.com/greatcz/simple-lfu-cache-demo/cache"
)

/**
defaultWeepOutAmount = 4
defaultCapacity      = 10
*/
func main() {
	c := cache.New()
	var freq int
	for i := 0; i < 10; {
		freq = 0
		for val := i; val < 12; {
			c.Set(strconv.Itoa(i), val)
			val++
			freq++
		}
		fmt.Printf("key:%d freq:%d \n", i, freq)
		i++
	}

	for i := 0; i < 10; {
		fmt.Printf("key is:%d freq is:%d \n", i, c.GetFreq(strconv.Itoa(i)))
		i++
	}
}
