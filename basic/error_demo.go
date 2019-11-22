package basic

import "os"
import "errors"
import "fmt"

func Error_use() {

	if _, err := os.Open("not-exist"); err != nil {
		var pathError *os.PathError
		if errors.As(err, &pathError) { // 如果是true，将err的属性复制到pathError
			fmt.Printf("path err %s\n", pathError.Path)
		} else {
			fmt.Print(err)
		}
	}

	e1 := errors.New("error1")
	e2 := errors.New("error2")

	fmt.Printf("e1 is e2? %t\n", errors.Is(e1, e2))
	fmt.Printf("e1 is e2? %t\n", errors.Is(e1, errors.New("error1")))

	// 如果传入的err对象中有%w关键字的格式化类容，则会在返回值中解析出这个原始error
	//多层嵌套只返回第一个，否则返回nil

	e3 := errors.New("error3")
	e4 := fmt.Errorf("error4: %w", e3)
	e5 := fmt.Errorf("error4: %w", e4)
	fmt.Println(e5)
	fmt.Println(errors.Unwrap(e5))
	fmt.Println(e4)
	fmt.Println(errors.Unwrap(e4))
	fmt.Println(e3)
}
