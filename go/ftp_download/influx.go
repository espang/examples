package main

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/cihub/seelog"
	"github.com/influxdb/influxdb/client"
)

type MyInfluxClient struct {
	client *client.Client
	db     string
}

func NewInfluxClient(surl, db string) (*MyInfluxClient, error) {
	url, err := url.Parse(surl)
	if err != nil {
		seelog.Errorf("Error parsing url: %s", err)
		return nil, err
	}

	cfg := client.NewConfig()
	cfg.URL = *url
	client, err := client.NewClient(cfg)
	if err != nil {
		seelog.Errorf("Error connecting to influxdb: %s", err)
		return nil, err
	}

	return &MyInfluxClient{client, db}, nil
}

func (ic *MyInfluxClient) Write(datas []*Data, measures []string) error {
	if len(datas) != len(measures) {
		return fmt.Errorf("Expect equals length from datas (len:%d) and measures (len:%d)", len(datas), len(measures))
	}
	erroneous := false
	for i, d := range datas {
		err := ic.writeSingle(d, measures[i])
		if err != nil {
			seelog.Warnf("Error writing %d values to %s: %v", len(d.points), measures[i], err)
			erroneous = true
			continue
		}
	}
	if erroneous {
		return fmt.Errorf("Errors during writing the data!")
	}
	return nil
}
func (ic *MyInfluxClient) writeSingle(data *Data, measure string) error {
	start := time.Now()
	lines := make([]string, 0, len(data.points))
	for _, p := range data.points {
		lines = append(lines, fmt.Sprintf("%s value=%f %d", measure, p.v, p.t.UnixNano()))
	}
	n := len(lines)

	size := 2500
	max := len(data.points)
	for index := 0; index < n; {
		end := index + size
		if end >= max {
			end = max - 1
		}
		txt := strings.Join(lines[index:end], "\n")
		_, err := ic.client.WriteLineProtocol(txt, ic.db, "", "", "")
		if err != nil {
			seelog.Warnf("Error writing lines to influxdb: %v", err)
			return err
		}
		index = index + size
	}

	seelog.Tracef("Took %s to write %d values to %s", time.Since(start), len(data.points), measure)
	return nil
}
