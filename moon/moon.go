package moon

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", "welcome!")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/html/index.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}

// Start 启动方法
func Start() {
	router := mux.NewRouter()
	// gorilla/mux 和 http在处理匹配地址时不同
	// 静态路由
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("template/css"))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("template/js"))))

	// 动态路由
	router.HandleFunc("/", sayHello)
	router.HandleFunc("/index", indexHandler)
	err := http.ListenAndServe(":8125", router)
	// err := http.ListenAndServeTLS(":443", "resources/server.crt", "resources/server.key", router)
	if err != nil {
		log.Fatalf("Moon Start error! %v", err)
	}
}
