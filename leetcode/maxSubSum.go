package leetcode

func MaxSubSum(numbers []int) int {
	if numbers == nil || len(numbers) <= 0 {
		return 0
	}
	var max1 int = 0
	var sum1 int = 0
	for _, value := range numbers {
		if sum1 < 0 {
			sum1 = value
		} else {
			sum1 += value
		}
		if max1 < sum1 {
			max1 = sum1
		}
	}
	return max1
}
