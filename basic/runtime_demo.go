package basic

import (
	"fmt"
	"runtime"
)

func first() {

	fmt.Println(runtime.Version())
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(-1))

	go say("workd")
	say("hello")

}

func say(w string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() // 出让当前cpu执行时间片
		fmt.Println(w)
	}
}

// 此方法 会有性能瓶颈，因此除非必要，尽可能不使用
// TODO 关于匿名方法的函数名，待处理
func runtime_callers() {

	// runtime.Caller 返回当前 goroutine 的栈上的函数调用信息. 主要有当前的 pc 值和调用的文件和行号等信息. 若无法获得信息, 返回的 ok 值为 false
	for skip := 0; ; skip++ {

		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		fmt.Printf("caller: pc=%d,file=%s,line=%d\n", pc, file, line)
	}

	// runtime.Callers 把调用它的函数Go程栈上的程序计数器填入切片 pc 中. 参数 skip 为开始在 pc 中记录之前所要跳过的栈帧数, 若为0则表示 runtime.Callers 自身的栈帧, 若为1则表示调用者的栈帧. 该函数返回写入到 pc 切片中的项数
	fmt.Println("=======runtime.Callers============")
	pcs := make([]uintptr, 256)
	for skip := 0; ; skip++ {
		n := runtime.Callers(skip, pcs)
		if n <= 0 {
			break
		}
		fmt.Printf("callers: skip=%d,n=%d,pcs=%v\n", skip, n, pcs[:n])

		if skip == 0 {
			for _, j := range pcs[:n] {
				u := j - 1
				p := runtime.FuncForPC(u)
				file, line := p.FileLine(0)
				fmt.Printf("FuncForPC: pc = %d, file=%s, line=%d, name=%s, entry=%v\n ", u, file, line, p.Name(), p.Entry())
			}
		}
	}
}

func runtime_memstate() {

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	format := "%-40s : %d bytes\n"
	fmt.Printf(format, "bytes allocated and still in use", m.HeapAlloc)
	fmt.Printf(format, "bytes obtained from system", m.HeapSys)
	fmt.Printf(format, "bytes in idle spans", m.HeapIdle)
	fmt.Printf(format, "bytes in non-idle span", m.HeapInuse)
	fmt.Printf(format, "bytes released to the OS", m.HeapReleased)
	fmt.Printf(format, "total number of allocated objects", m.HeapObjects)
}

// debug
// router := httprouter.New()
// router.HandlerFunc("GET", "/debug/pprof", pprof.Index)
// router.Handler("GET", "/debug/heap", pprof.Handler("heap"))
// router.Handler("GET", "/debug/goroutine", pprof.Handler("goroutine"))
// router.Handler("GET", "/debug/block", pprof.Handler("block"))
// router.Handler("GET", "/debug/threadcreate", pprof.Handler("threadcreate"))
// // 启动时的命令,比如 bin/debug -a=1
// router.HandlerFunc("GET", "/debug/pprof/cmdline", pprof.Cmdline)
// router.HandlerFunc("GET", "/debug/pprof/symbol", pprof.Symbol)
// router.HandlerFunc("GET", "/debug/pprof/profile", pprof.Profile)
// router.HandlerFunc("GET", "/debug/pprof/trace", pprof.Trace)
// http.ListenAndServe(":8080", router)
