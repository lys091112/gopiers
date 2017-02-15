package basic

import (
	"fmt"
	"os"
	"regexp"
	"text/template"
)

var pat = `\d{4}\s*[A-HJ-Y][A-HJ-Z]\d*`

func regex_demo() {
	fmt.Println("hello")
	//text := ioutil.ReadFile("")
	var text = "2323 AH20 2324 EH33 2222 EH22"
	reg, _ := regexp.Compile(pat)
	for i, v := range reg.FindAllString(text, 3) {
		fmt.Printf("%d %s \n", i, v)
	}
}

func template_demo() {
	var text = `
	How many roads must {{.}} walk down
	Before they call him {{.}}
	`
	tmp := new(template.Template)
	tmp.Parse(text)
	tmp.Execute(os.Stdout, "a man")
}
