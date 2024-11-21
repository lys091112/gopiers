# gopiers


## 1. use module
- export GO111MODULE=on 开启 GO module 
- go list -m all #列出所有的依赖
- go list -m -versions methodName #列出某个依赖的相关版本
- go get golang.org/x/text #发现指定module的最新版本并更新
- go get rsc.io/sampler@v1.3.1 #通过@加版本号来更新到指定版本

可以通过如下方式引入同一个module的不同版本
  ```go
    import (
    "rsc.io/quote"
    quoteV3 "rsc.io/quote/v3"
    )
  ```

<!-- go mod tidy removes unused dependencies -->


## 2. go build 编译

 - go help buildmod 可以查看build的类型

## 3. 测试

- go tool cover 用来显示help， 覆盖率测试的命令
- go test -coverprofile=c.out  用来生成覆盖率测试文件到c.out
- go tool cover -html=c.out -o coverage.html 将c.out转化为html， 可以通过浏览器进行查看

### 3.1 go benchmark
- go test -bench=. -test.benchmem 用来统计以Benchmark开头的测试方法的性能
支持的参数有：

      -test.bench regexp
            run only benchmarks matching regexp
      -test.benchmem
            print memory allocations for benchmarks
      -test.benchtime d
            run each benchmark for duration d (default 1s)
      -test.blockprofile file
            write a goroutine blocking profile to file
      -test.blockprofilerate rate
            set blocking profile rate (see runtime.SetBlockProfileRate) (default 1)
      -test.count n
            run tests and benchmarks n times (default 1)
      -test.coverprofile file
            write a coverage profile to file
      -test.cpu list
            comma-separated list of cpu counts to run each test with
      -test.cpuprofile file
            write a cpu profile to file
      -test.failfast
            do not start new tests after the first test failure
      -test.fuzz regexp
            run the fuzz test matching regexp
      -test.fuzzcachedir string
            directory where interesting fuzzing inputs are stored (for use only by cmd/go)
      -test.fuzzminimizetime value
            time to spend minimizing a value after finding a failing input (default 1m0s)
      -test.fuzztime value
            time to spend fuzzing; default is to run indefinitely
      -test.fuzzworker
            coordinate with the parent process to fuzz random values (for use only by cmd/go)
      -test.list regexp
            list tests, examples, and benchmarks matching regexp then exit
      -test.memprofile file
            write an allocation profile to file
      -test.memprofilerate rate
            set memory allocation profiling rate (see runtime.MemProfileRate)
      -test.mutexprofile string
            write a mutex contention profile to the named file after execution
      -test.mutexprofilefraction int
            if >= 0, calls runtime.SetMutexProfileFraction() (default 1)
      -test.outputdir dir
            write profiles to dir
      -test.paniconexit0
            panic on call to os.Exit(0)
      -test.parallel n
            run at most n tests in parallel (default 16)
      -test.run regexp
            run only tests and examples matching regexp
      -test.short
            run smaller test suite to save time
      -test.shuffle string
            randomize the execution order of tests and benchmarks (default "off")
      -test.testlogfile file
            write test action log to file (for use only by cmd/go)
      -test.timeout d
            panic test binary after duration d (default 0, timeout disabled)
      -test.trace file
            write an execution trace to file
      -test.v
            verbose: print additional output


### 3.2 测试文件

- 测试单个文件 ``go test -v basic/json_demo.go basic/json_demo_test.go``
- 测试单个方法 ``go test -v -test.run $methodName``

> 白盒测试的一种技巧： 针对于方法测试中，包含有邮件发送，webhook发送等第三方依赖的代码，为了保证测试不依赖这些，可以将这些包装成一个接口，在测试的时候，传递的一个空实现，不修改现有代码逻辑，还可以屏蔽第三方的依赖影响


## 4. 制定go的编译方式

该方式一般以：``// go:`` 开头 

- //go:noinline 顾名思义，不要内联。
```
  Inline: 是在编译期间发生的，将函数调用调用处替换为被调用函数主体的一种编译器优化手段
```
  
- //go:nosplit 跳过栈溢出检测

```
  一个 Goroutine 的起始栈大小是有限制的，且比较小的，才可以做到支持并发很多 Goroutine，并高效调度。
  stack.go 源码中可以看到，_StackMin 是 2048 字节，也就是 2k，它不是一成不变的，当不够用时，它会动态地增长。
  那么，必然有一个检测的机制，来保证可以及时地知道栈不够用了，然后再去增长, nosplit 就是将这个跳过这个机制。
  
  不执行栈溢出检查，可以提高性能，但同时也有可能发生 stack overflow 而导致编译失败
```

- //go:noescape 禁止逃逸，而且它必须指示一个只有声明没有主体的函数

```
  Go 相比 C、C++ 是内存更为安全的语言，主要一个点就体现在它可以自动地将超出自身生命周期的变量，从函数栈转移到堆中，逃逸就是指这种行为
  
  好处： GC 压力变小了。
      因为它已经告诉编译器，下面的函数无论如何都不会逃逸，那么当函数返回时，其中的资源也会一并都被销毁
      但是绕过编译器的逃逸检查，一旦进入运行时，就有可能导致严重的错误及后果
```  

- //go:norace  跳过竞态检测
```
  在多线程程序中，难免会出现数据竞争，正常情况下，当编译器检测到有数据竞争，就会给出提示
  执行 go run -race main.go 利用 -race 来使编译器报告数据竞争问题
  使用 norace 减少编译时间, 但缺点却很明显，那就是数据竞争会导致程序的不确定性
```


