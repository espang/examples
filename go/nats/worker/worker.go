package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/cihub/seelog"
	"github.com/nats-io/nats"
)

var s *Stats

type Work struct {
	WorkFor time.Duration
	Values  []float64
	NameIn  string
	NameOut string
}

type Stats struct {
	sync.Mutex
	TotalTime time.Duration
	WorkCount uint64
}

func (s *Stats) add(t time.Duration) {
	s.Lock()
	s.TotalTime += t
	s.WorkCount++
	s.Unlock()
}

func (s Stats) String() string {
	return fmt.Sprintf("%d Task completted with a total worktime of %s", s.WorkCount, s.TotalTime)
}

func (s *Stats) log() {
	s.Lock()
	seelog.Infof("Stats: %s", s)
	s.Unlock()
}

func worker(tasks <-chan *Work, done chan<- *Work) {
	seelog.Info("Started a worker!")

	for w := range tasks {
		s.add(w.WorkFor)
		w.NameOut = w.NameIn
		w.NameIn = ""
		time.Sleep(w.WorkFor)
		done <- w
	}

}

func main() {
	defer seelog.Flush()
	seelog.Trace("Start nats worker")

	s = &Stats{TotalTime: time.Duration(0), WorkCount: 0}

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		seelog.Errorf("Connecting: %s", err)
		os.Exit(1)
	}

	nc.Opts.AllowReconnect = false
	// Report async errors.
	nc.Opts.AsyncErrorCB = func(nc *nats.Conn, sub *nats.Subscription, err error) {
		seelog.Errorf("NATS: Received an async error! %v\n", err)
		os.Exit(1)
	}

	// Report a disconnect scenario.
	nc.Opts.DisconnectedCB = func(nc *nats.Conn) {
		seelog.Errorf("Getting behind! %d\n", nc.OutMsgs-nc.InMsgs)
		os.Exit(1)
	}

	ec, err := nats.NewEncodedConn(nc, nats.GOB_ENCODER)
	if err != nil {
		seelog.Errorf("Encoded connection: %s", err)
		os.Exit(1)
	}

	recvChannel := make(chan *Work, 20)
	ec.BindRecvQueueChan("Work", "job_workers", recvChannel)

	doneChannel := make(chan *Work, 20)
	ec.BindSendChan("Done", doneChannel)

	go worker(recvChannel, doneChannel)
	go worker(recvChannel, doneChannel)
	go worker(recvChannel, doneChannel)
	go worker(recvChannel, doneChannel)

	tick := time.Tick(time.Minute)
	for _ = range tick {
		s.log()
	}
}
