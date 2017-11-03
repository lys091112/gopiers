package function

import "fmt"

func A() func() int {
	value := 0
	return func() int {
		value++
		return value
	}
}

// 在循环中他们都绑定了外层的运行环境中的变量，因为外层环境随时会进行更改
//所以直到外层循环结束，这3个routine都不会进行运行
//now i is  0
// now i is  1
// now i is  2
// 3
// 3
// 3
func cycly() {
	done := make(chan int, 3)
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println(i)
			done <- 1
		}()
		fmt.Println("now i is ", i)

	}

	for i := 0; i < 3; i++ {
		<-done
	}
}

//结果是正常的不可预知状态
// now i is  0
// now i is  1
// now i is  2
// index  2
// index  0
// index  1
func cycly2() {
	lens := 5
	done := make(chan int, lens)
	for i := 0; i < lens; i++ {
		go func(index int) {
			fmt.Println("index ", index)
			done <- 1
		}(i)
		fmt.Println("now i is ", i)

	}
	for i := 0; i < lens; i++ {
		<-done
	}
}

/**
func main() {
	// 因为闭包的影响，导致B实际是A的一个函数。 B的作用就是每次使A的局部变量加1
	B := A()

	fmt.Println(B())
	fmt.Println(B())
	fmt.Println(B())

	cycly()
	cycly2()
}
*/
