package basic

import "fmt"

/**
* 流程控制:
	if a > 0 {
	}
	switch x := f(); {
		case (表达式 Or key1,key2,key3) :
			...
			fallthrough (表示继续向下执行)
			...
			default:
			...
		}

	for true {}
	for (表达式) {}

	for i,v := range r(数组,切片,字符串,映射)

	标号是由一个标识符和冒号组成,可作为goto,break,continue
	语句的目的跳转
	L: for i < n{
		switch i {
			case 5:
				break L
		}
	}
**/

func Control_TestA() {

	fmt.Println("this is Test a")
}

func Control_TestB() (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("test b error: %v", x)
		}
	}()

	panic("demo panic")
}

func Control_TestC() {
	fmt.Println("this is Test c")
}

func Controle_panic() {
	Control_TestA()
	err := Control_TestB()
	fmt.Println(err)
	Control_TestC()

}
