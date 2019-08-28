package moon

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", "this is from 2")
}

func Start() {
	router := mux.NewRouter()
	// gorilla/mux 和 http在处理匹配地址时不同
	router.HandleFunc("/", sayHello)
	router.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir("/Users/langle/Downloads/"))))
	//router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/Users/langle/"))))
	router.HandleFunc("/ai-to-alert", getProxyInfo)
	err := http.ListenAndServe(":8125", router)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}

func getProxyInfo(w http.ResponseWriter, r *http.Request) {
	res := make(chan string)
	go getInfo(res)
	fmt.Fprintf(w, "%v", <-res)
}

func getInfo(r chan string) {
	res, _ := http.Get("http://localhost:8000/distribute/alert-with-ai.html")
	body := res.Body
	defer body.Close()

	if res.StatusCode == 200 {
		bodyBytes, _ := ioutil.ReadAll(body)
		str := string(bodyBytes)
		r <- str
	}
	r <- "error"
}
