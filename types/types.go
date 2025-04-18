package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) IncrementBy(factor int) {
	c.total += factor
	c.lastUpdated = time.Now()
}
func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func main() {
	counter := Counter{total: 0}
	counter.IncrementBy(1)
	fmt.Println(counter)
	counterStr := counter.String()
	fmt.Println(counterStr)
}
