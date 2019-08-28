package base

import "fmt"

/**
* 并查集算法
并查集的优化
1、Find_Set(x)时 路径压缩
寻找祖先时我们一般采用递归查找，但是当元素很多亦或是整棵树变为一条链时，每次Find_Set(x)都是O(n)的复杂度，有没有办法减小这个复杂度呢？
答案是肯定的，这就是路径压缩，即当我们经过"递推"找到祖先节点后，"回溯"的时候顺便将它的子孙节点都直接指向祖先，这样以后再次Find_Set(x)时复杂度就变成O(1)了，如下图所示；可见，路径压缩方便了以后的查找。

2、Union(x,y)时 按秩合并
即合并的时候将元素少的集合合并到元素多的集合中，这样合并之后树的高度会相对较小
**/
const (
	MAX int = 10
)

type UnionFind2 struct {
}

var father [MAX]int
var rank [MAX]int

func NewUnionFind() *UnionFind2 {
	return &UnionFind2{}
}

// 设置每个节点都是一个单独的集合
func (unionFind *UnionFind2) Init_set() {
	for i := 0; i < MAX; i++ {
		father[i] = i
		rank[i] = 0
	}
	fmt.Println("hello, world")
}

func (unionFind *UnionFind2) Find(x int) int {
	if x != father[x] {
		father[x] = unionFind.Find(father[x]) //回溯时进行路径压缩
	}
	return father[x]
}

// 将两个集合进行融合,
//按照秩进行合并,并实时的更新秩
func (unionFind *UnionFind2) Union(x, y int) {
	x = unionFind.Find(x)
	y = unionFind.Find(y)
	if x == y {
		return
	}
	if rank[x] > rank[y] {
		father[y] = x
	} else {
		if rank[x] == rank[y] {
			rank[x] += 1
			rank[y] += 1
		}
		father[x] = y
	}
}

func (unionFind *UnionFind2) PrintUnion() {
	for _, k := range father {
		fmt.Print(k)
		fmt.Print(" ")
	}
	fmt.Println()
	for _, k := range rank {
		fmt.Print(k)
		fmt.Print(" ")
	}
}

/*
func main() {
	fmt.Println("vim-go")
	Init_set()
	Union(1, 2)
	Union(2, 5)
	Union(7, 5)
	Union(4, 5)
	Union(3, 7)
	PrintUnion()
}*/
