package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/jlaffaye/ftp"
)

type MyFtpClient struct {
	conn  *ftp.ServerConn
	login bool
}

type FtpCfg struct {
	host, user, pw string
	port           int
}

func NewFtpClient(cfg FtpCfg) (*MyFtpClient, error) {
	addr := fmt.Sprintf("%s:%d", cfg.host, cfg.port)

	conn, err := ftp.DialTimeout(addr, 5*time.Second)
	if err != nil {
		log.Printf("Error connecting: %s", err)
		return nil, err
	}
	err = conn.Login(cfg.user, cfg.pw)
	if err != nil {
		log.Printf("Error login in: %s", err)
		conn.Quit()
		return nil, err
	}

	return &MyFtpClient{conn, true}, nil
}
func (f *MyFtpClient) Close() {
	if f.login {
		f.conn.Logout()
	}
	f.conn.Quit()
}

func (f *MyFtpClient) Download(fname string) ([]byte, error) {
	reader, err := f.conn.Retr(fname)
	if err != nil {
		log.Printf("Error retrieving '%s': %s", fname, err)
		return nil, err
	}
	defer reader.Close()
	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Printf("Error reading content of '%s': %s", fname, err)
		return nil, err
	}
	return buf, nil
}
