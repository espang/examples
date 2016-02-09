// +build ignore
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func MetricStdOutHandler(times <-chan int64) {
	var mu sync.Mutex
	var count, last, delta, total, avg int64

	go func(mu *sync.Mutex) {
		for t := range times {
			mu.Lock()
			count++
			total += t
			mu.Unlock()
		}

	}(&mu)

	tick := time.Tick(1 * time.Second)
	for _ = range tick {
		mu.Lock()
		delta = count - last
		last = count
		if total > 0 {
			avg = delta / total
		}
		mu.Unlock()
		fmt.Printf("Last second: %d Requests with avg time %s \n", delta, time.Duration(avg))
	}
}

var clientPool = sync.Pool{
	New: func() interface{} { return &http.Client{} },
}

func GetIt(times chan<- int64) {
	client := clientPool.Get().(*http.Client)

	s := time.Now()
	resp, err := client.Get("http://localhost:8080/")
	if err != nil {
		log.Printf("error sending get requests: %v", err)
	}
	resp.Body.Close()
	times <- time.Since(s).Nanoseconds()
	clientPool.Put(client)
}

func main() {

	times := make(chan int64, 1000)
	go MetricStdOutHandler(times)

	// 1000 pro sec
	tick := time.Tick(time.Millisecond)
	for _ = range tick {
		go GetIt(times)
	}

}
