package network

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func StartServer() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/webhook", webhook)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UrlPath= %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	fmt.Fprintf(w, "count = %d\n", count)
	mu.Unlock()
}

func webhook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
	r.Context()
	method := r.Method
	fmt.Fprintf(w, "method is %s", method)
}
