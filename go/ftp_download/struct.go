package main

import "strings"

type FtpToInflux struct {
	Filename     string
	Measurements []string
}

func (f FtpToInflux) String() string {
	return "Tmp: " + f.Filename + ":" + strings.Join(f.Measurements, ", ")
}

type MyFtpClient struct{}

type FtpCfg struct {
	host, user, pw string
	port           int
}

func NewFtpClient(cfg FtpCfg) (*MyFtpClient, error) {
	return nil, nil
}

func (f *MyFtpClient) Download(fname string) ([]byte, error) {
	return nil, nil
}

type MyInfluxClient struct{}

func NewInfluxClient(path string) (*MyInfluxClient, error) {
	return nil, nil
}

func (i *MyInfluxClient) Write(buf []byte, measures []string) error {
	return nil
}
