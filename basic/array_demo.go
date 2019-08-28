package basic

import "fmt"

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

}
