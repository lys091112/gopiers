package basic

import (
	"fmt"
)

func map_demo() {
	var mapList map[string]int
	mapList = map[string]int{"one": 1, "two": 2}
	mapList["three"] = 3

	mapCreate := make(map[string]float32)
	mapCreate["key1"] = 3.14
	mapCreate["key2"] = 5.26

	fmt.Printf("map literal at \"one\" is : %d\n", mapList["one"])
	fmt.Printf("\"key2\" is: %2f\n", mapCreate["key2"])
	fmt.Printf("assign at two is : %d\n", mapList["two"])

	x := make(map[string]int)
	x["A"] = 0
	x["B"] = 20
	x["C"] = 40

	val, ok := x["E"]
	if ok {
		fmt.Println("----", val)
	} else {
		fmt.Println("E is not exist")
	}

	fmt.Println(x)
	fmt.Println(len(x))
	delete(x, "A")
	fmt.Println(x)

	for key, value := range x {
		fmt.Printf("key is %s -- value is: %d\n", key, value)
	}

	intMap := map[int]int{}

	intMap[0] += 2

	fmt.Printf("intMap=%v\n", intMap)

	fmt.Printf("intMap.2=%d\n", intMap[2])

}
