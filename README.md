#simple-lfu-cache-demo
a demo about [lfu(Least Frequently Used)](http://dhruvbird.com/lfu.pdf) cache implement in Golang

## example

- when capacity is not enough to hold element then  weed out a amount of low frequency keys
```go
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
	for i := 0; i < 12; {
		freq = 0
		for val := i; val < 12; {
			c.Set(strconv.Itoa(i), val)
			val ++
			freq ++
		}
		fmt.Printf("key:%d freq:%d \n", i, freq)
		i++
	}

	for i := 0 ; i < 12; {
		fmt.Printf("key is:%d freq is:%d \n", i, c.GetFreq(strconv.Itoa(i)))
		i++
	}
}

```
output
```
key:0 freq:12
key:1 freq:11
key:2 freq:10
key:3 freq:9
key:4 freq:8
key:5 freq:7
key:6 freq:6
key:7 freq:5
key:8 freq:4
key:9 freq:3
key:10 freq:2
key:11 freq:1
key is:0 freq is:12
key is:1 freq is:11
key is:2 freq is:10
key is:3 freq is:9
key is:4 freq is:8
key is:5 freq is:7
key is:6 freq is:0
key is:7 freq is:0
key is:8 freq is:0
key is:9 freq is:0
key is:10 freq is:2
key is:11 freq is:1

# freq is 0 mean that the key has been weeded out
```
- capacity is enough
```go
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
			val ++
			freq ++
		}
		fmt.Printf("key:%d freq:%d \n", i, freq)
		i++
	}

	for i := 0 ; i < 10; {
		fmt.Printf("key is:%d freq is:%d \n", i, c.GetFreq(strconv.Itoa(i)))
		i++
	}
}


```
- output
```
key:0 freq:12
key:1 freq:11
key:2 freq:10
key:3 freq:9
key:4 freq:8
key:5 freq:7
key:6 freq:6
key:7 freq:5
key:8 freq:4
key:9 freq:3
key is:0 freq is:12
key is:1 freq is:11
key is:2 freq is:10
key is:3 freq is:9
key is:4 freq is:8
key is:5 freq is:7
key is:6 freq is:6
key is:7 freq is:5
key is:8 freq is:4
key is:9 freq is:3

```

## wiki
http://dhruvbird.com/lfu.pdf
## reference
https://ieftimov.com/post/when-why-least-frequently-used-cache-implementation-golang/
