package insect

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/**
* 遇到的问题
* http://studygolang.com/articles/7692
* http://stackoverflow.com/questions/16280176/go-panic-runtime-error-invalid-memory-address-or-nil-pointer-dereference
* TLS handshake timeout
 */
func download(url string) string {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("[E] %v ", r)
		}
	}()

	startTime := time.Now().UnixNano()
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	client := http.DefaultClient
	client.Timeout = 60 * time.Second
	res, e := client.Do(req)
	fmt.Println()
	if e != nil {
		cost := time.Now().UnixNano() - startTime
		fmt.Printf("cost time is %d", cost/(1000000000))
		fmt.Println()
		log.Fatalf("get %s error: %s", url, e)
		fmt.Println()
	}
	fmt.Println("-------" + url)
	body := res.Body
	defer body.Close()

	if res.StatusCode == 200 {
		bodyBytes, _ := ioutil.ReadAll(body)
		str := string(bodyBytes)
		fmt.Print("content", str)
	}
	cost := time.Now().UnixNano() - startTime
	fmt.Printf("cost time is %d", cost/(1000000000))
	return "hello"
}

func Start() {
	url := "https://www.kernel.org/doc/Documentation/"
	download(url)
}
