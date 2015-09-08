package main

import (
	"bytes"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/cihub/seelog"
)

type FtpToInflux struct {
	Filename     string
	Measurements []string
}

func (f FtpToInflux) String() string {
	return "Tmp: " + f.Filename + ":" + strings.Join(f.Measurements, ", ")
}

type Point struct {
	t time.Time
	v float64
}

type Data struct {
	points []*Point
}

var (
	bufSep       = []byte{';'}
	missingValue = []byte("XXX")
	cet          *time.Location
	layout       = "02.01.2006 15:04"
)

func init() {
	var err error
	cet, err = time.LoadLocation("CET")
	if err != nil {
		seelog.Errorf("Error getting CET-Local: %s", err)
		os.Exit(1)
	}
}

func Transform(buf []byte) []*Data {
	//Split Byte array in single lines
	//The first line is the header
	// followed by an arbitrary number
	// of lines with values
	blines := bytes.Split(buf, []byte{'\r', '\n'})

	//headerline
	//"";zr://BoFiT.D_common_data.MarketData.Load.Load[1].wert
	// The First columns are dates, followed by an arbitrary
	// number of value columns
	// --> count the seperators in a line to get the number
	// of values
	numberOfValues := bytes.Count(blines[0], bufSep)

	datas := make([]*Data, 0, numberOfValues)
	for i := 0; i < numberOfValues; i++ {
		datas = append(datas, &Data{make([]*Point, 0, len(blines)-2)})
	}

	for _, bline := range blines[1 : len(blines)-1] {
		arr := bytes.Split(bline, bufSep)
		t, err := time.ParseInLocation(layout, string(arr[0]), cet)
		if err != nil {
			seelog.Warnf("Could not parse timestring '%s': %s", string(arr[0]), err)
			continue
		}
		for i, s := range arr[1:] {
			if bytes.Compare(s, missingValue) == 0 {
				continue
			}
			fs := strings.TrimRight(strings.Replace(string(s), ",", ".", 1), "\r")
			value, err := strconv.ParseFloat(fs, 64)
			if err != nil {
				seelog.Warnf("Could not convert '%s' to float: %s", fs, err)
				continue
			}
			datas[i].points = append(datas[i].points, &Point{t, value})
		}

	}
	return datas

}
