package palindromic

type PalindromicTree struct {
}

var pNext [][26]int // 编号为i的节点表示的回文串在两边添加字符c以后变成的回文串的的节点的编号
var pFail []int     // 指向的节点表示的回文串是i节点表示回文串的一个最长后缀
var pCnt []int      // 表示节点i表示的本质不同的串的个数(建树时求出的不是完全的，最后count()函数跑一遍以后才是正确的)
var pNum []int      // 表示以节点i表示的回文串的最右端点为回文串结尾的回文串个数
var pLen []int      // 编号为i的节点表示的回文串的长度(一个节点表示一个回文串)
var S []int         // 表示第i次添加的字符(一开始设S[0]=-1,表示一个在串S中不会出现的字符
var pLast int       // 指向新添加一个字母后所形成的最长回文串表示的节点
var n int           // 添加的字符个数
var p int           // 添加的节点个数

func NewPalindromicTree(n int) PalindromicTree {
	setUp(n)
	return PalindromicTree{}
}

func newNode(l int) int {
	pCnt[p], pNum[p] = 0, 0
	pLen[p] = l
	p += 1
	return p - 1
}

func setUp(n int) {
	pNext = make([][26]int, n)
	pFail = make([]int, n)
	pCnt = make([]int, n)
	pNum = make([]int, n)
	pLen = make([]int, n)
	S = make([]int, n)
	newNode(0)
	newNode(-1)
	pLast, n = 0, 0
	S[n] = -1
	pFail[0] = 1
}

func getFail(x int) int {
	for {
		if S[n-pLen[x]-1] != S[n] {
			x = pFail[x]
		} else {
			return x
		}
	}
}

func PalindromicAdd(c int) {
	c -= 'a'
	n += 1
	S[n] = c
	cur := getFail(pLast)
	if pNext[cur][c] == 0 {
		now := newNode(pLen[cur] + 2)
		pFail[now] = pNext[getFail(pFail[cur])][c]
		pNext[cur][c] = now
		pNum[now] = pNum[pFail[now]] + 1
	}
	pLast = pNext[cur][c]
	pCnt[pLast] += 1
}

// 建树时，求的都是当前节点的数量，但并未计算其指向的fail节点的的数量,所以需要再求和
func count() {
	for i := p - 1; i >= 0; i-- {
		pCnt[pFail[i]] += pCnt[i]
	}
}
