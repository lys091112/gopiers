package mystrings

import "bytes"
import "fmt"
import "strings"
import "time"

const (
	Sprintf = iota
	StringAdd
	Join
	BufferAppender
)

func StringAppender(choice, n int) (d time.Duration) {
	v := "Be Alive Keep Fighting!"
	var s string
	var buf bytes.Buffer

	t0 := time.Now()
	for i := 0; i < n; i++ {
		switch choice {
		case Sprintf:
			s = fmt.Sprintf("%s[%s]", s, v)
		case StringAdd:
			s = s + "[" + v + "]"
		case Join:
			s = strings.Join([]string{s, "[", v, "]"}, "")
		case BufferAppender:
			buf.WriteString("[")
			buf.WriteString(v)
			buf.WriteString("]")
		}
	}

	if choice == BufferAppender {
		s = buf.String()
	}

	fmt.Printf("length---> %d", len(s))
	t1 := time.Now()
	d = t1.Sub(t0)
	fmt.Printf("time of way(%d)=%v\n", choice, d)
	return d
}
