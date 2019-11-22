package basic

import "fmt"
import "reflect"
import "unsafe"

// unsafe.Pointer
// （1）任何类型的指针都可以被转化为Pointer
// （2）Pointer可以被转化为任何类型的指针
// （3）uintptr可以被转化为Pointer
// （4）Pointer可以被转化为uintptr
func Float64toBits(f float64) uint64 {
	fmt.Println(reflect.TypeOf(unsafe.Pointer(&f)))
	fmt.Println(reflect.TypeOf((*uint64)(unsafe.Pointer(&f))))
	return *(*uint64)(unsafe.Pointer(&f))
}

func unsafe_assigne() {
	i1 := uint(12)
	i2 := int(13)

	fmt.Println(reflect.TypeOf(i1))
	fmt.Println(reflect.TypeOf(i2))

	fmt.Println(reflect.TypeOf(&i1))
	fmt.Println(reflect.TypeOf(&i2))

	p := &i1
	fmt.Println(reflect.TypeOf(p))

	// cannot use &i2 (type *int) as type *uint in assignment
	// p := &i2
	// so you can use
	p = (*uint)(unsafe.Pointer(&i2))

	fmt.Println(reflect.TypeOf(p))
	fmt.Println(&p)

}

// uintptr is an integer type that is large enough to hold the bit pattern of any pointer.
// 一个unsafe.Pointer指针也可以被转化为uintptr类型，然后保存到指针型数值变量中（注：这只是和当前指针相同的一个数字值，并不是一个指针），然后用以做必要的指针数值运算。（uintptr是一个无符号的整型数，足以保存一个地址）
// 这种转换虽然也是可逆的，但是将uintptr转为unsafe.Pointer指针可能会破坏类型系统，因为并不是所有的数字都是有效的内存地址

func uintprt_demo() {
	var x struct {
		a bool
		b int16
		c []int
	}

	// 和 pb := &x.b 等价
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Print(x.b)

	/**
	 * 有时候垃圾回收器会移动一些变量以降低内存碎片等问题。这类垃圾回收器被称为移动GC。当一个变量被移动，所有的保存改变量旧地址的指针必须同时被更新为变量移动后的新地址。从垃圾收集器的视角来看，一个unsafe.Pointer是一个指向变量的指针，因此当变量被移动是对应的指针也必须被更新；但是uintptr类型的临时变量只是一个普通的数字，所以其值不应该被改变。上面错误的代码因为引入一个非指针的临时变量tmp，导致垃圾收集器无法正确识别这个是一个指向变量x的指针。当第二个语句执行时，变量x可能已经被转移，这时候临时变量tmp也就不再是现在的&x.b地址。**第三个向之前无效地址空间的赋值语句将彻底摧毁整个程序！
	 */
	// NOTE: subtly incorrect!
	// tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	// pb := (*int16)(unsafe.Pointer(tmp))
	// *pb = 42
	// fmt.Print(x.b)
}

type unsafe_W struct {
	b byte
	i int32
	j int64
}

type unsafe_W2 struct {
	b byte
	i int32
	j int16
}

func sizeof_demo() {
	var w *unsafe_W = new(unsafe_W)
	// size = 16
	fmt.Printf("size=%d\n", unsafe.Sizeof(*w))

	// w.j的size为8 ，w.b 占1 ，然后剩余3B用作对齐，w.i占4
	fmt.Printf("w.j size=%d\n", unsafe.Sizeof(w.j))

	fmt.Printf("w.b address = %#016x\n", &w.b)
	fmt.Printf("w.i address = %#016x\n", &w.i)
	fmt.Printf("w.j address = %#016x\n", &w.j)

	p := (*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(w)) + uintptr(unsafe.Sizeof(int32(0)))))
	fmt.Printf("w address = %#016x\n", uintptr(unsafe.Pointer(w)))
	fmt.Printf("p address = %#016x\n", p)
	*p = 33
	fmt.Printf("w.b=%d\n", w.b)
	fmt.Printf("w.i=%d\n", w.i)

	//输出长度为1
	fmt.Printf("size=%d\n", unsafe.Sizeof(struct {
		i8 int8
	}{}))

	//输出长度为16
	fmt.Printf("size=%d\n", unsafe.Sizeof(struct {
		i8 int8
		p  *int8
	}{}))

	var w2 *unsafe_W2 = new(unsafe_W2)
	// size = 12
	fmt.Printf("w2 size=%d\n", unsafe.Sizeof(*w2))

	fmt.Printf("w2.b address = %#016x\n", &w2.b)
	fmt.Printf("w2.i address = %#016x\n", &w2.i)
	fmt.Printf("w2.j address = %#016x\n", &w2.j)
}
