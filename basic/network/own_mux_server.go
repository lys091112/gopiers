package network

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type myhandler struct {
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", "Hello,this is from FileServer3.")
}
func (*myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
}

func startOwn() {
	server := http.Server{
		Addr:        ":9090",
		Handler:     &myhandler{},
		ReadTimeout: 5 * time.Second,
	}
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = sayHello
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
