package basic

import (
	"bytes"
	"fmt"
)

/**
一、数组初始化方式
var [length]Type
var array [5]int //这种方式，只是初始化，不带初始化值，数组长度，已经定义好, 但是其实初始化的值，已经有了并且是该类型的最小值（bool false）,int 0, string ' ' 其他，自行验证

var [N]Type{value1, value2, ... , valueN}
var array  = [5]int{1, 2, 3, 4, 5}  // 这种方式，既初始化变量，又带了初始化值，数组长度已经定义好

var [...]Type{value1, value2, ... , valueN}
var array  = [...]int{1, 2, 3, 4, 5} // 这种方式，既初始化变量，也是带了初始值，数组长度，根据初始值的个数而定，也就是五个

:=[N]Type{value1, value2, ... , valueN}
array :=[5]int{1, 2, 3, 4, 5} // 这种方式，省去 var 关键词，将初始化变量和赋值，放在一起操作,这种方式简单，明了。

:= [...]Type{value1, value2, ... , valueN}
array := [...]int{1, 2, 3, 4, 5} // 这种方式，既初始化变量，也是带了初始值，数组长度，根据初始值的个数而定，也就是五个


二、切片初始化方式
make ( []Type ,length, capacity )
slice := make([]int, 3, 5)

make ( []Type, length)
slice := make([]int, 5)

[]Type{}
slice := []int{}

[]Type{value1 , value2 , ... , valueN }
slice := []int{1, 2, 3, 4, 5} // 这种方式，len是根据初始化长度而定的
*/

// HACK 切片使用共享底层数值的方式来避免数据的拷贝，扩展等操作，因此高效，但是带来的问题是如果已较小的切片引用了
// 一个较大的数据块，从而造成该数据块无法被回收，那么就会产生内存泄漏。 因此碰到这种情况，建议使用copy将需要的数据copy到
// 自己的切片中

/**
 * make 分配大小时，默认会填充零
 */
func make_use() {
	a := make([]int, 4)
	a = append(a, 5)
	fmt.Printf("a.Len = %d\n", len(a)) // 5
	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("--------------------")
	b := make([]int, 0)
	b = append(b, 5)
	fmt.Printf("b.Len = %d\n", len(b))
	for _, v := range b {
		fmt.Println(v)
	}
}

/**
 * slice 初始化时指定分配大小
 * append 方法会动态修改slice的大小，如果当前的slice能容纳元素，则不会扩容
 * 		  如果当前的slice不能容纳元素， 则会先扩容，然后在容纳元素，并返回扩容后的地址
 */
func slice_demo() {
	res := make([]string, 1)
	res[0] = "xxx"
	res = append(res, "a ")
	fmt.Printf("res=%v\n", res)

	var buffer [256]byte

	// inclusive , exclusive
	sli := buffer[100:150]

	addOneToEachElement(sli)
	fmt.Printf("after add, sli=%v\n", sli)

	/// newSlice 和 sli 引用的是同一个底层数组，但是他们的对象不是同一个地址
	newSlice := subtractOneFromLength(sli)
	fmt.Printf("newSlice size=%d, slice size=%d\n", len(newSlice), len(sli))

	fmt.Printf("before len=%d\n", len(sli))
	PtrSubtractOneFromLength(&sli)
	fmt.Println("after len=", len(sli))

	pathName := path("/usr/bin/clean") // Conversion from string to path.
	pathName.TruncateAtFinalSlash()
	fmt.Printf("%s\n", pathName)

	pathName = path("/usr/bin/clean")
	pathName.ToUpper()
	fmt.Printf("%s\n", pathName)

	sli01 := make([]int, 10, 15)
	fmt.Printf("len: %d, cap: %d\n", len(sli01), cap(sli01))

	// Create a couple of starter slices.
	slice := []int{1, 2, 3}
	slice2 := []int{55, 66, 77}
	fmt.Println("Start slice: ", slice)
	fmt.Println("Start slice2:", slice2)

	// Add an item to a slice.
	slice = append(slice, 4)
	fmt.Println("Add one item:", slice)

	// Add one slice to another.
	slice = append(slice, slice2...)
	fmt.Println("Add one slice:", slice)

	// Make a copy of a slice (of int).
	// nil slice
	slice3 := append([]int(nil), slice...)
	fmt.Println("Copy a slice:", slice3)

	// Copy a slice to the end of itself.
	fmt.Println("Before append to self:", slice)
	slice = append(slice, slice...)
	fmt.Println("After append to self:", slice)

}

type path []byte

func (p *path) TruncateAtFinalSlash() {
	i := bytes.LastIndex(*p, []byte("/"))
	if i >= 0 {
		*p = (*p)[0:i]
	}
}

func (p path) ToUpper() {
	for i, b := range p {
		if 'a' <= b && b <= 'z' {
			p[i] = b + 'A' - 'a'
		}
	}
}

// HACK 切片传递的是引用，因此可以修改切片引用对象的值
func addOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}

// 去掉尾元素 传递的是对象引用
func subtractOneFromLength(slice []byte) []byte {
	slice = slice[0 : len(slice)-1]
	return slice
}

// 传递的是指针，因此可以通过修改指针中的引用，来修改sli指向的值
// 与之上的方法比，从而可以感知指针和引用的区别
func PtrSubtractOneFromLength(slicePtr *[]byte) {
	slice := *slicePtr
	*slicePtr = slice[0 : len(slice)-1]
}

/**
 * 二维数组初始化
 * 下表从0开始， 最后一行也要以，结尾
 */
func test() {
	a := [3][4]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
	}
	fmt.Println(a)
}
