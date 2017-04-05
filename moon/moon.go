package moon

import "net/http"
import "fmt"
import "log"

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
	mux.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("/home/langle/xianyue"))))
	err := http.ListenAndServe(":8123", mux)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
