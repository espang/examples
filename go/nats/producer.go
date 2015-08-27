package main

import (
	"log"
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/nats-io/nats"
)

type Timeseries struct {
	Id     uint16
	Times  []time.Time
	Values []float64
	Sum    float64
}

type Stats struct {
	sync.Mutex
	starts    map[uint16]time.Time
	durations map[uint16]time.Duration
}

func (s *Stats) start(id uint16) {
	s.Lock()
	s.starts[id] = time.Now()
	s.Unlock()
}

func (s *Stats) finish(id uint16) {
	s.Lock()
	if val, ok := s.starts[id]; ok {
		s.durations[id] = time.Since(val)
	}
	s.Unlock()
}

var s *Stats

func producer() {

	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, nats.DEFAULT_ENCODER)
	defer c.Close()

	t1 := time.Date(2015, 1, 1, 0, 0, 0, 0, time.Local)
	t2 := time.Date(2015, 1, 2, 0, 0, 0, 0, time.Local)
	times := make([]time.Time, 0, 8760)
	values := make([]float64, 0, 8760)

	total := 0.
	for t2.After(t1) {
		times = append(times, t1)
		val := rand.NormFloat64()
		values = append(values, val)
		total += val
		t1 = t1.Add(time.Hour)
	}

	_, _ = c.Subscribe("timeseries done", func(ts *Timeseries) {
		s.finish(ts.Id)
	})

	var i uint16
	log.Printf("Producer produces")
	for i = 0; i < 2; i++ {
		ts := &Timeseries{Id: i, Times: times, Values: values, Sum: 0.0}
		s.start(i)
		c.Publish("timeseries", ts)
		c.Flush()
	}
	log.Printf("Producer done")
}

func consumer(done <-chan bool) {
	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, nats.DEFAULT_ENCODER)
	defer c.Close()

	_, _ = c.Subscribe("timeseries", func(ts *Timeseries) {
		total := 0.0
		for _, val := range ts.Values {
			total += val
		}
		ts.Sum = total
		c.Publish("timeseries done", ts)
		c.Flush()
	})
	log.Printf("Consumer waiting for done")
	<-done
	log.Printf("Consumer done")
}

func main() {
	s = &Stats{starts: make(map[uint16]time.Time), durations: make(map[uint16]time.Duration)}
	done := make(chan bool)
	go consumer(done)
	producer()
	done <- true
	total := time.Duration(0)
	min := time.Duration(math.MaxInt64)
	max := time.Duration(0)
	for _, dur := range s.durations {
		total += dur
		if max < dur {
			max = dur
		}
		if min > dur {
			min = dur
		}
	}
	mean := float64(total) / float64(len(s.durations))

	log.Printf("Stats: written %d| max %s| min %s| mean %s", len(s.durations), max, min, time.Duration(int64(mean)))
}
