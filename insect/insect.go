package insect

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/tl"
)

func download(url string) string {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[E]", r)
		}
	}()

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	client := http.DefaultClient
	tr := &http.Transport{
		TLSClientConfig:&tls.Con
	}
	res, e := client.Do(req)
	//	res, e := http.Get(url)
	if e != nil {
		fmt.Errorf("get %s error: %s", url, e)
	}
	fmt.Println("-------" + url)

	if res.StatusCode == 200 {
		body := res.Body
		defer body.Close()
		bodyBytes, _ := ioutil.ReadAll(body)
		str := string(bodyBytes)
		fmt.Print(str)
	}
	return "hello"
}

func Start() {
	url := "https://www.kernel.org/doc/Documentation/"
	download(url)
}
