package network

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetUrl(urls ...string) {
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		d := resp.ContentLength
		resp.Body.Close()
		if err != nil {
			fmt.Fprint(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
		fmt.Printf("%d", d)
	}
}
