package moon

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type myHandler struct {
}

var mux map[string]func(w http.ResponseWriter, r *http.Request)

func (*myHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	var path = r.URL.Path
	fmt.Fprintf(w, "path is %s", path)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", "this is from 2")
}

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", sayHello)
	mux.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("/home/langle/xianyue/download"))))
	mux.HandleFunc("/ai-to-alert", get_proxy_info)
	err := http.ListenAndServe(":8125", mux)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}

func get_proxy_info(w http.ResponseWriter, r *http.Request) {
	res := make(chan string)
	go get_info(res)
	fmt.Fprintf(w, "%v", <-res)
}

func get_info(restr chan string) {
	res, _ := http.Get("http://localhost:8000/distribute/alert-with-ai.html")
	body := res.Body
	defer body.Close()

	if res.StatusCode == 200 {
		bodyBytes, _ := ioutil.ReadAll(body)
		str := string(bodyBytes)
		restr <- str
	}
	restr <- "erro"
}
