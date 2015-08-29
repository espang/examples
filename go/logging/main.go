package main

import "github.com/cihub/seelog"

func main() {
	defer seelog.Flush()
	seelog.Info("Hello!")
}
