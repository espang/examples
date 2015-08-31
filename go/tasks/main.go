package main

import (
	"time"

	"github.com/cihub/seelog"
	cron "gopkg.in/robfig/cron.v2"
)

type task struct {
	folder      string
	prefix      string
	layout      string
	measurement string
	spec        string
}

func main() {
	defer seelog.Flush()

	c := cron.New()

	c.AddFunc("1-59/1 1 * * * *", func() { seelog.Info("Every Minute!") })
	c.AddFunc("1-59/1 * * * * *", func() { seelog.Info("Every Second!") })
	c.ad
	c.Start()

	time.Sleep(2 * time.Minute)

	c.Stop()
}
