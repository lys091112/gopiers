package channel

import "fmt"

// channel 消费示例
func consume() {
	ch := make(chan int) //创建一个无缓存channel

	//新建一个goroutine
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i //往通道写数据
		}
		//不需要再写数据时，关闭channel
		close(ch)
		//ch <- 666 //关闭channel后无法再发送数据

	}()

	// 通过range可以不停的消费channel中的数据，知道channel停止
	for num := range ch {
		fmt.Println("num = ", num)
	}

}
