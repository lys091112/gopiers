# gopiers


## use module

    go list -m all #列出所有的依赖
    go list -m -versions methodName #列出某个依赖的相关版本

    go get golang.org/x/text #发现指定module的最新版本并更新
    go get rsc.io/sampler@v1.3.1 #通过@加版本号来更新到指定版本

    可以通过如下方式引入同一个module的不同版本
    import (
    "rsc.io/quote"
    quoteV3 "rsc.io/quote/v3"
    )

    go mod tidy removes unused dependencies


## go build

    go help buildmod 可以查看build的类型

###测试

- go tool cover 用来显示help， 覆盖率测试的命令
- go test -coverprofile=c.out  用来生成覆盖率测试文件到c.out
- go tool cover -html=c.out -o coverage.html 将c.out转化为html， 可以通过浏览器进行查看
- go test -bench=. 用来统计以Benchmark开头的测试方法的性能

#### 测试单个文件

```go
  go test -v basic/json_demo.go basic/json_demo_test.go
```

#### 测试单个方法

```go
  go test -v -test.run $methodName 
```

白盒测试的一种技巧： 针对于方法测试中，包含有邮件发送，webhook发送等第三方依赖的代码，为了保证测试不依赖这些，可以将这些包装成一个接口，在测试的时候，传递的一个空实现，不修改现有代码逻辑，还可以屏蔽第三方的依赖影响


## // go: 制定go的编译方式

//go:noinline 顾名思义，不要内联。
```
  Inline: 是在编译期间发生的，将函数调用调用处替换为被调用函数主体的一种编译器优化手段
```
  
//go:nosplit 跳过栈溢出检测

```
  一个 Goroutine 的起始栈大小是有限制的，且比较小的，才可以做到支持并发很多 Goroutine，并高效调度。
  stack.go 源码中可以看到，_StackMin 是 2048 字节，也就是 2k，它不是一成不变的，当不够用时，它会动态地增长。
  那么，必然有一个检测的机制，来保证可以及时地知道栈不够用了，然后再去增长, nosplit 就是将这个跳过这个机制。
  
  不执行栈溢出检查，可以提高性能，但同时也有可能发生 stack overflow 而导致编译失败
```

//go:noescape 禁止逃逸，而且它必须指示一个只有声明没有主体的函数

```
  Go 相比 C、C++ 是内存更为安全的语言，主要一个点就体现在它可以自动地将超出自身生命周期的变量，从函数栈转移到堆中，逃逸就是指这种行为
  
  好处： GC 压力变小了。
      因为它已经告诉编译器，下面的函数无论如何都不会逃逸，那么当函数返回时，其中的资源也会一并都被销毁
      但是绕过编译器的逃逸检查，一旦进入运行时，就有可能导致严重的错误及后果
```  

//go:norace  跳过竞态检测
```
  在多线程程序中，难免会出现数据竞争，正常情况下，当编译器检测到有数据竞争，就会给出提示
  执行 go run -race main.go 利用 -race 来使编译器报告数据竞争问题
  使用 norace 减少编译时间, 但缺点却很明显，那就是数据竞争会导致程序的不确定性
```


#### 开启 GO module

    export GO111MODULE=on
