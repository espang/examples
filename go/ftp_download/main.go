package main

import (
	"bufio"
	"flag"
	"os"
	"strings"

	"github.com/cihub/seelog"
)

const (
	commentPrefix = "#"
	space         = " "
	sep           = ","
)

var host = flag.String("host", "", "The Ftp-Server")
var user = flag.String("user", "", "User of ftp server")
var pw = flag.String("pw", "", "Password of ftp server")
var port = flag.Int("port", 21, "Port of ftp server")
var url = flag.String("influx", "http://localhost:8086", "Url of influx")

func main() {
	defer seelog.Flush()
	flag.Parse()

	cfg := FtpCfg{*host, *user, *pw, *port}
	fClient, err := NewFtpClient(cfg)
	if err != nil {
		panic(err)
	}
	iClient, err := NewInfluxClient(*url)
	if err != nil {
		panic(err)
	}

	files := make([]*FtpToInflux, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		seelog.Tracef("Handle line '%s'", line)
		if strings.HasPrefix(line, commentPrefix) {
			//Comment
			continue
		}
		splittedLine := strings.Split(line, space)
		if len(splittedLine) != 2 {
			seelog.Warnf("Line '%s' has more than one space", line)
			continue
		}
		data := &FtpToInflux{splittedLine[0], strings.Split(splittedLine[1], sep)}
		files = append(files, data)
		seelog.Tracef("Data: %v", data)
	}

	for _, f := range files {
		buf, err := fClient.Download(f.Filename)
		if err != nil {
			seelog.Warnf("Error downloading file '%s': %v", f.Filename, err)
			continue
		}
		err = iClient.Write(buf, f.Measurements)
		if err != nil {
			seelog.Warnf("Error writing Data: %v", err)
			continue
		}
		seelog.Tracef("File '%s' downloaded and written to %d measurements!", f.Filename, len(f.Measurements))
	}
}
