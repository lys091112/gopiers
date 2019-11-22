package algorithm

import "testing"

func TestIncreasingTriplet(test *testing.T) {
	elements := [...]int{1, 2, 3, 5, 4}
	res := IncreasingTriplet(elements[:])
	if res {
		test.Log("success")
	} else {
		test.Errorf("except Result is true, but actual is %t", res)
	}
}
