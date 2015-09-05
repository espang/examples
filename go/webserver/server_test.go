package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("Error getting %s: %s", ts.URL, err)
	}
	defer res.Body.Close()
	fmt.Println(res)

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error reading body: %s", ts.URL, err)
	}

	fmt.Println(string(buf))
}
