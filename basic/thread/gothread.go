package thread

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// GetID 获取协程ID
// SLOW
func GetID() int64 {
	var buf [64]byte
	s := buf[:runtime.Stack(buf[:], false)]
	s = s[len("goroutine "):]
	s = s[:bytes.IndexByte(s, ' ')]
	gid, _ := strconv.ParseInt(string(s), 10, 64)
	return gid
}
var l sync.Mutex

func getLock() {

	go func() {
		l.Lock()
		time.Sleep(10 * time.Second)
		l.Unlock()
	}()

	time.Sleep(3 * time.Second)
	l.Lock()
	fmt.Print("--------------------")
	time.Sleep(10 * time.Minute)
}