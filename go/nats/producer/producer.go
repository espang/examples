package main

import (
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/cihub/seelog"
	"github.com/nats-io/nats"
)

type Work struct {
	WorkFor time.Duration
	Values  []float64
	NameIn  string
	NameOut string
}

func main() {
	defer seelog.Flush()
	seelog.Trace("Start nats worker")

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

	sendChannel := make(chan *Work, 20)
	ec.BindSendChan("Work", sendChannel)

	doneChannel := make(chan *Work, 20)
	ec.BindRecvChan("Done", doneChannel)

	var i uint16
	var total int64 = 0
	start := time.Now()
	var maxBytesBehind uint64 = 512 * 512
	go func() {
		for i = 0; i < math.MaxUint16; i++ {

			randTime := time.Duration(rand.Int63n(10)) * time.Millisecond
			total += int64(randTime)
			values := make([]float64, 0, 2500)
			for j := 0; j < 2500; j++ {
				values = append(values, rand.Float64())
			}
			bytesDeltaOver := nc.OutBytes-nc.InBytes >= maxBytesBehind
			if bytesDeltaOver {
				time.Sleep(1 * time.Millisecond)
			}
			seelog.Infof("Work %d ||| bytesDelta %d", i, nc.OutBytes-nc.InBytes)
			sendChannel <- &Work{WorkFor: randTime, NameIn: "In", NameOut: "", Values: values}

		}
	}()

	count := 0
	for count < math.MaxUint16 {
		<-doneChannel
		count++

		if count%10000 == 0 {
			seelog.Infof(" %d of %d messages back", count, math.MaxUint16)
		}
	}
	seelog.Infof("Work done %d tasks finished", count)
	seelog.Infof("Workload of %s, done in %s", time.Duration(total), time.Since(start))
}
