package leetcode

import (
	"sort"
)

const MAC_COUNT = 26

type Trie struct {
	isEnd  bool
	childs map[string]*Trie
}

/** Initialize your data structure here. */
func TireConstructor() *Trie {
	return &Trie{
		isEnd:  false,
		childs: make(map[string]*Trie, MAC_COUNT)}

}

/** Inserts a word into the trie. */
func (this *Trie) insert(word string) {
	if word == "" {
		return
	}

	ws := []rune(word)

	tmp := this
	for i := 0; i < len(ws); i++ {
		key := string(ws[i])
		if c, ok := tmp.childs[key]; ok {
			tmp = c
		} else {
			newTrie := TireConstructor()
			tmp.childs[key] = newTrie
			tmp = newTrie
		}
	}
	tmp.isEnd = true
}

// 212 trie + 深度遍历
func findWords(board [][]byte, words []string) []string {

	if len(board) <= 0 {
		return []string{}
	}

	root := TireConstructor()

	for _, v := range words {
		root.insert(v)
	}

	var keys = make(map[string]bool)
	result := make([]string, 0)
	for i := 0; i < len(board); i ++ {
		for j := 0; j < len(board[0]); j++ {
			tmp := root
			if t, ok := tmp.childs[string(board[i][j])]; !ok {
				continue
			} else {
				var ways = make(map[int]bool)
				strs := find(t, board, i, j, ways)
				if len(strs) <= 0 {
					continue
				}
				for _, v := range strs {
					if _, ok := keys[v]; !ok {
						result = append(result, v)
						keys[v] = true
					}
				}
			}
		}
	}

	sort.Strings(result)
	return result
}

// 1. [i][j] [i-1][j] [i+1][j] [i][j-1] [i][j+1]
// 2. 走过的路不能在走回来
// 3. 如果未走完，不能停止,可能有多重匹配
// 4. 子递归走过的路，在其他分支仍可以重新走
// 通过递归遍历查找
func find(trie *Trie, board [][]byte, i int, j int, ways map[int]bool) []string {
	col := len(board[0])
	ways[i*col+j] = true

	result := make([]string, 0)

	if i > 0 && !ways[(i-1)*col+j] && trie.childs[string(board[i-1][j])] != nil {
		r1 := find(trie.childs[string(board[i-1][j])], board, i-1, j, ways)
		if len(r1) > 0 {
			for _, v := range r1 {
				result = append(result, string(board[i][j])+v)
			}
		}
		delete(ways, (i-1)*col+j) // 为了让下一重递归重新走
	}

	if i < len(board)-1 && !ways[(i+1)*col+j] && trie.childs[string(board[i+1][j])] != nil {
		r1 := find(trie.childs[string(board[i+1][j])], board, i+1, j, ways)
		if len(r1) > 0 {
			for _, v := range r1 {
				result = append(result, string(board[i][j])+v)
			}
		}
		delete(ways, (i+1)*col+j)
	}
	if j > 0 && !ways[i*col+j-1] && trie.childs[string(board[i][j-1])] != nil {
		r1 := find(trie.childs[string(board[i][j-1])], board, i, j-1, ways)
		if len(r1) > 0 {
			for _, v := range r1 {
				result = append(result, string(board[i][j])+v)
			}
		}
		delete(ways, i*col+j-1)
	}

	if j < len(board[0])-1 && !ways[i*col+j+1] && trie.childs[string(board[i][j+1])] != nil {
		r1 := find(trie.childs[string(board[i][j+1])], board, i, j+1, ways)
		if len(r1) > 0 {
			for _, v := range r1 {
				result = append(result, string(board[i][j])+v)
			}
		}
		delete(ways, i*col+j+1)
	}

	if trie.isEnd {
		return append(result, string(board[i][j]))
	}
	return result
}
