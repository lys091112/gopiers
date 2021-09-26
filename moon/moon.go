package moon

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	// "golang.org/x/sync/semaphore"
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

type Daisy struct {
	Description string
	ImagePath string
	Prev int
	Next int
}

func daisyHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Printf("query=%v\n",query)
	param,ok := query["index"]
	index := 1
	if ok {
		index,_ = strconv.Atoi(param[0])
	}
	fmt.Println("----------------------" , index)
	if index == 1 {
		t :=template.Must(template.ParseFiles("template/html/view/daisy.html"))
		t.Execute(w, map[string]interface{}{"Description": "test","ImagePath": "image01","Prev":1,"Next":index+1})
	}else {
		d := Daisy{Description: "test02",ImagePath: "image02",Prev:index-1,Next:index+1}
		res,_ :=json.Marshal(&d)
		fmt.Printf("%v\n",string(res))
		fmt.Fprintf(w, "%v", string(res))
	}
}

// Start 启动方法
func Start() {
	router := mux.NewRouter()
	// gorilla/mux 和 http在处理匹配地址时不同
	// 静态路由
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("template/css"))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("template/js"))))
	// router.PathPrefix("/download/").Handler(http.StripPrefix("/download/", http.FileServer(http.Dir("/Users/langle/xianyue/own"))))
	router.PathPrefix("/download/").Handler(http.StripPrefix("/download/", http.FileServer(http.Dir("."))))

	// 动态路由
	router.HandleFunc("/", sayHello)
	router.HandleFunc("/index", indexHandler)
	router.HandleFunc("/view/daisy", daisyHandler)
	err := http.ListenAndServe(":8125", router)
	// err := http.ListenAndServeTLS(":443", "resources/server.crt", "resources/server.key", router)
	if err != nil {
		log.Fatalf("Moon Start error! %v", err)
	}
}
