package basic

import (
	"fmt"
	"testing"
)

func Test_map_demo(test *testing.T) {
	map_demo()
	test.Log("success")

	intMap := map[int]int{}
	intMap[0] += 2
	fmt.Printf("intMap=%v\n", intMap)
	fmt.Printf("intMap.2=%d\n", intMap[2])

	sMap := map[string]string{}

	if sMap["x"] == "" {
		fmt.Println("default value")
	}
	fmt.Printf("v = %s\n", sMap["x"])
}
