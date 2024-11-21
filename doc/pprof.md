# golang 性能分析工具

Go语言内置了获取程序运行数据的工具，包括以下两个标准库：

``runtime/pprof``: 采集工具型应用运行数据进行分析
``net/http/pprof``: 采集服务型应用运行时数据进行分析
pprof开启后，每隔一段时间(10ms)就会收集当前的堆栈信息，获取各个函数占用的CPU以及内存资源，然后通过对这些采样数据进行分析，形成一个性能分析报告。

``go tool pprof`` 是对应的命令行指令。它的源数据既可以是一个http地址，也可以是已经获取到的profile文件。使用go tool pprof命令时，既可以采用交互式终端，也可以采用web进行可视化分析，除此之外可以直接将数据生成svg图片，进行静态的分析。


> 非性能分析或测试场景，不要打开pprof

## 2. 常用工具
| Profile项	| 说明	| 详情 |
| -- | --| --|
| allocs |	内存分配		| 从程序启动开始，分配的全部内存 |
| block	| 阻塞		| 导致同步原语阻塞的堆栈跟踪 |
| cmdline | 	命令行调用		| 当前程序的命令行调用 |
| goroutine	| gorouting |	所有当前 goroutine 的堆栈跟踪 |
| heap		| 堆		| 活动对象的内存分配抽样。您可以指定 gc 参数以在获取堆样本之前运行 GC |
| mutex		| 互斥锁		| 争用互斥锁持有者的堆栈跟踪 |
| profile		| CPU分析		| CPU 使用率分析。可以在url中，通过seconds指定持续时间（默认30s）。获取配置文件后，使用 go tool pprof 命令分析CPU使用情况 |
| threadcreate		| 线程创建		| 导致创建新操作系统线程的堆栈跟踪 |
| trace		| 追踪		| 当前程序的执行轨迹。可以在url中，通过seconds指定持续时间（默认30s）。获取跟踪文件后，使用 go tool trace 命令调查跟踪 |

## 3.  使用

### 3.1 初始化

- 在h使用go默认的路由时，初始化时加载 ``import   _ "net/http/pprof"`` 会自动把相应链接加入到路由中

- 使用第g三方路由，需要手动进行路由初始化
```go
    r := http.NewServeMux()
    r.HandleFunc("/", hiHandler)

    // Register pprof handlers
    r.HandleFunc("/debug/pprof/", pprof.Index)
    r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
    r.HandleFunc("/debug/pprof/profile", pprof.Profile)
    r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
    r.HandleFunc("/debug/pprof/trace", pprof.Trace)
```

### 3.2 查看火焰图

以监听 8080 端口为例

1. go tool pprof http://localhost:8080/debug/pprof/goroutine\?second\=20

```
// 查询性能数据，并存储在 /home/xxx/pprof/pprof.main.goroutine.002.pb.gz
Fetching profile over HTTP from http://localhost:8080/debug/pprof/goroutine?second=20
Saved profile in /home/xxx/pprof/pprof.main.goroutine.002.pb.gz
File: main
Type: goroutine
Time: Jul 4, 2022 at 11:00pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof)
```

2. ``go tool pprof -http 0.0.0.0:8081 /home/langle/pprof/pprof.main.goroutine.002.pb.gz``

会自动在浏览器打开界面，通过该界面的view-》flame graph 可以查看火焰图



### 3.3 待补充，包括性能分析如何看等问题
 

## 参考链接

1. [Profiling and optimizing Go web applications](https://artem.krylysov.com/blog/2017/03/13/profiling-and-optimizing-go-web-applications/)
2. [Gin框架中使用pprof](http://liumurong.org/2019/12/gin_pprof/)