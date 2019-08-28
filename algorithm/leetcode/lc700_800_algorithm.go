package leetcode

import (
	"math"
	"sort"
)

/**
 *  N: 720
 *
 * 如果单词的长度为1， 那么必定为 空字符加上本身
 */
func longestWord(words []string) string {
	if len(words) <= 0 {
		return ""
	}

	var kvs = make(map[string]int)
	for k, v := range words {
		kvs[v] = k
	}

	// 按字符串长度排序
	p := StringSlice2(words)
	sort.Sort(p)

	maxLen := math.MaxInt64
	result := make([]string, 0)
	for _, v := range p {
		if checkWords(v, kvs) {
			if maxLen == math.MaxInt64 {
				maxLen = len(v)
			} else if maxLen > len(v) {
				break
			}
			result = append(result, v)
		}
	}

	if len(result) <= 0 {
		return ""
	}
	if len(result) == 1 {
		return result[0]
	}

	sort.Sort(sort.StringSlice(result))
	return result[0]

}

// 检测单词是否存在
func checkWords(s string, kvs map[string]int) bool {
	if len(s) == 1 {
		return true
	}
	var tmp = s[0 : len(s)-1]

	for {
		if len(tmp) <= 0 {
			return false
		}
		if _, ok := kvs[tmp]; !ok {
			return false
		}
		if len(tmp) == 1 {
			return true
		} else {
			tmp = tmp[0 : len(tmp)-1]
		}
	}

}

type StringSlice2 []string

func (p StringSlice2) Len() int           { return len(p) }
func (p StringSlice2) Less(i, j int) bool { return len(p[i]) > len(p[j]) }
func (p StringSlice2) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
