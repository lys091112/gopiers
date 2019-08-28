package base

// 这是数组形式，也可以使用链表形式
type UnionFind struct {
	Data []int
}

func (u *UnionFind) MakeSet(size int) {
	u.Data = make([]int, size)
	for idx := range u.Data {
		u.Data[idx] = idx
	}
}

func (u *UnionFind) Union(x, y int) {
	if x >= len(u.Data) || y >= len(u.Data) {
		panic("invalid param")
	}
	pX := u.find(x)
	pY := u.find(y)

	if pX == pY {
		return
	}

	if pX < pY {
		u.Data[pY] = pX
	} else {
		u.Data[pX] = pY
	}

}

// 找寻父节点，并进行路径压缩
func (u *UnionFind) find(i int) int {
	tmp := i
	r := i // 记录根节点
	for {
		if u.Data[i] == i {
			r = i
			break
		} else {
			i = u.Data[i]
		}
	}

	// 路径压缩，将遍历过的点都指向根节点
	j, k := tmp, 0
	for {
		if u.Data[j] != r {
			k = u.Data[j]
			u.Data[j] = r
			j = k
		} else {
			break
		}

	}

	return r
}

// 返回集合分类,最终寻找的同一个父节点的
func (u *UnionFind) collections() map[int][]int {
	res := make(map[int][]int)
	for idx, v := range u.Data {
		pI := u.find(v) // this is the code
		if s, ok := res[pI]; ok {
			res[pI] = append(s, idx)
		} else {
			s := make([]int, 1, 1)
			s[0] = idx
			res[pI] = s
		}
	}
	return res
}
