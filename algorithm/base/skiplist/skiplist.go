package skiplist

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/lys091112/gopiers/basic/container"
)

// SkipList 跳表
type SkipList struct {
	level int
	head *SkipNode
}

// SkipNode 跳表节点
// TODO key value 可以设计为interface{},然后通过自定义方法进行转化
type SkipNode struct {
	key int 
	value interface{}
	right,down *SkipNode
}

const (
	maxLevel = 32
)

// New 初始化
func New() *SkipList{
	return &SkipList{level:1,head:&SkipNode{key:math.MinInt32,value: nil,}}
}

// Search 查询节点
func (list *SkipList) Search(key int) *SkipNode {
	tmp := list.head
	for {
		if tmp == nil {
			break
		}
		if (*tmp).key == key {
			return tmp
		}else if tmp.right == nil || tmp.right.key > key {
			// 只能向下移动
			tmp = tmp.down
		}else {
			// 向右移动
			tmp = tmp.right
		} 
	}

	return nil
}

// Delete 删除元素
func (list *SkipList) Delete(key int) {
	tmp := list.head
	for {
		if tmp == nil {
			break
		}

		if tmp.right == nil || tmp.right.key > key {
			tmp = tmp.down
		}else if tmp.right.key == key {
			// 通过利用左侧记录比当前key小的值记录
			tmp.right = tmp.right.right
			tmp = tmp.down
		}else  {
			tmp = tmp.right
		}
	}
}

// Insert 新增元素
func (list *SkipList) Insert(key int, value interface{}) error{

	node := list.Search(key)
	if node != nil {
		node.value = value
	}

	tmp := list.head
	stack := container.NewStack()
	for {
		if tmp == nil {
			break
		}

		// 如果右侧为空
		if tmp.right == nil || tmp.right.key > key {
			stack.Push(tmp)
			tmp = tmp.down
		}else if tmp.right.key < key {
			tmp = tmp.right
		}
	}
	currentLevel := 1
	var downNode *SkipNode
	for {
		node := stack.Pop()
		if node == nil {
			break
		}

		// 更新当前节点
		newNode := &SkipNode{key:key,value:value}
		newNode.down = downNode
		downNode = newNode 
		if node.(*SkipNode).right == nil {
			node.(*SkipNode).right = newNode	
		}else {
			newNode.right = node.(*SkipNode).right
			node.(*SkipNode).right = newNode
		}

		// 如果不需要向上层扩展
		// TODO 如果该节点已经扩展到了顶层，就不需要在扩展了
		rand.Seed(time.Now().Local().UnixNano())
		if currentLevel > maxLevel || rand.Float32() > 0.5 {
			break
		}

		currentLevel++
		if currentLevel > list.level {
			highHeadNode := &SkipNode{key:math.MinInt32,value: nil,}
			highHeadNode.down = list.head 
			list.head = highHeadNode
			list.level = currentLevel
			stack.Push(highHeadNode)
		}
	}
	// list.print()
	// fmt.Println("-----------------------------------------------------------")
	return nil
}

func (list *SkipList) print() {
	if nil == list.head {
		log.Print("")
		return
	}

	// 记录底层节点
	bottom := list.head
	for bottom.down != nil {
		bottom = bottom.down
	}


	var res bytes.Buffer
	tempNode := list.head
	for tempNode != nil {						
		seedNode := bottom
		enumNode := tempNode.right
		for enumNode != nil {
			if enumNode.key == seedNode.key {
				res.WriteString(strconv.Itoa(enumNode.key))
				res.WriteString("->")
				res.WriteString(strconv.Itoa(enumNode.value.(int)))
			 	enumNode = enumNode.right
			}else {
				// do nothing
				wt := len(strconv.Itoa(enumNode.key)) + 2 + len(strconv.Itoa(enumNode.value.(int))) 
				for i := 0;i < wt; i++ {
					res.WriteString(" ")
				}
			}
			res.WriteString("\t")
			seedNode = seedNode.right
		}
		res.WriteString("\n")
		tempNode = tempNode.down
	}

	// // 先读取第一层，在读取第二层
	// tmp := list.head
	// for {
	// 	if tmp == nil {
	// 		break
	// 	}
	// 	tr := tmp
	// 	for tr != nil {
	// 		if tr.value != nil {
	// 			res.WriteString(strconv.Itoa(tr.key))
	// 			res.WriteString("->")
	// 			res.WriteString(strconv.Itoa(tr.value.(int)))
	// 		}
	// 		res.WriteString("\t")
	// 		tr = tr.right
	// 	}
	// 	res.WriteString("\n")
	// 	tmp = tmp.down
	// }

	fmt.Print(res.String()) 
}