package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cihub/seelog"
	cron "gopkg.in/robfig/cron.v2"
)

type task struct {
	spec string
	url  string
	name string
}

func getTaskFunction(url string, name string) func() {

	return func() {
		resp, err := http.Get(url)
		if err != nil {
			seelog.Warnf("Error starting task '%s': %s", name, err)
			return
		}
		defer resp.Body.Close()

		buf, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			seelog.Warnf("Error reading body: %s", err)
			return
		}

		seelog.Tracef("Responded with state %d", resp.StatusCode)
		seelog.Tracef("Response body: %s", buf)
	}

}

func main() {
	defer seelog.Flush()

	t1 := task{
		spec: "1-59/5 * * * * *",
		url:  "http://localhost:8080/",
		name: "ok-5sec",
	}
	t2 := task{
		spec: "1-59/15 * * * * *",
		url:  "http://localhost:8081/",
		name: "nok-15sec",
	}

	c := cron.New()

	c.AddFunc(t1.spec, getTaskFunction(t1.url, t1.name))
	c.AddFunc(t2.spec, getTaskFunction(t2.url, t2.name))

	c.Start()

	time.Sleep(2 * time.Minute)

	c.Stop()
}
