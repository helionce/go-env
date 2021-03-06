package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync/atomic"
)

var count uint64

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("%+v\n", req)
	fmt.Fprintln(w, strings.Join(os.Environ(), "\n"))
	if req.URL.Path == "/crash" {
		os.Exit(1)
	}
}

func main() {
	http.HandleFunc("/", handler)
	addr := ":" + os.Getenv("PORT")
	fmt.Printf("Listening on %v\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func Count() uint64 {
	atomic.AddUint64(&count, 1)
	return count
}
