package leetcode

import (
	"container/list"
	"sort"
	"strconv"
	"strings"
)

func threeSum(nums []int) [][]int {

	if nums == nil || len(nums) < 3 {
		return nil
	}

	dumplicateStore := make(map[string]int)

	var result = list.New()
	sort.Sort(sort.IntSlice(nums))
	for idx, value := range nums {
		if value > 0 {
			break
		}

		target := 0 - value
		findTwoSums(nums[idx + 1:], target, value, result, dumplicateStore)
	}

	length := result.Len()

	if length <= 0 {
		return nil
	}

	r := [][]int{}
	for e := result.Front(); e != nil; e = e.Next() {
		r = append(r, e.Value.([]int))
	}

	return r
}

func findTwoSums(newNums []int, target int, v1 int, result *list.List, dumplicateStore map[string]int) {
	if newNums == nil || len(newNums) < 2 {
		return
	}

	store := make(map[int]int)

	for idx, v := range newNums {
		store[v] = idx
	}

	for idx, v2 := range newNums {
		v3 := target - v2
		if value, ok := store[v3]; ok && idx < value {
			t := []int{v1, v2, v3}
			unique := join(t, "_")
			if _, ok := dumplicateStore[unique]; !ok {
				dumplicateStore[unique] = 1
				result.PushBack(t)
			}
		}
	}
}

func join(nums []int, seperate string) string {
	if nums == nil || len(nums) <= 0 {
		return ""
	}

	var builder strings.Builder

	for _, value := range nums {
		builder.WriteString(strconv.Itoa(value))
		builder.WriteString(seperate)
	}
	return builder.String()
}
