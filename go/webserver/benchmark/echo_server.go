// +build ignore
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var host = flag.String("listen", "localhost", "")
var port = flag.Int("port", 8080, "")

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Got Request")
}

func main() {
	addr := fmt.Sprintf("%s:%d", *host, *port)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(addr, nil))
}
