package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu sync.Mutex
	c  map[string]int
}

func (c *Counter) Inc(key string) {
	c.c[key]++
}

func (c *Counter) Value(key string) int {
	return c.c[key]
}

func main() {
	key := "test"
	c := Counter{c: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc(key)
	}
	time.Sleep(3 * time.Second)
	fmt.Println(c.Value(key))
}
