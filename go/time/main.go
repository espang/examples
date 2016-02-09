package main

import (
	"fmt"
	"time"
)

func main() {
	cet, _ := time.LoadLocation("CET")
	t1 := time.Now().In(cet)
	t2, _ := time.Parse(time.RFC3339, "2015-12-24T02:45:00Z")
	t2 = t2.In(cet)

	fmt.Printf("t1: %s, t2: %s\n", t1, t2)
	fmt.Printf("t1: %d, t2: %d\n", t1.Day(), t2.Day())
}
