package main

import (
	"fmt"
	"math/rand"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/cihub/seelog"
	"github.com/influxdb/influxdb/client"
)

func write_in_batches(data []string, client *client.Client, size int, db string, wg *sync.WaitGroup) {
	start := time.Now()
	n := len(data)
	for index := 0; index < n; {
		txt := strings.Join(data[index:index+size], "\n")
		_, err := client.WriteLineProtocol(txt, db, "", "", "")
		if err != nil {
			panic(err)
		}
		index = index + size
	}
	wg.Done()
	seelog.Infof("Write done in %s", time.Since(start))
}

func main() {

	url, err := url.Parse("http://localhost:8086")
	if err != nil {
		panic(err)
	}

	cfg := client.NewConfig()
	cfg.URL = *url
	client, err := client.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	t := time.Now()
	var i int = 1
	var size int = 1e5
	data := make([]string, 0, size)
	data2 := make([]string, 0, size)
	data3 := make([]string, 0, size)
	data4 := make([]string, 0, size)
	for ; i < size; i++ {
		t = t.Add(time.Second)
		data = append(data, fmt.Sprintf("power_price_go3 value=%f %d", rand.Float64(), t.UnixNano()))
		data2 = append(data2, fmt.Sprintf("power_price_go4 value=%f %d", rand.Float64(), t.UnixNano()))
		data3 = append(data3, fmt.Sprintf("power_price_go5 value=%f %d", rand.Float64(), t.UnixNano()))
		data4 = append(data4, fmt.Sprintf("power_price_go6 value=%f %d", rand.Float64(), t.UnixNano()))
	}

	start := time.Now()
	seelog.Info("Start writing data")
	var wg sync.WaitGroup
	wg.Add(4)
	go write_in_batches(data, client, 10000, "marketdata", &wg)
	go write_in_batches(data2, client, 10000, "marketdata", &wg)
	go write_in_batches(data3, client, 10000, "marketdata", &wg)
	go write_in_batches(data4, client, 10000, "marketdata", &wg)

	wg.Wait()

	seelog.Infof("Took %s to write %d points", time.Since(start), len(data))
	defer seelog.Flush()
}
