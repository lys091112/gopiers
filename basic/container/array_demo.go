package container

import (
	"container/list"
	"fmt"
	"strconv"
)

/**
*不同与c / c++ ,go语言的数组是按值传递的,因此传递一个数组的代价特别的大,索性在go中我们可以用切片(slice)来代替

# 切片的含义
s :=[] int {1,2,3 }
直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3
s := arr[:]
初始化切片s,是数组arr的引用
s := arr[startIndex:endIndex]
将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片
s := arr[startIndex:]
缺省endIndex时将表示一直到arr的最后一个元素
s := arr[:endIndex]
缺省startIndex时将表示从arr的第一个元素开始
s1 := s[startIndex:endIndex]
通过切片s初始化切片s1
s :=make([]int,len,cap)
通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片
**/
func array() {
	var x [5]int
	var y = [5]int{1, 2, 3, 4, 5}
	var city = [...]string{"benjing", "tianjin", "hebei"}

	x[0] = 10
	x[1] = 20
	x[2] = 40
	fmt.Printf("%-8T %v\n", x, x)
	fmt.Printf("%-8T %v\n", y, y)
	fmt.Printf("the lenght of city %d \n", len(city))

	var grid1 [3][3]int
	grid1[1][0], grid1[1][1], grid1[1][2] = 2, 4, 8
	grid2 := [3][3]int{{4, 8}, {3, 5, 6}}
	fmt.Printf("%-8T %v\n", grid1, grid1)
	fmt.Printf("%-8T %v\n", grid2, grid2)

	var slice1 []int
	slice1 = y[3:5]
	slice2 := y[3:5]
	fmt.Printf("%p\n", &slice2)
	fmt.Printf("%p\n", &slice1)
	slice1 = append(slice1, 5, 6, 7, 8) //追加后切片的地址没有改变
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", slice1)
	fmt.Printf("%p\n", &slice1)

	slice2 = append(slice2, slice1...) //一个切片合并另一个切片时需要添加...
	fmt.Printf("%v\n", slice2)
	fmt.Printf("%p\n", &slice2)
}

func listdemo() {
	l := list.New()
	for i := 0; i < 5; i++ {
		l.PushBack("kk" + strconv.Itoa(i))
	}
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
